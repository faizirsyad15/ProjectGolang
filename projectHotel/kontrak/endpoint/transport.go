package endpoint

import (
	"context"

	scv "projectHotel/hotel/kontrak/server"

	pb "projectHotel/hotel/kontrak/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcKontrakServer struct {
	addKontrak           grpctransport.Handler
	readKontrakBySelesai grpctransport.Handler
	readKontrak          grpctransport.Handler
	updateKontrak        grpctransport.Handler
	readKontrakByMulai   grpctransport.Handler
}

func NewGRPCKontrakServer(endpoints KontrakEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.KontrakServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcKontrakServer{
		addKontrak: grpctransport.NewServer(endpoints.AddKontrakEndpoint,
			decodeAddKontrakRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddKontrak", logger)))...),
		readKontrakBySelesai: grpctransport.NewServer(endpoints.ReadKontrakBySelesaiEndpoint,
			decodeReadKontrakBySelesaiRequest,
			encodeReadKontrakBySelesaiResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "readKontrakBySelesai", logger)))...),
		readKontrak: grpctransport.NewServer(endpoints.ReadKontrakEndpoint,
			decodeReadKontrakRequest,
			encodeReadKontrakResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKontrak", logger)))...),
		updateKontrak: grpctransport.NewServer(endpoints.UpdateKontrakEndpoint,
			decodeUpdateKontrakRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateKontrak", logger)))...),
		readKontrakByMulai: grpctransport.NewServer(endpoints.ReadKontrakByMulaiEndpoint,
			decodeReadKontrakByMulaiRequest,
			encodeReadKontrakByMulaiResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKontrakByMulai", logger)))...),
	}
}

func decodeAddKontrakRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddKontrakReq)
	return scv.Kontrak{IdKontrak: req.GetIdKontrak(), TanggalMulai: req.GetTanggalMulai(), TanggalSelesai: req.GetTanggalSelesai(),
		Keterangan: req.GetKeterangan(),
		Status:     req.GetStatus()}, nil
}

func decodeReadKontrakBySelesaiRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKontrakBySelesaiReq)
	return scv.Kontrak{TanggalSelesai: req.TanggalSelesai}, nil
}

func decodeReadKontrakRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateKontrakRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateKontrakReq)
	return scv.Kontrak{IdKontrak: req.IdKontrak, TanggalMulai: req.TanggalMulai, TanggalSelesai: req.TanggalSelesai,
		Keterangan: req.Keterangan,
		Status:     req.Status}, nil
}

func decodeReadKontrakByMulaiRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKontrakByMulaiReq)
	return scv.Kontrak{TanggalMulai: req.TanggalMulai}, nil

}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadKontrakBySelesaiResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Kontrak)
	return &pb.ReadKontrakBySelesaiResp{IdKontrak: resp.IdKontrak, TanggalMulai: resp.TanggalMulai, TanggalSelesai: resp.TanggalSelesai,
		Keterangan: resp.Keterangan,
		Status:     resp.Status}, nil
}

func encodeReadKontrakResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Kontraks)

	rsp := &pb.ReadKontrakResp{}

	for _, v := range resp {
		itm := &pb.ReadKontrakBySelesaiResp{
			IdKontrak:      v.IdKontrak,
			TanggalMulai:   v.TanggalMulai,
			TanggalSelesai: v.TanggalSelesai,
			Keterangan:     v.Keterangan,
			Status:         v.Status,
		}
		rsp.AllKontrak = append(rsp.AllKontrak, itm)
	}
	return rsp, nil
}

func encodeReadKontrakByMulaiResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Kontrak)
	return &pb.ReadKontrakByMulaiResp{IdKontrak: resp.IdKontrak, TanggalMulai: resp.TanggalMulai, TanggalSelesai: resp.TanggalSelesai,
		Keterangan: resp.Keterangan,
		Status:     resp.Status}, nil
}

func (s *grpcKontrakServer) AddKontrak(ctx oldcontext.Context, kontrak *pb.AddKontrakReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addKontrak.ServeGRPC(ctx, kontrak)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcKontrakServer) ReadKontrakBySelesai(ctx oldcontext.Context, telepon *pb.ReadKontrakBySelesaiReq) (*pb.ReadKontrakBySelesaiResp, error) {
	_, resp, err := s.readKontrakBySelesai.ServeGRPC(ctx, telepon)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKontrakBySelesaiResp), nil
}

func (s *grpcKontrakServer) ReadKontrak(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadKontrakResp, error) {
	_, resp, err := s.readKontrak.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKontrakResp), nil
}

func (s *grpcKontrakServer) UpdateKontrak(ctx oldcontext.Context, ktk *pb.UpdateKontrakReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateKontrak.ServeGRPC(ctx, ktk)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcKontrakServer) ReadKontrakByMulai(ctx oldcontext.Context, nama *pb.ReadKontrakByMulaiReq) (*pb.ReadKontrakByMulaiResp, error) {
	_, resp, err := s.readKontrakByMulai.ServeGRPC(ctx, nama)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKontrakByMulaiResp), nil
}
