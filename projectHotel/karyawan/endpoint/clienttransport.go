package endpoint

import (
	"context"
	"time"

	svc "projectHotel/hotel/karyawan/server"

	pb "projectHotel/hotel/karyawan/grpc"

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
	grpcName = "grpc.KaryawanService"
)

func NewGRPCKaryawanClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.KaryawanService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addKaryawanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddKaryawanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addKaryawanEp = retry
	}

	var readKaryawanByTeleponEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKaryawanByTeleponEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKaryawanByTeleponEp = retry
	}

	var readKaryawanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKaryawanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKaryawanEp = retry
	}

	var updateKaryawanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateKaryawan, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateKaryawanEp = retry
	}

	var readKaryawanByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKaryawanByNama, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKaryawanByNamaEp = retry
	}

	var readKaryawanByKeteranganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKaryawanByKeteranganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKaryawanByKeteranganEp = retry
	}

	return KaryawanEndpoint{AddKaryawanEndpoint: addKaryawanEp, ReadKaryawanByTeleponEndpoint: readKaryawanByTeleponEp,
		ReadKaryawanEndpoint: readKaryawanEp, UpdateKaryawanEndpoint: updateKaryawanEp,
		ReadKaryawanByNamaEndpoint:       readKaryawanByNamaEp,
		ReadKaryawanByKeteranganEndpoint: readKaryawanByKeteranganEp}, nil
}

func encodeAddKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.AddKaryawanReq{
		IdKaryawan:   req.IdKaryawan,
		NamaKaryawan: req.NamaKaryawan,
		Alamat:       req.Alamat,
		NoTelepon:    req.NoTelepon,
		Keterangan:   req.Keterangan,
		IdJabatan:    req.IdJabatan,
		IdDepartemen: req.IdDepartemen,
		IdGolongan:   req.IdGolongan,
		IdKontrak:    req.IdKontrak,
		Status:       req.Status,
	}, nil
}

func encodeReadKaryawanByTeleponRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.ReadKaryawanByTeleponReq{NoTelepon: req.NoTelepon}, nil
}

func encodeReadKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.UpdateKaryawanReq{
		IdKaryawan:   req.IdKaryawan,
		NamaKaryawan: req.NamaKaryawan,
		Alamat:       req.Alamat,
		NoTelepon:    req.NoTelepon,
		Keterangan:   req.Keterangan,
		IdJabatan:    req.IdJabatan,
		IdDepartemen: req.IdDepartemen,
		IdGolongan:   req.IdGolongan,
		IdKontrak:    req.IdKontrak,
		Status:       req.Status,
	}, nil
}

func encodeReadKaryawanByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.ReadKaryawanByNamaReq{NamaKaryawan: req.NamaKaryawan}, nil
}

func encodeReadKaryawanByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.ReadKaryawanByKeteranganReq{Keterangan: req.Keterangan}, nil
}

func decodeKaryawanResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadKaryawanByTeleponRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKaryawanByTeleponResp)
	return svc.Karyawan{
		IdKaryawan:   resp.IdKaryawan,
		NamaKaryawan: resp.NamaKaryawan,
		Alamat:       resp.Alamat,
		NoTelepon:    resp.NoTelepon,
		Keterangan:   resp.Keterangan,
		IdJabatan:    resp.IdJabatan,
		IdDepartemen: resp.IdDepartemen,
		IdGolongan:   resp.IdGolongan,
		IdKontrak:    resp.IdKontrak,
		Status:       resp.Status,
	}, nil
}

func decodeReadKaryawanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKaryawanResp)
	var rsp svc.Karyawans

	for _, v := range resp.AllKaryawan {
		itm := svc.Karyawan{
			IdKaryawan:   v.IdKaryawan,
			NamaKaryawan: v.NamaKaryawan,
			Alamat:       v.Alamat,
			NoTelepon:    v.NoTelepon,
			Keterangan:   v.Keterangan,
			IdJabatan:    v.IdJabatan,
			IdDepartemen: v.IdDepartemen,
			IdGolongan:   v.IdGolongan,
			IdKontrak:    v.IdKontrak,
			Status:       v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadKaryawanbyNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKaryawanByNamaResp)
	return svc.Karyawan{
		IdKaryawan:   resp.IdKaryawan,
		NamaKaryawan: resp.NamaKaryawan,
		Alamat:       resp.Alamat,
		NoTelepon:    resp.NoTelepon,
		Keterangan:   resp.Keterangan,
		IdJabatan:    resp.IdJabatan,
		IdDepartemen: resp.IdDepartemen,
		IdGolongan:   resp.IdGolongan,
		IdKontrak:    resp.IdKontrak,
		Status:       resp.Status,
	}, nil
}

func decodeReadKaryawanbyKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKaryawanByKeteranganResp)
	var rsp svc.Karyawans

	for _, v := range resp.AllKeterangan {
		itm := svc.Karyawan{
			IdKaryawan:   v.IdKaryawan,
			NamaKaryawan: v.NamaKaryawan,
			Alamat:       v.Alamat,
			NoTelepon:    v.NoTelepon,
			Keterangan:   v.Keterangan,
			IdJabatan:    v.IdJabatan,
			IdDepartemen: v.IdDepartemen,
			IdGolongan:   v.IdGolongan,
			IdKontrak:    v.IdKontrak,
			Status:       v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddKaryawanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddKaryawan",
		encodeAddKaryawanRequest,
		decodeKaryawanResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddKaryawan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddKaryawan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKaryawanByTeleponEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKaryawanByTelepon",
		encodeReadKaryawanByTeleponRequest,
		decodeReadKaryawanByTeleponRespones,
		pb.ReadKaryawanByTeleponResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKaryawanByTelepon")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKaryawanByTelepon",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKaryawanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKaryawan",
		encodeReadKaryawanRequest,
		decodeReadKaryawanResponse,
		pb.ReadKaryawanResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKaryawan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKaryawan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateKaryawan(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateKaryawan",
		encodeUpdateKaryawanRequest,
		decodeKaryawanResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateKaryawan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateKaryawan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKaryawanByNama(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKaryawanByNama",
		encodeReadKaryawanByNamaRequest,
		decodeReadKaryawanbyNamaRespones,
		pb.ReadKaryawanByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKaryawanByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKaryawanByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKaryawanByKeteranganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKaryawanByKeterangan",
		encodeReadKaryawanByKeteranganRequest,
		decodeReadKaryawanbyKeteranganResponse,
		pb.ReadKaryawanByKeteranganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKaryawanByKeterangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKaryawanByKeterangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
