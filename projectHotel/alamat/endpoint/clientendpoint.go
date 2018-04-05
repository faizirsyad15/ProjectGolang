package endpoint

import (
	"context"
	"fmt"

	sv "projectHotel/hotel/alamat/server"
)

func (ce AlamatEndpoint) AddAlamatService(ctx context.Context, alamat sv.Alamat) error {
	_, err := ce.AddAlamatEndpoint(ctx, alamat)
	return err
}

func (ce AlamatEndpoint) ReadAlamatByNoRumahService(ctx context.Context, norumah string) (sv.Alamat, error) {
	req := sv.Alamat{NoRumah: norumah}
	fmt.Println(req)
	resp, err := ce.ReadAlamatByNoRumahEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	amt := resp.(sv.Alamat)
	return amt, err
}

func (ce AlamatEndpoint) ReadAlamatService(ctx context.Context) (sv.Alamats, error) {
	resp, err := ce.ReadAlamatEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Alamats), err
}

func (ce AlamatEndpoint) UpdateAlamatService(ctx context.Context, amt sv.Alamat) error {
	_, err := ce.UpdateAlamatEndpoint(ctx, amt)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
