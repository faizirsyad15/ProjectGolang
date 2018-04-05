package endpoint

import (
	"context"
	"fmt"

	sv "projectHotel/hotel/tamu/server"
)

func (ce TamuEndpoint) AddTamuService(ctx context.Context, tamu sv.Tamu) error {
	_, err := ce.AddTamuEndpoint(ctx, tamu)
	return err
}

func (ce TamuEndpoint) ReadTamuByTeleponService(ctx context.Context, telepon string) (sv.Tamu, error) {
	req := sv.Tamu{NoTelepon: telepon}
	fmt.Println(req)
	resp, err := ce.ReadTamuByTeleponEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	tam := resp.(sv.Tamu)
	return tam, err
}

func (ce TamuEndpoint) ReadTamuService(ctx context.Context) (sv.Tamus, error) {
	resp, err := ce.ReadTamuEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Tamus), err
}

func (ce TamuEndpoint) UpdateTamuService(ctx context.Context, tam sv.Tamu) error {
	_, err := ce.UpdateTamuEndpoint(ctx, tam)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (ce TamuEndpoint) ReadTamuByNamaService(ctx context.Context, nama string) (sv.Tamu, error) {
	req := sv.Tamu{NamaTamu: nama}
	resp, err := ce.ReadTamuByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	tam := resp.(sv.Tamu)
	return tam, err
}
