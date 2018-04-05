package endpoint

import (
	"context"

	scv "projectHotel/hotel/tamu/server"

	pb "projectHotel/hotel/tamu/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcTamuServer struct {
	addTamu           grpctransport.Handler
	readTamuByTelepon grpctransport.Handler
	readTamu          grpctransport.Handler
	updateTamu        grpctransport.Handler
	readTamuByNama    grpctransport.Handler
}

func NewGRPCTamuServer(endpoints TamuEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.TamuServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcTamuServer{
		addTamu: grpctransport.NewServer(endpoints.AddTamuEndpoint,
			decodeAddTamuRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddTamu", logger)))...),
		readTamuByTelepon: grpctransport.NewServer(endpoints.ReadTamuByTeleponEndpoint,
			decodeReadTamuByTeleponRequest,
			encodeReadTamuByTeleponResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "readTamuByTelepon", logger)))...),
		readTamu: grpctransport.NewServer(endpoints.ReadTamuEndpoint,
			decodeReadTamuRequest,
			encodeReadTamuResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadTamu", logger)))...),
		updateTamu: grpctransport.NewServer(endpoints.UpdateTamuEndpoint,
			decodeUpdateTamuRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateTamu", logger)))...),
		readTamuByNama: grpctransport.NewServer(endpoints.ReadTamuByNamaEndpoint,
			decodeReadTamuByNamaRequest,
			encodeReadTamuByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadTamuByNama", logger)))...),
	}
}

func decodeAddTamuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddTamuReq)
	return scv.Tamu{IdTamu: req.GetIdTamu(), NamaTamu: req.GetNamaTamu(), NoTelepon: req.GetNoTelepon(),
		JenisKelamin: req.GetJenisKelamin(), IdAlamat: req.GetIdAlamat(),
		Status: req.GetStatus()}, nil
}

func decodeReadTamuByTeleponRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadTamuByTeleponReq)
	return scv.Tamu{NoTelepon: req.NoTelepon}, nil
}

func decodeReadTamuRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateTamuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateTamuReq)
	return scv.Tamu{IdTamu: req.IdTamu, NamaTamu: req.NamaTamu, NoTelepon: req.NoTelepon,
		JenisKelamin: req.JenisKelamin, IdAlamat: req.IdAlamat,
		Status: req.Status}, nil
}

func decodeReadTamuByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadTamuByNamaReq)
	return scv.Tamu{NamaTamu: req.NamaTamu}, nil

}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadTamuByTeleponResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Tamu)
	return &pb.ReadTamuByTeleponResp{IdTamu: resp.IdTamu, NamaTamu: resp.NamaTamu, NoTelepon: resp.NoTelepon,
		JenisKelamin: resp.JenisKelamin, IdAlamat: resp.IdAlamat,
		Status: resp.Status}, nil
}

func encodeReadTamuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Tamus)

	rsp := &pb.ReadTamuResp{}

	for _, v := range resp {
		itm := &pb.ReadTamuByTeleponResp{
			IdTamu:       v.IdTamu,
			NamaTamu:     v.NamaTamu,
			NoTelepon:    v.NoTelepon,
			JenisKelamin: v.JenisKelamin,
			IdAlamat:     v.IdAlamat,
			Status:       v.Status,
		}
		rsp.AllTamu = append(rsp.AllTamu, itm)
	}
	return rsp, nil
}

func encodeReadTamuByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Tamu)
	return &pb.ReadTamuByNamaResp{IdTamu: resp.IdTamu, NamaTamu: resp.NamaTamu, NoTelepon: resp.NoTelepon,
		JenisKelamin: resp.JenisKelamin, IdAlamat: resp.IdAlamat,
		Status: resp.Status}, nil
}

func (s *grpcTamuServer) AddTamu(ctx oldcontext.Context, tamu *pb.AddTamuReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addTamu.ServeGRPC(ctx, tamu)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcTamuServer) ReadTamuByTelepon(ctx oldcontext.Context, telepon *pb.ReadTamuByTeleponReq) (*pb.ReadTamuByTeleponResp, error) {
	_, resp, err := s.readTamuByTelepon.ServeGRPC(ctx, telepon)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadTamuByTeleponResp), nil
}

func (s *grpcTamuServer) ReadTamu(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadTamuResp, error) {
	_, resp, err := s.readTamu.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadTamuResp), nil
}

func (s *grpcTamuServer) UpdateTamu(ctx oldcontext.Context, tam *pb.UpdateTamuReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateTamu.ServeGRPC(ctx, tam)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcTamuServer) ReadTamuByNama(ctx oldcontext.Context, nama *pb.ReadTamuByNamaReq) (*pb.ReadTamuByNamaResp, error) {
	_, resp, err := s.readTamuByNama.ServeGRPC(ctx, nama)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadTamuByNamaResp), nil
}
