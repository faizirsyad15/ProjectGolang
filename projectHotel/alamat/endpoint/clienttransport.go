package endpoint

import (
	"context"
	"time"

	svc "projectHotel/hotel/alamat/server"

	pb "projectHotel/hotel/alamat/grpc"

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
	grpcName = "grpc.AlamatService"
)

func NewGRPCAlamatClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.AlamatService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addAlamatEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddAlamatEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addAlamatEp = retry
	}

	var readAlamatByNoRumahEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadAlamatByNoRumahEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readAlamatByNoRumahEp = retry
	}

	var readAlamatEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadAlamatEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readAlamatEp = retry
	}

	var updateAlamatEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateAlamat, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateAlamatEp = retry
	}

	return AlamatEndpoint{AddAlamatEndpoint: addAlamatEp, ReadAlamatByNoRumahEndpoint: readAlamatByNoRumahEp,
		ReadAlamatEndpoint: readAlamatEp, UpdateAlamatEndpoint: updateAlamatEp}, nil
}

func encodeAddAlamatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Alamat)
	return &pb.AddAlamatReq{
		IdAlamat:    req.IdAlamat,
		AlamatRumah: req.AlamatRumah,
		RtRw:        req.RtRw,
		NoRumah:     req.NoRumah,
		IdKota:      req.IdKota,
		IdNegara:    req.IdNegara,
		Status:      req.Status,
	}, nil
}

func encodeReadAlamatByNoRumahRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Alamat)
	return &pb.ReadAlamatByNoRumahReq{NoRumah: req.NoRumah}, nil
}

func encodeReadAlamatRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateAlamatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Alamat)
	return &pb.UpdateAlamatReq{
		IdAlamat:    req.IdAlamat,
		AlamatRumah: req.AlamatRumah,
		RtRw:        req.RtRw,
		NoRumah:     req.NoRumah,
		IdKota:      req.IdKota,
		IdNegara:    req.IdNegara,
		Status:      req.Status,
	}, nil
}

func decodeAlamatResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadAlamatByNoRumahRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadAlamatByNoRumahResp)
	return svc.Alamat{
		IdAlamat:    resp.IdAlamat,
		AlamatRumah: resp.AlamatRumah,
		RtRw:        resp.RtRw,
		NoRumah:     resp.NoRumah,
		IdKota:      resp.IdKota,
		IdNegara:    resp.IdNegara,
		Status:      resp.Status,
	}, nil
}

func decodeReadAlamatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadAlamatResp)
	var rsp svc.Alamats

	for _, v := range resp.AllAlamat {
		itm := svc.Alamat{
			IdAlamat:    v.IdAlamat,
			AlamatRumah: v.AlamatRumah,
			RtRw:        v.RtRw,
			NoRumah:     v.NoRumah,
			IdKota:      v.IdKota,
			IdNegara:    v.IdNegara,
			Status:      v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddAlamatEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddAlamat",
		encodeAddAlamatRequest,
		decodeAlamatResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddAlamat")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddAlamat",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadAlamatByNoRumahEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadAlamatByNoRumah",
		encodeReadAlamatByNoRumahRequest,
		decodeReadAlamatByNoRumahRespones,
		pb.ReadAlamatByNoRumahResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadAlamatByNoRumah")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadAlamatByNoRumah",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadAlamatEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadAlamat",
		encodeReadAlamatRequest,
		decodeReadAlamatResponse,
		pb.ReadAlamatResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadAlamat")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadAlamat",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateAlamat(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateAlamat",
		encodeUpdateAlamatRequest,
		decodeAlamatResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateAlamat")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateAlamat",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
