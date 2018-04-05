package endpoint

import (
	"context"

	svc "projectHotel/hotel/alamat/server"

	kit "github.com/go-kit/kit/endpoint"
)

type AlamatEndpoint struct {
	AddAlamatEndpoint           kit.Endpoint
	ReadAlamatByNoRumahEndpoint kit.Endpoint
	ReadAlamatEndpoint          kit.Endpoint
	UpdateAlamatEndpoint        kit.Endpoint
}

func NewAlamatEndpoint(service svc.AlamatService) AlamatEndpoint {
	addAlamatEp := makeAddAlamatEndpoint(service)
	readAlamatByNoRumahEp := makeReadAlamatByNoRumahEndpoint(service)
	readAlamatEp := makeReadAlamatEndpoint(service)
	updateAlamatEp := makeUpdateAlamatEndpoint(service)
	return AlamatEndpoint{AddAlamatEndpoint: addAlamatEp,
		ReadAlamatByNoRumahEndpoint: readAlamatByNoRumahEp,
		ReadAlamatEndpoint:          readAlamatEp,
		UpdateAlamatEndpoint:        updateAlamatEp,
	}
}

func makeAddAlamatEndpoint(service svc.AlamatService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Alamat)
		err := service.AddAlamatService(ctx, req)
		return nil, err
	}
}

func makeReadAlamatByNoRumahEndpoint(service svc.AlamatService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Alamat)
		result, err := service.ReadAlamatByNoRumahService(ctx, req.NoRumah)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadAlamatEndpoint(service svc.AlamatService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadAlamatService(ctx)
		return result, err
	}
}

func makeUpdateAlamatEndpoint(service svc.AlamatService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Alamat)
		err := service.UpdateAlamatService(ctx, req)
		return nil, err
	}
}
