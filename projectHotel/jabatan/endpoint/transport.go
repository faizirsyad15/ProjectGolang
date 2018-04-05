package endpoint

import (
	"context"

	scv "projectHotel/hotel/jabatan/server"

	pb "projectHotel/hotel/jabatan/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcJabatanServer struct {
	addJabatan        grpctransport.Handler
	readJabatan       grpctransport.Handler
	updateJabatan     grpctransport.Handler
	readJabatanByNama grpctransport.Handler
}

func NewGRPCJabatanServer(endpoints JabatanEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.JabatanServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcJabatanServer{
		addJabatan: grpctransport.NewServer(endpoints.AddJabatanEndpoint,
			decodeAddJabatanRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddJabatan", logger)))...),
		readJabatan: grpctransport.NewServer(endpoints.ReadJabatanEndpoint,
			decodeReadJabatanRequest,
			encodeReadJabatanResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadJabatan", logger)))...),
		updateJabatan: grpctransport.NewServer(endpoints.UpdateJabatanEndpoint,
			decodeUpdateJabatanRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateJabatan", logger)))...),
		readJabatanByNama: grpctransport.NewServer(endpoints.ReadJabatanByNamaEndpoint,
			decodeReadJabatanByNamaRequest,
			encodeReadJabatanByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadJabatanByNama", logger)))...),
	}
}

func decodeAddJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddJabatanReq)
	return scv.Jabatan{IdJabatan: req.GetIdJabatan(), NamaJabatan: req.GetNamaJabatan(),
		Status: req.GetStatus()}, nil
}

func decodeReadJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateJabatanReq)
	return scv.Jabatan{IdJabatan: req.IdJabatan, NamaJabatan: req.NamaJabatan,
		Status: req.Status}, nil
}

func decodeReadJabatanByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadJabatanByNamaReq)
	return scv.Jabatan{NamaJabatan: req.NamaJabatan}, nil

}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadJabatanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Jabatans)

	rsp := &pb.ReadJabatanResp{}

	for _, v := range resp {
		itm := &pb.ReadJabatanByNamaResp{
			IdJabatan:   v.IdJabatan,
			NamaJabatan: v.NamaJabatan,
			Status:      v.Status,
		}
		rsp.AllJabatan = append(rsp.AllJabatan, itm)
	}
	return rsp, nil
}

func encodeReadJabatanByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Jabatan)
	return &pb.ReadJabatanByNamaResp{IdJabatan: resp.IdJabatan, NamaJabatan: resp.NamaJabatan,
		Status: resp.Status}, nil
}

func (s *grpcJabatanServer) AddJabatan(ctx oldcontext.Context, jabatan *pb.AddJabatanReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addJabatan.ServeGRPC(ctx, jabatan)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcJabatanServer) ReadJabatan(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadJabatanResp, error) {
	_, resp, err := s.readJabatan.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadJabatanResp), nil
}

func (s *grpcJabatanServer) UpdateJabatan(ctx oldcontext.Context, kot *pb.UpdateJabatanReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateJabatan.ServeGRPC(ctx, kot)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcJabatanServer) ReadJabatanByNama(ctx oldcontext.Context, nama *pb.ReadJabatanByNamaReq) (*pb.ReadJabatanByNamaResp, error) {
	_, resp, err := s.readJabatanByNama.ServeGRPC(ctx, nama)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadJabatanByNamaResp), nil
}
