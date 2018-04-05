package endpoint

import (
	"context"

	scv "projectHotel/hotel/kamar/server"

	pb "projectHotel/hotel/kamar/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcKamarServer struct {
	addKamar    grpctransport.Handler
	readKamar   grpctransport.Handler
	updateKamar grpctransport.Handler
}

func NewGRPCKamarServer(endpoints KamarEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.KamarServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcKamarServer{
		addKamar: grpctransport.NewServer(endpoints.AddKamarEndpoint,
			decodeAddKamarRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddKamar", logger)))...),
		readKamar: grpctransport.NewServer(endpoints.ReadKamarEndpoint,
			decodeReadKamarRequest,
			encodeReadKamarResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKamar", logger)))...),
		updateKamar: grpctransport.NewServer(endpoints.UpdateKamarEndpoint,
			decodeUpdateKamarRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateKamar", logger)))...),
	}
}

func decodeAddKamarRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddKamarReq)
	return scv.Kamar{IdKamar: req.GetIdKamar(), IdTipeKamar: req.GetIdTipeKamar(), IdMenu: req.GetIdMenu(),
		Status: req.GetStatus()}, nil
}

func decodeReadKamarRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateKamarRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateKamarReq)
	return scv.Kamar{IdKamar: req.IdKamar, IdTipeKamar: req.IdTipeKamar, IdMenu: req.IdMenu,
		Status: req.Status}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadKamarResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Kamars)

	rsp := &pb.ReadKamarResp{}

	for _, v := range resp {
		itm := &pb.AddKamarReq{
			IdKamar:     v.IdKamar,
			IdTipeKamar: v.IdTipeKamar,
			IdMenu:      v.IdMenu,
			Status:      v.Status,
		}
		rsp.AllKamar = append(rsp.AllKamar, itm)
	}
	return rsp, nil
}

func (s *grpcKamarServer) AddKamar(ctx oldcontext.Context, kamar *pb.AddKamarReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addKamar.ServeGRPC(ctx, kamar)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcKamarServer) ReadKamar(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadKamarResp, error) {
	_, resp, err := s.readKamar.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKamarResp), nil
}

func (s *grpcKamarServer) UpdateKamar(ctx oldcontext.Context, tam *pb.UpdateKamarReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateKamar.ServeGRPC(ctx, tam)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
