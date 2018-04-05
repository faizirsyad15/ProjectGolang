package endpoint

import (
	"context"

	svc "projectHotel/hotel/kamar/server"

	kit "github.com/go-kit/kit/endpoint"
)

type KamarEndpoint struct {
	AddKamarEndpoint    kit.Endpoint
	ReadKamarEndpoint   kit.Endpoint
	UpdateKamarEndpoint kit.Endpoint
}

func NewKamarEndpoint(service svc.KamarService) KamarEndpoint {
	addKamarEp := makeAddKamarEndpoint(service)
	readKamarEp := makeReadKamarEndpoint(service)
	updateKamarEp := makeUpdateKamarEndpoint(service)
	return KamarEndpoint{AddKamarEndpoint: addKamarEp,
		ReadKamarEndpoint:   readKamarEp,
		UpdateKamarEndpoint: updateKamarEp,
	}
}

func makeAddKamarEndpoint(service svc.KamarService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kamar)
		err := service.AddKamarService(ctx, req)
		return nil, err
	}
}

func makeReadKamarEndpoint(service svc.KamarService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadKamarService(ctx)
		return result, err
	}
}

func makeUpdateKamarEndpoint(service svc.KamarService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kamar)
		err := service.UpdateKamarService(ctx, req)
		return nil, err
	}
}
