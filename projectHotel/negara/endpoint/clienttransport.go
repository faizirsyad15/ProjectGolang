package endpoint

import (
	"context"
	"time"

	svc "projectHotel/hotel/negara/server"

	pb "projectHotel/hotel/negara/grpc"

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
	grpcName = "grpc.NegaraService"
)

func NewGRPCNegaraClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.NegaraService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addNegaraEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddNegaraEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addNegaraEp = retry
	}

	var readNegaraEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadNegaraEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readNegaraEp = retry
	}

	var updateNegaraEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateNegara, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateNegaraEp = retry
	}

	var readNegaraByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadNegaraByNama, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readNegaraByNamaEp = retry
	}
	return NegaraEndpoint{AddNegaraEndpoint: addNegaraEp,
		ReadNegaraEndpoint: readNegaraEp, UpdateNegaraEndpoint: updateNegaraEp,
		ReadNegaraByNamaEndpoint: readNegaraByNamaEp}, nil
}

func encodeAddNegaraRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Negara)
	return &pb.AddNegaraReq{
		IdNegara:   req.IdNegara,
		NamaNegara: req.NamaNegara,
		Status:     req.Status,
	}, nil
}

func encodeReadNegaraRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateNegaraRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Negara)
	return &pb.UpdateNegaraReq{
		IdNegara:   req.IdNegara,
		NamaNegara: req.NamaNegara,
		Status:     req.Status,
	}, nil
}

func encodeReadNegaraByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Negara)
	return &pb.ReadNegaraByNamaReq{NamaNegara: req.NamaNegara}, nil
}

func decodeNegaraResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadNegaraResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadNegaraResp)
	var rsp svc.Negaras

	for _, v := range resp.AllNegara {
		itm := svc.Negara{
			IdNegara:   v.IdNegara,
			NamaNegara: v.NamaNegara,
			Status:     v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadNegarabyNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadNegaraByNamaResp)
	return svc.Negara{
		IdNegara:   resp.IdNegara,
		NamaNegara: resp.NamaNegara,
		Status:     resp.Status,
	}, nil
}

func makeClientAddNegaraEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddNegara",
		encodeAddNegaraRequest,
		decodeNegaraResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddNegara")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddNegara",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadNegaraEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadNegara",
		encodeReadNegaraRequest,
		decodeReadNegaraResponse,
		pb.ReadNegaraResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadNegara")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadNegara",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateNegara(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateNegara",
		encodeUpdateNegaraRequest,
		decodeNegaraResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateNegara")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateNegara",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadNegaraByNama(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadNegaraByNama",
		encodeReadNegaraByNamaRequest,
		decodeReadNegarabyNamaRespones,
		pb.ReadNegaraByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadNegaraByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadNegaraByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
