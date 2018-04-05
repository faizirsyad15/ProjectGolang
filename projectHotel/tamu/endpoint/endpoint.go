package endpoint

import (
	"context"

	svc "projectHotel/hotel/tamu/server"

	kit "github.com/go-kit/kit/endpoint"
)

type TamuEndpoint struct {
	AddTamuEndpoint           kit.Endpoint
	ReadTamuByTeleponEndpoint kit.Endpoint
	ReadTamuEndpoint          kit.Endpoint
	UpdateTamuEndpoint        kit.Endpoint
	ReadTamuByNamaEndpoint    kit.Endpoint
}

func NewTamuEndpoint(service svc.TamuService) TamuEndpoint {
	addTamuEp := makeAddTamuEndpoint(service)
	readTamuByTeleponEp := makeReadTamuByTeleponEndpoint(service)
	readTamuEp := makeReadTamuEndpoint(service)
	updateTamuEp := makeUpdateTamuEndpoint(service)
	readTamuByNamaEp := makeReadTamuByNamaEndpoint(service)
	return TamuEndpoint{AddTamuEndpoint: addTamuEp,
		ReadTamuByTeleponEndpoint: readTamuByTeleponEp,
		ReadTamuEndpoint:          readTamuEp,
		UpdateTamuEndpoint:        updateTamuEp,
		ReadTamuByNamaEndpoint:    readTamuByNamaEp,
	}
}

func makeAddTamuEndpoint(service svc.TamuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Tamu)
		err := service.AddTamuService(ctx, req)
		return nil, err
	}
}

func makeReadTamuByTeleponEndpoint(service svc.TamuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Tamu)
		result, err := service.ReadTamuByTeleponService(ctx, req.NoTelepon)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadTamuEndpoint(service svc.TamuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadTamuService(ctx)
		return result, err
	}
}

func makeUpdateTamuEndpoint(service svc.TamuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Tamu)
		err := service.UpdateTamuService(ctx, req)
		return nil, err
	}
}

func makeReadTamuByNamaEndpoint(service svc.TamuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Tamu)
		result, err := service.ReadTamuByNamaService(ctx, req.NamaTamu)
		return result, err
	}
}
