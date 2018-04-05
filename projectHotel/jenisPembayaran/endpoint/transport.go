package endpoint

import (
	"context"

	scv "projectHotel/hotel/jenisPembayaran/server"

	pb "projectHotel/hotel/jenisPembayaran/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcJenisPembayaranServer struct {
	addJenisPembayaran          grpctransport.Handler
	readJenisPembayaran         grpctransport.Handler
	updateJenisPembayaran       grpctransport.Handler
	readJenisPembayaranByMetode grpctransport.Handler
}

func NewGRPCJenisPembayaranServer(endpoints JenisPembayaranEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.JenisPembayaranServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcJenisPembayaranServer{
		addJenisPembayaran: grpctransport.NewServer(endpoints.AddJenisPembayaranEndpoint,
			decodeAddJenisPembayaranRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddJenisPembayaran", logger)))...),
		readJenisPembayaran: grpctransport.NewServer(endpoints.ReadJenisPembayaranEndpoint,
			decodeReadJenisPembayaranRequest,
			encodeReadJenisPembayaranResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadJenisPembayaran", logger)))...),
		updateJenisPembayaran: grpctransport.NewServer(endpoints.UpdateJenisPembayaranEndpoint,
			decodeUpdateJenisPembayaranRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateJenisPembayaran", logger)))...),
		readJenisPembayaranByMetode: grpctransport.NewServer(endpoints.ReadJenisPembayaranByMetodeEndpoint,
			decodeReadJenisPembayaranByMetodeRequest,
			encodeReadJenisPembayaranByMetodeResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadJenisPembayaranByMetode", logger)))...),
	}
}

func decodeAddJenisPembayaranRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddJenisPembayaranReq)
	return scv.JenisPembayaran{IdJenisPembayaran: req.GetIdJenisPembayaran(), MetodePembayaran: req.GetMetodePembayaran(),
		Status: req.GetStatus()}, nil
}

func decodeReadJenisPembayaranRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateJenisPembayaranRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateJenisPembayaranReq)
	return scv.JenisPembayaran{IdJenisPembayaran: req.IdJenisPembayaran, MetodePembayaran: req.MetodePembayaran,
		Status: req.Status}, nil
}

func decodeReadJenisPembayaranByMetodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadJenisPembayaranByMetodeReq)
	return scv.JenisPembayaran{MetodePembayaran: req.MetodePembayaran}, nil

}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadJenisPembayaranResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.JenisPembayarans)

	rsp := &pb.ReadJenisPembayaranResp{}

	for _, v := range resp {
		itm := &pb.ReadJenisPembayaranByMetodeResp{
			IdJenisPembayaran: v.IdJenisPembayaran,
			MetodePembayaran:  v.MetodePembayaran,
			Status:            v.Status,
		}
		rsp.AllJenisPembayaran = append(rsp.AllJenisPembayaran, itm)
	}
	return rsp, nil
}

func encodeReadJenisPembayaranByMetodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.JenisPembayaran)
	return &pb.ReadJenisPembayaranByMetodeResp{IdJenisPembayaran: resp.IdJenisPembayaran, MetodePembayaran: resp.MetodePembayaran,
		Status: resp.Status}, nil
}

func (s *grpcJenisPembayaranServer) AddJenisPembayaran(ctx oldcontext.Context, jenisPembayaran *pb.AddJenisPembayaranReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addJenisPembayaran.ServeGRPC(ctx, jenisPembayaran)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcJenisPembayaranServer) ReadJenisPembayaran(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadJenisPembayaranResp, error) {
	_, resp, err := s.readJenisPembayaran.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadJenisPembayaranResp), nil
}

func (s *grpcJenisPembayaranServer) UpdateJenisPembayaran(ctx oldcontext.Context, jen *pb.UpdateJenisPembayaranReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateJenisPembayaran.ServeGRPC(ctx, jen)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcJenisPembayaranServer) ReadJenisPembayaranByMetode(ctx oldcontext.Context, nama *pb.ReadJenisPembayaranByMetodeReq) (*pb.ReadJenisPembayaranByMetodeResp, error) {
	_, resp, err := s.readJenisPembayaranByMetode.ServeGRPC(ctx, nama)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadJenisPembayaranByMetodeResp), nil
}
