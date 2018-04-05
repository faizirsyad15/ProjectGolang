package endpoint

import (
	"context"

	svc "projectHotel/hotel/jabatan/server"

	kit "github.com/go-kit/kit/endpoint"
)

type JabatanEndpoint struct {
	AddJabatanEndpoint        kit.Endpoint
	ReadJabatanEndpoint       kit.Endpoint
	UpdateJabatanEndpoint     kit.Endpoint
	ReadJabatanByNamaEndpoint kit.Endpoint
}

func NewJabatanEndpoint(service svc.JabatanService) JabatanEndpoint {
	addJabatanEp := makeAddJabatanEndpoint(service)
	readJabatanEp := makeReadJabatanEndpoint(service)
	updateJabatanEp := makeUpdateJabatanEndpoint(service)
	readJabatanByNamaEp := makeReadJabatanByNamaEndpoint(service)
	return JabatanEndpoint{AddJabatanEndpoint: addJabatanEp,
		ReadJabatanEndpoint:       readJabatanEp,
		UpdateJabatanEndpoint:     updateJabatanEp,
		ReadJabatanByNamaEndpoint: readJabatanByNamaEp,
	}
}

func makeAddJabatanEndpoint(service svc.JabatanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Jabatan)
		err := service.AddJabatanService(ctx, req)
		return nil, err
	}
}

func makeReadJabatanEndpoint(service svc.JabatanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadJabatanService(ctx)
		return result, err
	}
}

func makeUpdateJabatanEndpoint(service svc.JabatanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Jabatan)
		err := service.UpdateJabatanService(ctx, req)
		return nil, err
	}
}

func makeReadJabatanByNamaEndpoint(service svc.JabatanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Jabatan)
		result, err := service.ReadJabatanByNamaService(ctx, req.NamaJabatan)
		return result, err
	}
}
