package endpoint

import (
	"context"
	"fmt"

	sv "projectHotel/hotel/tipeKamar/server"
)

func (ce TipeKamarEndpoint) AddTipeKamarService(ctx context.Context, tipeKamar sv.TipeKamar) error {
	_, err := ce.AddTipeKamarEndpoint(ctx, tipeKamar)
	return err
}

func (ce TipeKamarEndpoint) ReadTipeKamarByHargaService(ctx context.Context, harga int32) (sv.TipeKamar, error) {
	req := sv.TipeKamar{HargaKamar: harga}
	fmt.Println(req)
	resp, err := ce.ReadTipeKamarByHargaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	tam := resp.(sv.TipeKamar)
	return tam, err
}

func (ce TipeKamarEndpoint) ReadTipeKamarService(ctx context.Context) (sv.TipeKamars, error) {
	resp, err := ce.ReadTipeKamarEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.TipeKamars), err
}

func (ce TipeKamarEndpoint) UpdateTipeKamarService(ctx context.Context, tam sv.TipeKamar) error {
	_, err := ce.UpdateTipeKamarEndpoint(ctx, tam)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (ce TipeKamarEndpoint) ReadTipeKamarByNamaService(ctx context.Context, nama string) (sv.TipeKamar, error) {
	req := sv.TipeKamar{NamaTipeKamar: nama}
	resp, err := ce.ReadTipeKamarByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	tam := resp.(sv.TipeKamar)
	return tam, err
}
