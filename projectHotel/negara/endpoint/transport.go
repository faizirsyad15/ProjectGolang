package endpoint

import (
	"context"

	scv "projectHotel/hotel/negara/server"

	pb "projectHotel/hotel/negara/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcNegaraServer struct {
	addNegara        grpctransport.Handler
	readNegara       grpctransport.Handler
	updateNegara     grpctransport.Handler
	readNegaraByNama grpctransport.Handler
}

func NewGRPCNegaraServer(endpoints NegaraEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.NegaraServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcNegaraServer{
		addNegara: grpctransport.NewServer(endpoints.AddNegaraEndpoint,
			decodeAddNegaraRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddNegara", logger)))...),
		readNegara: grpctransport.NewServer(endpoints.ReadNegaraEndpoint,
			decodeReadNegaraRequest,
			encodeReadNegaraResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadNegara", logger)))...),
		updateNegara: grpctransport.NewServer(endpoints.UpdateNegaraEndpoint,
			decodeUpdateNegaraRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateNegara", logger)))...),
		readNegaraByNama: grpctransport.NewServer(endpoints.ReadNegaraByNamaEndpoint,
			decodeReadNegaraByNamaRequest,
			encodeReadNegaraByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadNegaraByNama", logger)))...),
	}
}

func decodeAddNegaraRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddNegaraReq)
	return scv.Negara{IdNegara: req.GetIdNegara(), NamaNegara: req.GetNamaNegara(),
		Status: req.GetStatus()}, nil
}

func decodeReadNegaraRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateNegaraRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateNegaraReq)
	return scv.Negara{IdNegara: req.IdNegara, NamaNegara: req.NamaNegara,
		Status: req.Status}, nil
}

func decodeReadNegaraByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadNegaraByNamaReq)
	return scv.Negara{NamaNegara: req.NamaNegara}, nil

}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadNegaraResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Negaras)

	rsp := &pb.ReadNegaraResp{}

	for _, v := range resp {
		itm := &pb.ReadNegaraByNamaResp{
			IdNegara:   v.IdNegara,
			NamaNegara: v.NamaNegara,
			Status:     v.Status,
		}
		rsp.AllNegara = append(rsp.AllNegara, itm)
	}
	return rsp, nil
}

func encodeReadNegaraByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Negara)
	return &pb.ReadNegaraByNamaResp{IdNegara: resp.IdNegara, NamaNegara: resp.NamaNegara,
		Status: resp.Status}, nil
}

func (s *grpcNegaraServer) AddNegara(ctx oldcontext.Context, negara *pb.AddNegaraReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addNegara.ServeGRPC(ctx, negara)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcNegaraServer) ReadNegara(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadNegaraResp, error) {
	_, resp, err := s.readNegara.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadNegaraResp), nil
}

func (s *grpcNegaraServer) UpdateNegara(ctx oldcontext.Context, neg *pb.UpdateNegaraReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateNegara.ServeGRPC(ctx, neg)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcNegaraServer) ReadNegaraByNama(ctx oldcontext.Context, nama *pb.ReadNegaraByNamaReq) (*pb.ReadNegaraByNamaResp, error) {
	_, resp, err := s.readNegaraByNama.ServeGRPC(ctx, nama)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadNegaraByNamaResp), nil
}
