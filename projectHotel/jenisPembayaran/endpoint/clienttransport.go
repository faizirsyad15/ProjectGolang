package endpoint

import (
	"context"
	"time"

	svc "projectHotel/hotel/jenisPembayaran/server"

	pb "projectHotel/hotel/jenisPembayaran/grpc"

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
	grpcName = "grpc.JenisPembayaranService"
)

func NewGRPCJenisPembayaranClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.JenisPembayaranService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addJenisPembayaranEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddJenisPembayaranEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addJenisPembayaranEp = retry
	}

	var readJenisPembayaranEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadJenisPembayaranEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readJenisPembayaranEp = retry
	}

	var updateJenisPembayaranEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateJenisPembayaran, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateJenisPembayaranEp = retry
	}

	var readJenisPembayaranByMetodeEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadJenisPembayaranByMetode, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readJenisPembayaranByMetodeEp = retry
	}
	return JenisPembayaranEndpoint{AddJenisPembayaranEndpoint: addJenisPembayaranEp,
		ReadJenisPembayaranEndpoint: readJenisPembayaranEp, UpdateJenisPembayaranEndpoint: updateJenisPembayaranEp,
		ReadJenisPembayaranByMetodeEndpoint: readJenisPembayaranByMetodeEp}, nil
}

func encodeAddJenisPembayaranRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.JenisPembayaran)
	return &pb.AddJenisPembayaranReq{
		IdJenisPembayaran: req.IdJenisPembayaran,
		MetodePembayaran:  req.MetodePembayaran,
		Status:            req.Status,
	}, nil
}

func encodeReadJenisPembayaranRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateJenisPembayaranRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.JenisPembayaran)
	return &pb.UpdateJenisPembayaranReq{
		IdJenisPembayaran: req.IdJenisPembayaran,
		MetodePembayaran:  req.MetodePembayaran,
		Status:            req.Status,
	}, nil
}

func encodeReadJenisPembayaranByMetodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.JenisPembayaran)
	return &pb.ReadJenisPembayaranByMetodeReq{MetodePembayaran: req.MetodePembayaran}, nil
}

func decodeJenisPembayaranResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadJenisPembayaranResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadJenisPembayaranResp)
	var rsp svc.JenisPembayarans

	for _, v := range resp.AllJenisPembayaran {
		itm := svc.JenisPembayaran{
			IdJenisPembayaran: v.IdJenisPembayaran,
			MetodePembayaran:  v.MetodePembayaran,
			Status:            v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadJenisPembayaranbyMetodeRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadJenisPembayaranByMetodeResp)
	return svc.JenisPembayaran{
		IdJenisPembayaran: resp.IdJenisPembayaran,
		MetodePembayaran:  resp.MetodePembayaran,
		Status:            resp.Status,
	}, nil
}

func makeClientAddJenisPembayaranEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddJenisPembayaran",
		encodeAddJenisPembayaranRequest,
		decodeJenisPembayaranResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddJenisPembayaran")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddJenisPembayaran",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadJenisPembayaranEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadJenisPembayaran",
		encodeReadJenisPembayaranRequest,
		decodeReadJenisPembayaranResponse,
		pb.ReadJenisPembayaranResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadJenisPembayaran")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadJenisPembayaran",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateJenisPembayaran(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateJenisPembayaran",
		encodeUpdateJenisPembayaranRequest,
		decodeJenisPembayaranResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateJenisPembayaran")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateJenisPembayaran",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadJenisPembayaranByMetode(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadJenisPembayaranByMetode",
		encodeReadJenisPembayaranByMetodeRequest,
		decodeReadJenisPembayaranbyMetodeRespones,
		pb.ReadJenisPembayaranByMetodeResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadJenisPembayaranByMetode")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadJenisPembayaranByMetode",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
