package endpoint

import (
	"context"
	"fmt"

	sv "projectHotel/hotel/golongan/server"
)

func (ce GolonganEndpoint) AddGolonganService(ctx context.Context, golongan sv.Golongan) error {
	_, err := ce.AddGolonganEndpoint(ctx, golongan)
	return err
}

func (ce GolonganEndpoint) ReadGolonganService(ctx context.Context) (sv.Golongans, error) {
	resp, err := ce.ReadGolonganEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Golongans), err
}

func (ce GolonganEndpoint) UpdateGolonganService(ctx context.Context, gol sv.Golongan) error {
	_, err := ce.UpdateGolonganEndpoint(ctx, gol)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (ce GolonganEndpoint) ReadGolonganByNamaService(ctx context.Context, nama string) (sv.Golongan, error) {
	req := sv.Golongan{NamaGolongan: nama}
	resp, err := ce.ReadGolonganByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	gol := resp.(sv.Golongan)
	return gol, err
}
