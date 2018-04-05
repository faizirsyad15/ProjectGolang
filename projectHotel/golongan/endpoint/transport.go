package endpoint

import (
	"context"

	scv "projectHotel/hotel/golongan/server"

	pb "projectHotel/hotel/golongan/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcGolonganServer struct {
	addGolongan        grpctransport.Handler
	readGolongan       grpctransport.Handler
	updateGolongan     grpctransport.Handler
	readGolonganByNama grpctransport.Handler
}

func NewGRPCGolonganServer(endpoints GolonganEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.GolonganServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcGolonganServer{
		addGolongan: grpctransport.NewServer(endpoints.AddGolonganEndpoint,
			decodeAddGolonganRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddGolongan", logger)))...),
		readGolongan: grpctransport.NewServer(endpoints.ReadGolonganEndpoint,
			decodeReadGolonganRequest,
			encodeReadGolonganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadGolongan", logger)))...),
		updateGolongan: grpctransport.NewServer(endpoints.UpdateGolonganEndpoint,
			decodeUpdateGolonganRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateGolongan", logger)))...),
		readGolonganByNama: grpctransport.NewServer(endpoints.ReadGolonganByNamaEndpoint,
			decodeReadGolonganByNamaRequest,
			encodeReadGolonganByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadGolonganByNama", logger)))...),
	}
}

func decodeAddGolonganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddGolonganReq)
	return scv.Golongan{IdGolongan: req.GetIdGolongan(), NamaGolongan: req.GetNamaGolongan(),
		Status: req.GetStatus()}, nil
}

func decodeReadGolonganRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateGolonganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateGolonganReq)
	return scv.Golongan{IdGolongan: req.IdGolongan, NamaGolongan: req.NamaGolongan,
		Status: req.Status}, nil
}

func decodeReadGolonganByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadGolonganByNamaReq)
	return scv.Golongan{NamaGolongan: req.NamaGolongan}, nil

}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadGolonganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Golongans)

	rsp := &pb.ReadGolonganResp{}

	for _, v := range resp {
		itm := &pb.ReadGolonganByNamaResp{
			IdGolongan:   v.IdGolongan,
			NamaGolongan: v.NamaGolongan,
			Status:       v.Status,
		}
		rsp.AllGolongan = append(rsp.AllGolongan, itm)
	}
	return rsp, nil
}

func encodeReadGolonganByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Golongan)
	return &pb.ReadGolonganByNamaResp{IdGolongan: resp.IdGolongan, NamaGolongan: resp.NamaGolongan,
		Status: resp.Status}, nil
}

func (s *grpcGolonganServer) AddGolongan(ctx oldcontext.Context, golongan *pb.AddGolonganReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addGolongan.ServeGRPC(ctx, golongan)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcGolonganServer) ReadGolongan(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadGolonganResp, error) {
	_, resp, err := s.readGolongan.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadGolonganResp), nil
}

func (s *grpcGolonganServer) UpdateGolongan(ctx oldcontext.Context, gol *pb.UpdateGolonganReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateGolongan.ServeGRPC(ctx, gol)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcGolonganServer) ReadGolonganByNama(ctx oldcontext.Context, nama *pb.ReadGolonganByNamaReq) (*pb.ReadGolonganByNamaResp, error) {
	_, resp, err := s.readGolonganByNama.ServeGRPC(ctx, nama)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadGolonganByNamaResp), nil
}
