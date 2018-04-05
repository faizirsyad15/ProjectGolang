package endpoint

import (
	"context"

	scv "projectHotel/hotel/karyawan/server"

	pb "projectHotel/hotel/karyawan/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcKaryawanServer struct {
	addKaryawan              grpctransport.Handler
	readKaryawanByTelepon    grpctransport.Handler
	readKaryawan             grpctransport.Handler
	updateKaryawan           grpctransport.Handler
	readKaryawanByNama       grpctransport.Handler
	readKaryawanByKeterangan grpctransport.Handler
}

func NewGRPCKaryawanServer(endpoints KaryawanEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.KaryawanServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcKaryawanServer{
		addKaryawan: grpctransport.NewServer(endpoints.AddKaryawanEndpoint,
			decodeAddKaryawanRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddKaryawan", logger)))...),
		readKaryawanByTelepon: grpctransport.NewServer(endpoints.ReadKaryawanByTeleponEndpoint,
			decodeReadKaryawanByTeleponRequest,
			encodeReadKaryawanByTeleponResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "readKaryawanByTelepon", logger)))...),
		readKaryawan: grpctransport.NewServer(endpoints.ReadKaryawanEndpoint,
			decodeReadKaryawanRequest,
			encodeReadKaryawanResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKaryawan", logger)))...),
		updateKaryawan: grpctransport.NewServer(endpoints.UpdateKaryawanEndpoint,
			decodeUpdateKaryawanRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateKaryawan", logger)))...),
		readKaryawanByNama: grpctransport.NewServer(endpoints.ReadKaryawanByNamaEndpoint,
			decodeReadKaryawanByNamaRequest,
			encodeReadKaryawanByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKaryawanByNama", logger)))...),
		readKaryawanByKeterangan: grpctransport.NewServer(endpoints.ReadKaryawanByKeteranganEndpoint,
			decodeReadKaryawanByKeteranganRequest,
			encodeReadKaryawanByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKaryawanByNama", logger)))...),
	}
}

func decodeAddKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddKaryawanReq)
	return scv.Karyawan{IdKaryawan: req.GetIdKaryawan(), NamaKaryawan: req.GetNamaKaryawan(), Alamat: req.GetAlamat(), NoTelepon: req.GetNoTelepon(),
		Keterangan: req.GetKeterangan(), IdJabatan: req.GetIdJabatan(), IdDepartemen: req.GetIdDepartemen(), IdGolongan: req.GetIdGolongan(), IdKontrak: req.GetIdKontrak(),
		Status: req.GetStatus()}, nil
}

func decodeReadKaryawanByTeleponRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKaryawanByTeleponReq)
	return scv.Karyawan{NoTelepon: req.NoTelepon}, nil
}

func decodeReadKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateKaryawanReq)
	return scv.Karyawan{IdKaryawan: req.IdKaryawan, NamaKaryawan: req.NamaKaryawan, Alamat: req.Alamat, NoTelepon: req.NoTelepon,
		Keterangan: req.Keterangan, IdJabatan: req.IdJabatan, IdDepartemen: req.IdDepartemen, IdGolongan: req.IdGolongan, IdKontrak: req.IdKontrak,
		Status: req.Status}, nil
}

func decodeReadKaryawanByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKaryawanByNamaReq)
	return scv.Karyawan{NamaKaryawan: req.NamaKaryawan}, nil

}

func decodeReadKaryawanByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKaryawanByKeteranganReq)
	return scv.Karyawan{Keterangan: req.Keterangan}, nil

}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadKaryawanByTeleponResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Karyawan)
	return &pb.ReadKaryawanByTeleponResp{IdKaryawan: resp.IdKaryawan, NamaKaryawan: resp.NamaKaryawan, Alamat: resp.Alamat, NoTelepon: resp.NoTelepon,
		Keterangan: resp.Keterangan, IdJabatan: resp.IdJabatan, IdDepartemen: resp.IdDepartemen, IdGolongan: resp.IdGolongan, IdKontrak: resp.IdKontrak,
		Status: resp.Status}, nil
}

func encodeReadKaryawanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Karyawans)

	rsp := &pb.ReadKaryawanResp{}

	for _, v := range resp {
		itm := &pb.ReadKaryawanByTeleponResp{
			IdKaryawan:   v.IdKaryawan,
			NamaKaryawan: v.NamaKaryawan,
			Alamat:       v.Alamat,
			NoTelepon:    v.NoTelepon,
			Keterangan:   v.Keterangan,
			IdJabatan:    v.IdJabatan,
			IdDepartemen: v.IdDepartemen,
			IdGolongan:   v.IdGolongan,
			IdKontrak:    v.IdKontrak,
			Status:       v.Status,
		}
		rsp.AllKaryawan = append(rsp.AllKaryawan, itm)
	}
	return rsp, nil
}

func encodeReadKaryawanByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Karyawan)
	return &pb.ReadKaryawanByNamaResp{IdKaryawan: resp.IdKaryawan, NamaKaryawan: resp.NamaKaryawan, Alamat: resp.Alamat, NoTelepon: resp.NoTelepon,
		Keterangan: resp.Keterangan, IdJabatan: resp.IdJabatan, IdDepartemen: resp.IdDepartemen, IdGolongan: resp.IdGolongan, IdKontrak: resp.IdKontrak,
		Status: resp.Status}, nil
}

func encodeReadKaryawanByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Karyawans)
	rsp := &pb.ReadKaryawanByKeteranganResp{}

	for _, v := range resp {
		itm := &pb.ReadKaryawanByNamaResp{
			IdKaryawan:   v.IdKaryawan,
			NamaKaryawan: v.NamaKaryawan,
			Alamat:       v.Alamat,
			NoTelepon:    v.NoTelepon,
			Keterangan:   v.Keterangan,
			IdJabatan:    v.IdJabatan,
			IdDepartemen: v.IdDepartemen,
			IdGolongan:   v.IdGolongan,
			IdKontrak:    v.IdKontrak,
			Status:       v.Status,
		}
		rsp.AllKeterangan = append(rsp.AllKeterangan, itm)
	}
	return rsp, nil
}

func (s *grpcKaryawanServer) AddKaryawan(ctx oldcontext.Context, karyawan *pb.AddKaryawanReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addKaryawan.ServeGRPC(ctx, karyawan)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcKaryawanServer) ReadKaryawanByTelepon(ctx oldcontext.Context, telepon *pb.ReadKaryawanByTeleponReq) (*pb.ReadKaryawanByTeleponResp, error) {
	_, resp, err := s.readKaryawanByTelepon.ServeGRPC(ctx, telepon)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKaryawanByTeleponResp), nil
}

func (s *grpcKaryawanServer) ReadKaryawan(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadKaryawanResp, error) {
	_, resp, err := s.readKaryawan.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKaryawanResp), nil
}

func (s *grpcKaryawanServer) UpdateKaryawan(ctx oldcontext.Context, kar *pb.UpdateKaryawanReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateKaryawan.ServeGRPC(ctx, kar)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcKaryawanServer) ReadKaryawanByNama(ctx oldcontext.Context, nama *pb.ReadKaryawanByNamaReq) (*pb.ReadKaryawanByNamaResp, error) {
	_, resp, err := s.readKaryawanByNama.ServeGRPC(ctx, nama)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKaryawanByNamaResp), nil
}

func (s *grpcKaryawanServer) ReadKaryawanByKeterangan(ctx oldcontext.Context, keterangan *pb.ReadKaryawanByKeteranganReq) (*pb.ReadKaryawanByKeteranganResp, error) {
	_, resp, err := s.readKaryawanByKeterangan.ServeGRPC(ctx, keterangan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKaryawanByKeteranganResp), nil
}
