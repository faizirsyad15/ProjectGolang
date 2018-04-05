package endpoint

import (
	"context"
	"time"

	svc "projectHotel/hotel/tipeKamar/server"

	pb "projectHotel/hotel/tipeKamar/grpc"

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
	grpcName = "grpc.TipeKamarService"
)

func NewGRPCTipeKamarClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.TipeKamarService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addTipeKamarEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddTipeKamarEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addTipeKamarEp = retry
	}

	var readTipeKamarByHargaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadTipeKamarByHargaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readTipeKamarByHargaEp = retry
	}

	var readTipeKamarEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadTipeKamarEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readTipeKamarEp = retry
	}

	var updateTipeKamarEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateTipeKamar, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateTipeKamarEp = retry
	}

	var readTipeKamarByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadTipeKamarByNama, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readTipeKamarByNamaEp = retry
	}
	return TipeKamarEndpoint{AddTipeKamarEndpoint: addTipeKamarEp, ReadTipeKamarByHargaEndpoint: readTipeKamarByHargaEp,
		ReadTipeKamarEndpoint: readTipeKamarEp, UpdateTipeKamarEndpoint: updateTipeKamarEp,
		ReadTipeKamarByNamaEndpoint: readTipeKamarByNamaEp}, nil
}

func encodeAddTipeKamarRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.TipeKamar)
	return &pb.AddTipeKamarReq{
		IdTipeKamar:   req.IdTipeKamar,
		NamaTipeKamar: req.NamaTipeKamar,
		HargaKamar:    req.HargaKamar,
		Status:        req.Status,
	}, nil
}

func encodeReadTipeKamarByHargaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.TipeKamar)
	return &pb.ReadTipeKamarByHargaReq{HargaKamar: req.HargaKamar}, nil
}

func encodeReadTipeKamarRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateTipeKamarRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.TipeKamar)
	return &pb.UpdateTipeKamarReq{
		IdTipeKamar:   req.IdTipeKamar,
		NamaTipeKamar: req.NamaTipeKamar,
		HargaKamar:    req.HargaKamar,
		Status:        req.Status,
	}, nil
}

func encodeReadTipeKamarByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.TipeKamar)
	return &pb.ReadTipeKamarByNamaReq{NamaTipeKamar: req.NamaTipeKamar}, nil
}

func decodeTipeKamarResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadTipeKamarByHargaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadTipeKamarByHargaResp)
	return svc.TipeKamar{
		IdTipeKamar:   resp.IdTipeKamar,
		NamaTipeKamar: resp.NamaTipeKamar,
		HargaKamar:    resp.HargaKamar,
		Status:        resp.Status,
	}, nil
}

func decodeReadTipeKamarResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadTipeKamarResp)
	var rsp svc.TipeKamars

	for _, v := range resp.AllTipeKamar {
		itm := svc.TipeKamar{
			IdTipeKamar:   v.IdTipeKamar,
			NamaTipeKamar: v.NamaTipeKamar,
			HargaKamar:    v.HargaKamar,
			Status:        v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadTipeKamarbyNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadTipeKamarByNamaResp)
	return svc.TipeKamar{
		IdTipeKamar:   resp.IdTipeKamar,
		NamaTipeKamar: resp.NamaTipeKamar,
		HargaKamar:    resp.HargaKamar,
		Status:        resp.Status,
	}, nil
}

func makeClientAddTipeKamarEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddTipeKamar",
		encodeAddTipeKamarRequest,
		decodeTipeKamarResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddTipeKamar")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddTipeKamar",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadTipeKamarByHargaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadTipeKamarByHarga",
		encodeReadTipeKamarByHargaRequest,
		decodeReadTipeKamarByHargaRespones,
		pb.ReadTipeKamarByHargaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadTipeKamarByHarga")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadTipeKamarByHarga",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadTipeKamarEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadTipeKamar",
		encodeReadTipeKamarRequest,
		decodeReadTipeKamarResponse,
		pb.ReadTipeKamarResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadTipeKamar")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadTipeKamar",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateTipeKamar(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateTipeKamar",
		encodeUpdateTipeKamarRequest,
		decodeTipeKamarResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateTipeKamar")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateTipeKamar",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadTipeKamarByNama(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadTipeKamarByNama",
		encodeReadTipeKamarByNamaRequest,
		decodeReadTipeKamarbyNamaRespones,
		pb.ReadTipeKamarByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadTipeKamarByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadTipeKamarByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
