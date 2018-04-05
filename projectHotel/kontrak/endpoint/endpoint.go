package endpoint

import (
	"context"

	svc "projectHotel/hotel/kontrak/server"

	kit "github.com/go-kit/kit/endpoint"
)

type KontrakEndpoint struct {
	AddKontrakEndpoint           kit.Endpoint
	ReadKontrakBySelesaiEndpoint kit.Endpoint
	ReadKontrakEndpoint          kit.Endpoint
	UpdateKontrakEndpoint        kit.Endpoint
	ReadKontrakByMulaiEndpoint   kit.Endpoint
}

func NewKontrakEndpoint(service svc.KontrakService) KontrakEndpoint {
	addKontrakEp := makeAddKontrakEndpoint(service)
	readKontrakBySelesaiEp := makeReadKontrakBySelesaiEndpoint(service)
	readKontrakEp := makeReadKontrakEndpoint(service)
	updateKontrakEp := makeUpdateKontrakEndpoint(service)
	readKontrakByMulaiEp := makeReadKontrakByMulaiEndpoint(service)
	return KontrakEndpoint{AddKontrakEndpoint: addKontrakEp,
		ReadKontrakBySelesaiEndpoint: readKontrakBySelesaiEp,
		ReadKontrakEndpoint:          readKontrakEp,
		UpdateKontrakEndpoint:        updateKontrakEp,
		ReadKontrakByMulaiEndpoint:   readKontrakByMulaiEp,
	}
}

func makeAddKontrakEndpoint(service svc.KontrakService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kontrak)
		err := service.AddKontrakService(ctx, req)
		return nil, err
	}
}

func makeReadKontrakBySelesaiEndpoint(service svc.KontrakService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kontrak)
		result, err := service.ReadKontrakBySelesaiService(ctx, req.TanggalSelesai)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadKontrakEndpoint(service svc.KontrakService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadKontrakService(ctx)
		return result, err
	}
}

func makeUpdateKontrakEndpoint(service svc.KontrakService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kontrak)
		err := service.UpdateKontrakService(ctx, req)
		return nil, err
	}
}

func makeReadKontrakByMulaiEndpoint(service svc.KontrakService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kontrak)
		result, err := service.ReadKontrakByMulaiService(ctx, req.TanggalMulai)
		return result, err
	}
}
