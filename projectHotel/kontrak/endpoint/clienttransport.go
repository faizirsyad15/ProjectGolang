package endpoint

import (
	"context"
	"time"

	svc "projectHotel/hotel/kontrak/server"

	pb "projectHotel/hotel/kontrak/grpc"

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
	grpcName = "grpc.KontrakService"
)

func NewGRPCKontrakClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.KontrakService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addKontrakEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddKontrakEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addKontrakEp = retry
	}

	var readKontrakBySelesaiEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKontrakBySelesaiEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKontrakBySelesaiEp = retry
	}

	var readKontrakEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKontrakEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKontrakEp = retry
	}

	var updateKontrakEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateKontrak, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateKontrakEp = retry
	}

	var readKontrakByMulaiEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKontrakByMulai, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKontrakByMulaiEp = retry
	}
	return KontrakEndpoint{AddKontrakEndpoint: addKontrakEp, ReadKontrakBySelesaiEndpoint: readKontrakBySelesaiEp,
		ReadKontrakEndpoint: readKontrakEp, UpdateKontrakEndpoint: updateKontrakEp,
		ReadKontrakByMulaiEndpoint: readKontrakByMulaiEp}, nil
}

func encodeAddKontrakRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Kontrak)
	return &pb.AddKontrakReq{
		IdKontrak:      req.IdKontrak,
		TanggalMulai:   req.TanggalMulai,
		TanggalSelesai: req.TanggalSelesai,
		Keterangan:     req.Keterangan,
		Status:         req.Status,
	}, nil
}

func encodeReadKontrakBySelesaiRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Kontrak)
	return &pb.ReadKontrakBySelesaiReq{TanggalSelesai: req.TanggalSelesai}, nil
}

func encodeReadKontrakRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateKontrakRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Kontrak)
	return &pb.UpdateKontrakReq{
		IdKontrak:      req.IdKontrak,
		TanggalMulai:   req.TanggalMulai,
		TanggalSelesai: req.TanggalSelesai,
		Keterangan:     req.Keterangan,
		Status:         req.Status,
	}, nil
}

func encodeReadKontrakByMulaiRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Kontrak)
	return &pb.ReadKontrakByMulaiReq{TanggalMulai: req.TanggalMulai}, nil
}

func decodeKontrakResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadKontrakBySelesaiRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKontrakBySelesaiResp)
	return svc.Kontrak{
		IdKontrak:      resp.IdKontrak,
		TanggalMulai:   resp.TanggalMulai,
		TanggalSelesai: resp.TanggalSelesai,
		Keterangan:     resp.Keterangan,
		Status:         resp.Status,
	}, nil
}

func decodeReadKontrakResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKontrakResp)
	var rsp svc.Kontraks

	for _, v := range resp.AllKontrak {
		itm := svc.Kontrak{
			IdKontrak:      v.IdKontrak,
			TanggalMulai:   v.TanggalMulai,
			TanggalSelesai: v.TanggalSelesai,
			Keterangan:     v.Keterangan,
			Status:         v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadKontrakbyNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKontrakByMulaiResp)
	return svc.Kontrak{
		IdKontrak:      resp.IdKontrak,
		TanggalMulai:   resp.TanggalMulai,
		TanggalSelesai: resp.TanggalSelesai,
		Keterangan:     resp.Keterangan,
		Status:         resp.Status,
	}, nil
}

func makeClientAddKontrakEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddKontrak",
		encodeAddKontrakRequest,
		decodeKontrakResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddKontrak")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddKontrak",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKontrakBySelesaiEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKontrakBySelesai",
		encodeReadKontrakBySelesaiRequest,
		decodeReadKontrakBySelesaiRespones,
		pb.ReadKontrakBySelesaiResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKontrakBySelesai")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKontrakBySelesai",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKontrakEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKontrak",
		encodeReadKontrakRequest,
		decodeReadKontrakResponse,
		pb.ReadKontrakResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKontrak")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKontrak",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateKontrak(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateKontrak",
		encodeUpdateKontrakRequest,
		decodeKontrakResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateKontrak")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateKontrak",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKontrakByMulai(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKontrakByMulai",
		encodeReadKontrakByMulaiRequest,
		decodeReadKontrakbyNamaRespones,
		pb.ReadKontrakByMulaiResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKontrakByMulai")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKontrakByMulai",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
