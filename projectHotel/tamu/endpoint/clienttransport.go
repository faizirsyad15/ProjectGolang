package endpoint

import (
	"context"
	"time"

	svc "projectHotel/hotel/tamu/server"

	pb "projectHotel/hotel/tamu/grpc"

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
	grpcName = "grpc.TamuService"
)

func NewGRPCTamuClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.TamuService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addTamuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddTamuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addTamuEp = retry
	}

	var readTamuByTeleponEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadTamuByTeleponEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readTamuByTeleponEp = retry
	}

	var readTamuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadTamuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readTamuEp = retry
	}

	var updateTamuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateTamu, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateTamuEp = retry
	}

	var readTamuByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadTamuByNama, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readTamuByNamaEp = retry
	}
	return TamuEndpoint{AddTamuEndpoint: addTamuEp, ReadTamuByTeleponEndpoint: readTamuByTeleponEp,
		ReadTamuEndpoint: readTamuEp, UpdateTamuEndpoint: updateTamuEp,
		ReadTamuByNamaEndpoint: readTamuByNamaEp}, nil
}

func encodeAddTamuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Tamu)
	return &pb.AddTamuReq{
		IdTamu:       req.IdTamu,
		NamaTamu:     req.NamaTamu,
		NoTelepon:    req.NoTelepon,
		JenisKelamin: req.JenisKelamin,
		IdAlamat:     req.IdAlamat,
		Status:       req.Status,
	}, nil
}

func encodeReadTamuByTeleponRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Tamu)
	return &pb.ReadTamuByTeleponReq{NoTelepon: req.NoTelepon}, nil
}

func encodeReadTamuRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateTamuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Tamu)
	return &pb.UpdateTamuReq{
		IdTamu:       req.IdTamu,
		NamaTamu:     req.NamaTamu,
		NoTelepon:    req.NoTelepon,
		JenisKelamin: req.JenisKelamin,
		IdAlamat:     req.IdAlamat,
		Status:       req.Status,
	}, nil
}

func encodeReadTamuByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Tamu)
	return &pb.ReadTamuByNamaReq{NamaTamu: req.NamaTamu}, nil
}

func decodeTamuResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadTamuByTeleponRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadTamuByTeleponResp)
	return svc.Tamu{
		IdTamu:       resp.IdTamu,
		NamaTamu:     resp.NamaTamu,
		NoTelepon:    resp.NoTelepon,
		JenisKelamin: resp.JenisKelamin,
		IdAlamat:     resp.IdAlamat,
		Status:       resp.Status,
	}, nil
}

func decodeReadTamuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadTamuResp)
	var rsp svc.Tamus

	for _, v := range resp.AllTamu {
		itm := svc.Tamu{
			IdTamu:       v.IdTamu,
			NamaTamu:     v.NamaTamu,
			NoTelepon:    v.NoTelepon,
			JenisKelamin: v.JenisKelamin,
			IdAlamat:     v.IdAlamat,
			Status:       v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadTamubyNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadTamuByNamaResp)
	return svc.Tamu{
		IdTamu:       resp.IdTamu,
		NamaTamu:     resp.NamaTamu,
		NoTelepon:    resp.NoTelepon,
		JenisKelamin: resp.JenisKelamin,
		IdAlamat:     resp.IdAlamat,
		Status:       resp.Status,
	}, nil
}

func makeClientAddTamuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddTamu",
		encodeAddTamuRequest,
		decodeTamuResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddTamu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddTamu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadTamuByTeleponEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadTamuByTelepon",
		encodeReadTamuByTeleponRequest,
		decodeReadTamuByTeleponRespones,
		pb.ReadTamuByTeleponResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadTamuByTelepon")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadTamuByTelepon",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadTamuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadTamu",
		encodeReadTamuRequest,
		decodeReadTamuResponse,
		pb.ReadTamuResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadTamu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadTamu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateTamu(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateTamu",
		encodeUpdateTamuRequest,
		decodeTamuResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateTamu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateTamu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadTamuByNama(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadTamuByNama",
		encodeReadTamuByNamaRequest,
		decodeReadTamubyNamaRespones,
		pb.ReadTamuByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadTamuByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadTamuByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
