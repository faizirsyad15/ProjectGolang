package endpoint

import (
	"context"
	"time"

	svc "projectHotel/hotel/jabatan/server"

	pb "projectHotel/hotel/jabatan/grpc"

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
	grpcName = "grpc.JabatanService"
)

func NewGRPCJabatanClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.JabatanService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addJabatanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddJabatanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addJabatanEp = retry
	}

	var readJabatanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadJabatanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readJabatanEp = retry
	}

	var updateJabatanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateJabatan, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateJabatanEp = retry
	}

	var readJabatanByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadJabatanByNama, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readJabatanByNamaEp = retry
	}
	return JabatanEndpoint{AddJabatanEndpoint: addJabatanEp,
		ReadJabatanEndpoint: readJabatanEp, UpdateJabatanEndpoint: updateJabatanEp,
		ReadJabatanByNamaEndpoint: readJabatanByNamaEp}, nil
}

func encodeAddJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Jabatan)
	return &pb.AddJabatanReq{
		IdJabatan:   req.IdJabatan,
		NamaJabatan: req.NamaJabatan,
		Status:      req.Status,
	}, nil
}

func encodeReadJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Jabatan)
	return &pb.UpdateJabatanReq{
		IdJabatan:   req.IdJabatan,
		NamaJabatan: req.NamaJabatan,
		Status:      req.Status,
	}, nil
}

func encodeReadJabatanByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Jabatan)
	return &pb.ReadJabatanByNamaReq{NamaJabatan: req.NamaJabatan}, nil
}

func decodeJabatanResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadJabatanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadJabatanResp)
	var rsp svc.Jabatans

	for _, v := range resp.AllJabatan {
		itm := svc.Jabatan{
			IdJabatan:   v.IdJabatan,
			NamaJabatan: v.NamaJabatan,
			Status:      v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadJabatanbyNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadJabatanByNamaResp)
	return svc.Jabatan{
		IdJabatan:   resp.IdJabatan,
		NamaJabatan: resp.NamaJabatan,
		Status:      resp.Status,
	}, nil
}

func makeClientAddJabatanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddJabatan",
		encodeAddJabatanRequest,
		decodeJabatanResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddJabatan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddJabatan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadJabatanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadJabatan",
		encodeReadJabatanRequest,
		decodeReadJabatanResponse,
		pb.ReadJabatanResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadJabatan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadJabatan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateJabatan(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateJabatan",
		encodeUpdateJabatanRequest,
		decodeJabatanResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateJabatan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateJabatan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadJabatanByNama(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadJabatanByNama",
		encodeReadJabatanByNamaRequest,
		decodeReadJabatanbyNamaRespones,
		pb.ReadJabatanByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadJabatanByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadJabatanByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
