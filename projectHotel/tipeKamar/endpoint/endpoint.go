package endpoint

import (
	"context"

	svc "projectHotel/hotel/tipeKamar/server"

	kit "github.com/go-kit/kit/endpoint"
)

type TipeKamarEndpoint struct {
	AddTipeKamarEndpoint         kit.Endpoint
	ReadTipeKamarByHargaEndpoint kit.Endpoint
	ReadTipeKamarEndpoint        kit.Endpoint
	UpdateTipeKamarEndpoint      kit.Endpoint
	ReadTipeKamarByNamaEndpoint  kit.Endpoint
}

func NewTipeKamarEndpoint(service svc.TipeKamarService) TipeKamarEndpoint {
	addTipeKamarEp := makeAddTipeKamarEndpoint(service)
	readTipeKamarByHargaEp := makeReadTipeKamarByHargaEndpoint(service)
	readTipeKamarEp := makeReadTipeKamarEndpoint(service)
	updateTipeKamarEp := makeUpdateTipeKamarEndpoint(service)
	readTipeKamarByNamaEp := makeReadTipeKamarByNamaEndpoint(service)
	return TipeKamarEndpoint{AddTipeKamarEndpoint: addTipeKamarEp,
		ReadTipeKamarByHargaEndpoint: readTipeKamarByHargaEp,
		ReadTipeKamarEndpoint:        readTipeKamarEp,
		UpdateTipeKamarEndpoint:      updateTipeKamarEp,
		ReadTipeKamarByNamaEndpoint:  readTipeKamarByNamaEp,
	}
}

func makeAddTipeKamarEndpoint(service svc.TipeKamarService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.TipeKamar)
		err := service.AddTipeKamarService(ctx, req)
		return nil, err
	}
}

func makeReadTipeKamarByHargaEndpoint(service svc.TipeKamarService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.TipeKamar)
		result, err := service.ReadTipeKamarByHargaService(ctx, req.HargaKamar)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadTipeKamarEndpoint(service svc.TipeKamarService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadTipeKamarService(ctx)
		return result, err
	}
}

func makeUpdateTipeKamarEndpoint(service svc.TipeKamarService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.TipeKamar)
		err := service.UpdateTipeKamarService(ctx, req)
		return nil, err
	}
}

func makeReadTipeKamarByNamaEndpoint(service svc.TipeKamarService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.TipeKamar)
		result, err := service.ReadTipeKamarByNamaService(ctx, req.NamaTipeKamar)
		return result, err
	}
}
