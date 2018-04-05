package endpoint

import (
	"context"

	svc "projectHotel/hotel/karyawan/server"

	kit "github.com/go-kit/kit/endpoint"
)

type KaryawanEndpoint struct {
	AddKaryawanEndpoint              kit.Endpoint
	ReadKaryawanByTeleponEndpoint    kit.Endpoint
	ReadKaryawanEndpoint             kit.Endpoint
	UpdateKaryawanEndpoint           kit.Endpoint
	ReadKaryawanByNamaEndpoint       kit.Endpoint
	ReadKaryawanByKeteranganEndpoint kit.Endpoint
}

func NewKaryawanEndpoint(service svc.KaryawanService) KaryawanEndpoint {
	addKaryawanEp := makeAddKaryawanEndpoint(service)
	readKaryawanByTeleponEp := makeReadKaryawanByTeleponEndpoint(service)
	readKaryawanEp := makeReadKaryawanEndpoint(service)
	updateKaryawanEp := makeUpdateKaryawanEndpoint(service)
	readKaryawanByNamaEp := makeReadKaryawanByNamaEndpoint(service)
	readKaryawanByKeteranganEp := makeReadKaryawanByKeteranganEndpoint(service)
	return KaryawanEndpoint{AddKaryawanEndpoint: addKaryawanEp,
		ReadKaryawanByTeleponEndpoint:    readKaryawanByTeleponEp,
		ReadKaryawanEndpoint:             readKaryawanEp,
		UpdateKaryawanEndpoint:           updateKaryawanEp,
		ReadKaryawanByNamaEndpoint:       readKaryawanByNamaEp,
		ReadKaryawanByKeteranganEndpoint: readKaryawanByKeteranganEp,
	}
}

func makeAddKaryawanEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		err := service.AddKaryawanService(ctx, req)
		return nil, err
	}
}

func makeReadKaryawanByTeleponEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		result, err := service.ReadKaryawanByTeleponService(ctx, req.NoTelepon)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadKaryawanEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadKaryawanService(ctx)
		return result, err
	}
}

func makeUpdateKaryawanEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		err := service.UpdateKaryawanService(ctx, req)
		return nil, err
	}
}

func makeReadKaryawanByNamaEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		result, err := service.ReadKaryawanByNamaService(ctx, req.NamaKaryawan)
		return result, err
	}
}

func makeReadKaryawanByKeteranganEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		result, err := service.ReadKaryawanByKeteranganService(ctx, req.Keterangan)
		return result, err
	}
}
