package endpoint

import (
	"context"
	"fmt"

	sv "projectHotel/hotel/jenisPembayaran/server"
)

func (ce JenisPembayaranEndpoint) AddJenisPembayaranService(ctx context.Context, jenisPembayaran sv.JenisPembayaran) error {
	_, err := ce.AddJenisPembayaranEndpoint(ctx, jenisPembayaran)
	return err
}

func (ce JenisPembayaranEndpoint) ReadJenisPembayaranService(ctx context.Context) (sv.JenisPembayarans, error) {
	resp, err := ce.ReadJenisPembayaranEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.JenisPembayarans), err
}

func (ce JenisPembayaranEndpoint) UpdateJenisPembayaranService(ctx context.Context, jen sv.JenisPembayaran) error {
	_, err := ce.UpdateJenisPembayaranEndpoint(ctx, jen)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (ce JenisPembayaranEndpoint) ReadJenisPembayaranByMetodeService(ctx context.Context, metode string) (sv.JenisPembayaran, error) {
	req := sv.JenisPembayaran{MetodePembayaran: metode}
	resp, err := ce.ReadJenisPembayaranByMetodeEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	jen := resp.(sv.JenisPembayaran)
	return jen, err
}
