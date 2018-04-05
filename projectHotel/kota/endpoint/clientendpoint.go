package endpoint

import (
	"context"
	"fmt"

	sv "projectHotel/hotel/kota/server"
)

func (ce KotaEndpoint) AddKotaService(ctx context.Context, kota sv.Kota) error {
	_, err := ce.AddKotaEndpoint(ctx, kota)
	return err
}

func (ce KotaEndpoint) ReadKotaService(ctx context.Context) (sv.Kotas, error) {
	resp, err := ce.ReadKotaEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Kotas), err
}

func (ce KotaEndpoint) UpdateKotaService(ctx context.Context, kot sv.Kota) error {
	_, err := ce.UpdateKotaEndpoint(ctx, kot)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (ce KotaEndpoint) ReadKotaByNamaService(ctx context.Context, nama string) (sv.Kota, error) {
	req := sv.Kota{NamaKota: nama}
	resp, err := ce.ReadKotaByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	kot := resp.(sv.Kota)
	return kot, err
}
