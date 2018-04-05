package endpoint

import (
	"context"
	"time"

	svc "projectHotel/hotel/kamar/server"

	pb "projectHotel/hotel/kamar/grpc"

	util "projectHotel/hotel/util/grpc"
	disc "projectHotel/hotel/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.KamarService"
)

func NewGRPCKamarClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.KamarService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addKamarEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddKamarEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addKamarEp = retry
	}

	var readKamarEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKamarEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKamarEp = retry
	}

	var updateKamarEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateKamar, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateKamarEp = retry
	}

	return KamarEndpoint{AddKamarEndpoint: addKamarEp,
		ReadKamarEndpoint: readKamarEp, UpdateKamarEndpoint: updateKamarEp}, nil
}

func encodeAddKamarRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Kamar)
	return &pb.AddKamarReq{
		IdKamar:     req.IdKamar,
		IdTipeKamar: req.IdTipeKamar,
		IdMenu:      req.IdMenu,
		Status:      req.Status,
	}, nil
}

func encodeReadKamarRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateKamarRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Kamar)
	return &pb.UpdateKamarReq{
		IdKamar:     req.IdKamar,
		IdTipeKamar: req.IdTipeKamar,
		IdMenu:      req.IdMenu,
		Status:      req.Status,
	}, nil
}

func decodeKamarResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadKamarResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKamarResp)
	var rsp svc.Kamars

	for _, v := range resp.AllKamar {
		itm := svc.Kamar{
			IdKamar:     v.IdKamar,
			IdTipeKamar: v.IdTipeKamar,
			IdMenu:      v.IdMenu,
			Status:      v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddKamarEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddKamar",
		encodeAddKamarRequest,
		decodeKamarResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddKamar")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddKamar",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKamarEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKamar",
		encodeReadKamarRequest,
		decodeReadKamarResponse,
		pb.ReadKamarResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKamar")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKamar",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateKamar(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateKamar",
		encodeUpdateKamarRequest,
		decodeKamarResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateKamar")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateKamar",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
