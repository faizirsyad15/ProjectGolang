package endpoint

import (
	"context"
	"time"

	svc "projectHotel/hotel/kota/server"

	pb "projectHotel/hotel/kota/grpc"

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
	grpcName = "grpc.KotaService"
)

func NewGRPCKotaClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.KotaService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addKotaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddKotaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addKotaEp = retry
	}

	var readKotaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKotaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKotaEp = retry
	}

	var updateKotaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateKota, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateKotaEp = retry
	}

	var readKotaByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKotaByNama, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKotaByNamaEp = retry
	}
	return KotaEndpoint{AddKotaEndpoint: addKotaEp,
		ReadKotaEndpoint: readKotaEp, UpdateKotaEndpoint: updateKotaEp,
		ReadKotaByNamaEndpoint: readKotaByNamaEp}, nil
}

func encodeAddKotaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Kota)
	return &pb.AddKotaReq{
		IdKota:   req.IdKota,
		NamaKota: req.NamaKota,
		Status:   req.Status,
	}, nil
}

func encodeReadKotaRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateKotaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Kota)
	return &pb.UpdateKotaReq{
		IdKota:   req.IdKota,
		NamaKota: req.NamaKota,
		Status:   req.Status,
	}, nil
}

func encodeReadKotaByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Kota)
	return &pb.ReadKotaByNamaReq{NamaKota: req.NamaKota}, nil
}

func decodeKotaResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadKotaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKotaResp)
	var rsp svc.Kotas

	for _, v := range resp.AllKota {
		itm := svc.Kota{
			IdKota:   v.IdKota,
			NamaKota: v.NamaKota,
			Status:   v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadKotabyNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKotaByNamaResp)
	return svc.Kota{
		IdKota:   resp.IdKota,
		NamaKota: resp.NamaKota,
		Status:   resp.Status,
	}, nil
}

func makeClientAddKotaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddKota",
		encodeAddKotaRequest,
		decodeKotaResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddKota")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddKota",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKotaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKota",
		encodeReadKotaRequest,
		decodeReadKotaResponse,
		pb.ReadKotaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKota")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKota",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateKota(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateKota",
		encodeUpdateKotaRequest,
		decodeKotaResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateKota")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateKota",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKotaByNama(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKotaByNama",
		encodeReadKotaByNamaRequest,
		decodeReadKotabyNamaRespones,
		pb.ReadKotaByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKotaByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKotaByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
