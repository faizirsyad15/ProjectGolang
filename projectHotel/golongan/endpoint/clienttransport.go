package endpoint

import (
	"context"
	"time"

	svc "projectHotel/hotel/golongan/server"

	pb "projectHotel/hotel/golongan/grpc"

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
	grpcName = "grpc.GolonganService"
)

func NewGRPCGolonganClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.GolonganService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addGolonganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddGolonganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addGolonganEp = retry
	}

	var readGolonganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadGolonganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readGolonganEp = retry
	}

	var updateGolonganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateGolongan, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateGolonganEp = retry
	}

	var readGolonganByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadGolonganByNama, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readGolonganByNamaEp = retry
	}
	return GolonganEndpoint{AddGolonganEndpoint: addGolonganEp,
		ReadGolonganEndpoint: readGolonganEp, UpdateGolonganEndpoint: updateGolonganEp,
		ReadGolonganByNamaEndpoint: readGolonganByNamaEp}, nil
}

func encodeAddGolonganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Golongan)
	return &pb.AddGolonganReq{
		IdGolongan:   req.IdGolongan,
		NamaGolongan: req.NamaGolongan,
		Status:       req.Status,
	}, nil
}

func encodeReadGolonganRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateGolonganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Golongan)
	return &pb.UpdateGolonganReq{
		IdGolongan:   req.IdGolongan,
		NamaGolongan: req.NamaGolongan,
		Status:       req.Status,
	}, nil
}

func encodeReadGolonganByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Golongan)
	return &pb.ReadGolonganByNamaReq{NamaGolongan: req.NamaGolongan}, nil
}

func decodeGolonganResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadGolonganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadGolonganResp)
	var rsp svc.Golongans

	for _, v := range resp.AllGolongan {
		itm := svc.Golongan{
			IdGolongan:   v.IdGolongan,
			NamaGolongan: v.NamaGolongan,
			Status:       v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadGolonganbyNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadGolonganByNamaResp)
	return svc.Golongan{
		IdGolongan:   resp.IdGolongan,
		NamaGolongan: resp.NamaGolongan,
		Status:       resp.Status,
	}, nil
}

func makeClientAddGolonganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddGolongan",
		encodeAddGolonganRequest,
		decodeGolonganResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddGolongan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddGolongan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadGolonganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadGolongan",
		encodeReadGolonganRequest,
		decodeReadGolonganResponse,
		pb.ReadGolonganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadGolongan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadGolongan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateGolongan(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateGolongan",
		encodeUpdateGolonganRequest,
		decodeGolonganResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateGolongan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateGolongan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadGolonganByNama(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadGolonganByNama",
		encodeReadGolonganByNamaRequest,
		decodeReadGolonganbyNamaRespones,
		pb.ReadGolonganByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadGolonganByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadGolonganByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
