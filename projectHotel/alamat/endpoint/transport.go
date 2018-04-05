package endpoint

import (
	"context"

	scv "projectHotel/hotel/alamat/server"

	pb "projectHotel/hotel/alamat/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcAlamatServer struct {
	addAlamat           grpctransport.Handler
	readAlamatByNoRumah grpctransport.Handler
	readAlamat          grpctransport.Handler
	updateAlamat        grpctransport.Handler
}

func NewGRPCAlamatServer(endpoints AlamatEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.AlamatServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcAlamatServer{
		addAlamat: grpctransport.NewServer(endpoints.AddAlamatEndpoint,
			decodeAddAlamatRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddAlamat", logger)))...),
		readAlamatByNoRumah: grpctransport.NewServer(endpoints.ReadAlamatByNoRumahEndpoint,
			decodeReadAlamatByNoRumahRequest,
			encodeReadAlamatByNoRumahResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "readAlamatByNoRumah", logger)))...),
		readAlamat: grpctransport.NewServer(endpoints.ReadAlamatEndpoint,
			decodeReadAlamatRequest,
			encodeReadAlamatResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadAlamat", logger)))...),
		updateAlamat: grpctransport.NewServer(endpoints.UpdateAlamatEndpoint,
			decodeUpdateAlamatRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateAlamat", logger)))...),
	}
}

func decodeAddAlamatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddAlamatReq)
	return scv.Alamat{IdAlamat: req.GetIdAlamat(), AlamatRumah: req.GetAlamatRumah(), RtRw: req.GetRtRw(),
		NoRumah: req.GetNoRumah(), IdKota: req.GetIdKota(), IdNegara: req.GetIdNegara(),
		Status: req.GetStatus()}, nil
}

func decodeReadAlamatByNoRumahRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadAlamatByNoRumahReq)
	return scv.Alamat{NoRumah: req.NoRumah}, nil
}

func decodeReadAlamatRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateAlamatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateAlamatReq)
	return scv.Alamat{IdAlamat: req.IdAlamat, AlamatRumah: req.AlamatRumah, RtRw: req.RtRw,
		NoRumah: req.NoRumah, IdKota: req.IdKota, IdNegara: req.IdNegara,
		Status: req.Status}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadAlamatByNoRumahResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Alamat)
	return &pb.ReadAlamatByNoRumahResp{IdAlamat: resp.IdAlamat, AlamatRumah: resp.AlamatRumah, RtRw: resp.RtRw,
		NoRumah: resp.NoRumah, IdKota: resp.IdKota, IdNegara: resp.IdNegara,
		Status: resp.Status}, nil
}

func encodeReadAlamatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Alamats)

	rsp := &pb.ReadAlamatResp{}

	for _, v := range resp {
		itm := &pb.ReadAlamatByNoRumahResp{
			IdAlamat:    v.IdAlamat,
			AlamatRumah: v.AlamatRumah,
			RtRw:        v.RtRw,
			NoRumah:     v.NoRumah,
			IdKota:      v.IdKota,
			IdNegara:    v.IdNegara,
			Status:      v.Status,
		}
		rsp.AllAlamat = append(rsp.AllAlamat, itm)
	}
	return rsp, nil
}

func (s *grpcAlamatServer) AddAlamat(ctx oldcontext.Context, alamat *pb.AddAlamatReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addAlamat.ServeGRPC(ctx, alamat)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcAlamatServer) ReadAlamatByNoRumah(ctx oldcontext.Context, norumah *pb.ReadAlamatByNoRumahReq) (*pb.ReadAlamatByNoRumahResp, error) {
	_, resp, err := s.readAlamatByNoRumah.ServeGRPC(ctx, norumah)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadAlamatByNoRumahResp), nil
}

func (s *grpcAlamatServer) ReadAlamat(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadAlamatResp, error) {
	_, resp, err := s.readAlamat.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadAlamatResp), nil
}

func (s *grpcAlamatServer) UpdateAlamat(ctx oldcontext.Context, amt *pb.UpdateAlamatReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateAlamat.ServeGRPC(ctx, amt)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
