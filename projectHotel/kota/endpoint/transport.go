package endpoint

import (
	"context"

	scv "projectHotel/hotel/kota/server"

	pb "projectHotel/hotel/kota/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcKotaServer struct {
	addKota        grpctransport.Handler
	readKota       grpctransport.Handler
	updateKota     grpctransport.Handler
	readKotaByNama grpctransport.Handler
}

func NewGRPCKotaServer(endpoints KotaEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.KotaServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcKotaServer{
		addKota: grpctransport.NewServer(endpoints.AddKotaEndpoint,
			decodeAddKotaRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddKota", logger)))...),
		readKota: grpctransport.NewServer(endpoints.ReadKotaEndpoint,
			decodeReadKotaRequest,
			encodeReadKotaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKota", logger)))...),
		updateKota: grpctransport.NewServer(endpoints.UpdateKotaEndpoint,
			decodeUpdateKotaRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateKota", logger)))...),
		readKotaByNama: grpctransport.NewServer(endpoints.ReadKotaByNamaEndpoint,
			decodeReadKotaByNamaRequest,
			encodeReadKotaByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKotaByNama", logger)))...),
	}
}

func decodeAddKotaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddKotaReq)
	return scv.Kota{IdKota: req.GetIdKota(), NamaKota: req.GetNamaKota(),
		Status: req.GetStatus()}, nil
}

func decodeReadKotaRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateKotaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateKotaReq)
	return scv.Kota{IdKota: req.IdKota, NamaKota: req.NamaKota,
		Status: req.Status}, nil
}

func decodeReadKotaByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKotaByNamaReq)
	return scv.Kota{NamaKota: req.NamaKota}, nil

}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadKotaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Kotas)

	rsp := &pb.ReadKotaResp{}

	for _, v := range resp {
		itm := &pb.ReadKotaByNamaResp{
			IdKota:   v.IdKota,
			NamaKota: v.NamaKota,
			Status:   v.Status,
		}
		rsp.AllKota = append(rsp.AllKota, itm)
	}
	return rsp, nil
}

func encodeReadKotaByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Kota)
	return &pb.ReadKotaByNamaResp{IdKota: resp.IdKota, NamaKota: resp.NamaKota,
		Status: resp.Status}, nil
}

func (s *grpcKotaServer) AddKota(ctx oldcontext.Context, kota *pb.AddKotaReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addKota.ServeGRPC(ctx, kota)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcKotaServer) ReadKota(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadKotaResp, error) {
	_, resp, err := s.readKota.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKotaResp), nil
}

func (s *grpcKotaServer) UpdateKota(ctx oldcontext.Context, kot *pb.UpdateKotaReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateKota.ServeGRPC(ctx, kot)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcKotaServer) ReadKotaByNama(ctx oldcontext.Context, nama *pb.ReadKotaByNamaReq) (*pb.ReadKotaByNamaResp, error) {
	_, resp, err := s.readKotaByNama.ServeGRPC(ctx, nama)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKotaByNamaResp), nil
}
