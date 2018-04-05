package endpoint

import (
	"context"

	svc "projectHotel/hotel/negara/server"

	kit "github.com/go-kit/kit/endpoint"
)

type NegaraEndpoint struct {
	AddNegaraEndpoint        kit.Endpoint
	ReadNegaraEndpoint       kit.Endpoint
	UpdateNegaraEndpoint     kit.Endpoint
	ReadNegaraByNamaEndpoint kit.Endpoint
}

func NewNegaraEndpoint(service svc.NegaraService) NegaraEndpoint {
	addNegaraEp := makeAddNegaraEndpoint(service)
	readNegaraEp := makeReadNegaraEndpoint(service)
	updateNegaraEp := makeUpdateNegaraEndpoint(service)
	readNegaraByNamaEp := makeReadNegaraByNamaEndpoint(service)
	return NegaraEndpoint{AddNegaraEndpoint: addNegaraEp,
		ReadNegaraEndpoint:       readNegaraEp,
		UpdateNegaraEndpoint:     updateNegaraEp,
		ReadNegaraByNamaEndpoint: readNegaraByNamaEp,
	}
}

func makeAddNegaraEndpoint(service svc.NegaraService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Negara)
		err := service.AddNegaraService(ctx, req)
		return nil, err
	}
}

func makeReadNegaraEndpoint(service svc.NegaraService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadNegaraService(ctx)
		return result, err
	}
}

func makeUpdateNegaraEndpoint(service svc.NegaraService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Negara)
		err := service.UpdateNegaraService(ctx, req)
		return nil, err
	}
}

func makeReadNegaraByNamaEndpoint(service svc.NegaraService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Negara)
		result, err := service.ReadNegaraByNamaService(ctx, req.NamaNegara)
		return result, err
	}
}
