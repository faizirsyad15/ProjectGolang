package endpoint

import (
	"context"

	svc "projectHotel/hotel/jenisPembayaran/server"

	kit "github.com/go-kit/kit/endpoint"
)

type JenisPembayaranEndpoint struct {
	AddJenisPembayaranEndpoint          kit.Endpoint
	ReadJenisPembayaranEndpoint         kit.Endpoint
	UpdateJenisPembayaranEndpoint       kit.Endpoint
	ReadJenisPembayaranByMetodeEndpoint kit.Endpoint
}

func NewJenisPembayaranEndpoint(service svc.JenisPembayaranService) JenisPembayaranEndpoint {
	addJenisPembayaranEp := makeAddJenisPembayaranEndpoint(service)
	readJenisPembayaranEp := makeReadJenisPembayaranEndpoint(service)
	updateJenisPembayaranEp := makeUpdateJenisPembayaranEndpoint(service)
	readJenisPembayaranByMetodeEp := makeReadJenisPembayaranByMetodeEndpoint(service)
	return JenisPembayaranEndpoint{AddJenisPembayaranEndpoint: addJenisPembayaranEp,
		ReadJenisPembayaranEndpoint:         readJenisPembayaranEp,
		UpdateJenisPembayaranEndpoint:       updateJenisPembayaranEp,
		ReadJenisPembayaranByMetodeEndpoint: readJenisPembayaranByMetodeEp,
	}
}

func makeAddJenisPembayaranEndpoint(service svc.JenisPembayaranService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.JenisPembayaran)
		err := service.AddJenisPembayaranService(ctx, req)
		return nil, err
	}
}

func makeReadJenisPembayaranEndpoint(service svc.JenisPembayaranService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadJenisPembayaranService(ctx)
		return result, err
	}
}

func makeUpdateJenisPembayaranEndpoint(service svc.JenisPembayaranService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.JenisPembayaran)
		err := service.UpdateJenisPembayaranService(ctx, req)
		return nil, err
	}
}

func makeReadJenisPembayaranByMetodeEndpoint(service svc.JenisPembayaranService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.JenisPembayaran)
		result, err := service.ReadJenisPembayaranByMetodeService(ctx, req.MetodePembayaran)
		return result, err
	}
}
