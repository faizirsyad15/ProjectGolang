package endpoint

import (
	"context"

	scv "projectHotel/hotel/tipeKamar/server"

	pb "projectHotel/hotel/tipeKamar/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcTipeKamarServer struct {
	addTipeKamar         grpctransport.Handler
	readTipeKamarByHarga grpctransport.Handler
	readTipeKamar        grpctransport.Handler
	updateTipeKamar      grpctransport.Handler
	readTipeKamarByNama  grpctransport.Handler
}

func NewGRPCTipeKamarServer(endpoints TipeKamarEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.TipeKamarServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcTipeKamarServer{
		addTipeKamar: grpctransport.NewServer(endpoints.AddTipeKamarEndpoint,
			decodeAddTipeKamarRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddTipeKamar", logger)))...),
		readTipeKamarByHarga: grpctransport.NewServer(endpoints.ReadTipeKamarByHargaEndpoint,
			decodeReadTipeKamarByHargaRequest,
			encodeReadTipeKamarByHargaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "readTipeKamarByHarga", logger)))...),
		readTipeKamar: grpctransport.NewServer(endpoints.ReadTipeKamarEndpoint,
			decodeReadTipeKamarRequest,
			encodeReadTipeKamarResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadTipeKamar", logger)))...),
		updateTipeKamar: grpctransport.NewServer(endpoints.UpdateTipeKamarEndpoint,
			decodeUpdateTipeKamarRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateTipeKamar", logger)))...),
		readTipeKamarByNama: grpctransport.NewServer(endpoints.ReadTipeKamarByNamaEndpoint,
			decodeReadTipeKamarByNamaRequest,
			encodeReadTipeKamarByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadTipeKamarByNama", logger)))...),
	}
}

func decodeAddTipeKamarRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddTipeKamarReq)
	return scv.TipeKamar{IdTipeKamar: req.GetIdTipeKamar(), NamaTipeKamar: req.GetNamaTipeKamar(), HargaKamar: req.GetHargaKamar(),
		Status: req.GetStatus()}, nil
}

func decodeReadTipeKamarByHargaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadTipeKamarByHargaReq)
	return scv.TipeKamar{HargaKamar: req.HargaKamar}, nil
}

func decodeReadTipeKamarRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateTipeKamarRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateTipeKamarReq)
	return scv.TipeKamar{IdTipeKamar: req.IdTipeKamar, NamaTipeKamar: req.NamaTipeKamar, HargaKamar: req.HargaKamar,
		Status: req.Status}, nil
}

func decodeReadTipeKamarByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadTipeKamarByNamaReq)
	return scv.TipeKamar{NamaTipeKamar: req.NamaTipeKamar}, nil

}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadTipeKamarByHargaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.TipeKamar)
	return &pb.ReadTipeKamarByHargaResp{IdTipeKamar: resp.IdTipeKamar, NamaTipeKamar: resp.NamaTipeKamar, HargaKamar: resp.HargaKamar,
		Status: resp.Status}, nil
}

func encodeReadTipeKamarResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.TipeKamars)

	rsp := &pb.ReadTipeKamarResp{}

	for _, v := range resp {
		itm := &pb.ReadTipeKamarByHargaResp{
			IdTipeKamar:   v.IdTipeKamar,
			NamaTipeKamar: v.NamaTipeKamar,
			HargaKamar:    v.HargaKamar,
			Status:        v.Status,
		}
		rsp.AllTipeKamar = append(rsp.AllTipeKamar, itm)
	}
	return rsp, nil
}

func encodeReadTipeKamarByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.TipeKamar)
	return &pb.ReadTipeKamarByNamaResp{IdTipeKamar: resp.IdTipeKamar, NamaTipeKamar: resp.NamaTipeKamar, HargaKamar: resp.HargaKamar,
		Status: resp.Status}, nil
}

func (s *grpcTipeKamarServer) AddTipeKamar(ctx oldcontext.Context, tipeKamar *pb.AddTipeKamarReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addTipeKamar.ServeGRPC(ctx, tipeKamar)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcTipeKamarServer) ReadTipeKamarByHarga(ctx oldcontext.Context, telepon *pb.ReadTipeKamarByHargaReq) (*pb.ReadTipeKamarByHargaResp, error) {
	_, resp, err := s.readTipeKamarByHarga.ServeGRPC(ctx, telepon)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadTipeKamarByHargaResp), nil
}

func (s *grpcTipeKamarServer) ReadTipeKamar(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadTipeKamarResp, error) {
	_, resp, err := s.readTipeKamar.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadTipeKamarResp), nil
}

func (s *grpcTipeKamarServer) UpdateTipeKamar(ctx oldcontext.Context, tam *pb.UpdateTipeKamarReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateTipeKamar.ServeGRPC(ctx, tam)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcTipeKamarServer) ReadTipeKamarByNama(ctx oldcontext.Context, nama *pb.ReadTipeKamarByNamaReq) (*pb.ReadTipeKamarByNamaResp, error) {
	_, resp, err := s.readTipeKamarByNama.ServeGRPC(ctx, nama)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadTipeKamarByNamaResp), nil
}
