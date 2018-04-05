package endpoint

import (
	"context"

	svc "projectHotel/hotel/golongan/server"

	kit "github.com/go-kit/kit/endpoint"
)

type GolonganEndpoint struct {
	AddGolonganEndpoint        kit.Endpoint
	ReadGolonganEndpoint       kit.Endpoint
	UpdateGolonganEndpoint     kit.Endpoint
	ReadGolonganByNamaEndpoint kit.Endpoint
}

func NewGolonganEndpoint(service svc.GolonganService) GolonganEndpoint {
	addGolonganEp := makeAddGolonganEndpoint(service)
	readGolonganEp := makeReadGolonganEndpoint(service)
	updateGolonganEp := makeUpdateGolonganEndpoint(service)
	readGolonganByNamaEp := makeReadGolonganByNamaEndpoint(service)
	return GolonganEndpoint{AddGolonganEndpoint: addGolonganEp,
		ReadGolonganEndpoint:       readGolonganEp,
		UpdateGolonganEndpoint:     updateGolonganEp,
		ReadGolonganByNamaEndpoint: readGolonganByNamaEp,
	}
}

func makeAddGolonganEndpoint(service svc.GolonganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Golongan)
		err := service.AddGolonganService(ctx, req)
		return nil, err
	}
}

func makeReadGolonganEndpoint(service svc.GolonganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadGolonganService(ctx)
		return result, err
	}
}

func makeUpdateGolonganEndpoint(service svc.GolonganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Golongan)
		err := service.UpdateGolonganService(ctx, req)
		return nil, err
	}
}

func makeReadGolonganByNamaEndpoint(service svc.GolonganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Golongan)
		result, err := service.ReadGolonganByNamaService(ctx, req.NamaGolongan)
		return result, err
	}
}
