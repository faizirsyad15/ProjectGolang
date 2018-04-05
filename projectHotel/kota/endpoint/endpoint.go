package endpoint

import (
	"context"

	svc "projectHotel/hotel/kota/server"

	kit "github.com/go-kit/kit/endpoint"
)

type KotaEndpoint struct {
	AddKotaEndpoint        kit.Endpoint
	ReadKotaEndpoint       kit.Endpoint
	UpdateKotaEndpoint     kit.Endpoint
	ReadKotaByNamaEndpoint kit.Endpoint
}

func NewKotaEndpoint(service svc.KotaService) KotaEndpoint {
	addKotaEp := makeAddKotaEndpoint(service)
	readKotaEp := makeReadKotaEndpoint(service)
	updateKotaEp := makeUpdateKotaEndpoint(service)
	readKotaByNamaEp := makeReadKotaByNamaEndpoint(service)
	return KotaEndpoint{AddKotaEndpoint: addKotaEp,
		ReadKotaEndpoint:       readKotaEp,
		UpdateKotaEndpoint:     updateKotaEp,
		ReadKotaByNamaEndpoint: readKotaByNamaEp,
	}
}

func makeAddKotaEndpoint(service svc.KotaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kota)
		err := service.AddKotaService(ctx, req)
		return nil, err
	}
}

func makeReadKotaEndpoint(service svc.KotaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadKotaService(ctx)
		return result, err
	}
}

func makeUpdateKotaEndpoint(service svc.KotaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kota)
		err := service.UpdateKotaService(ctx, req)
		return nil, err
	}
}

func makeReadKotaByNamaEndpoint(service svc.KotaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kota)
		result, err := service.ReadKotaByNamaService(ctx, req.NamaKota)
		return result, err
	}
}
