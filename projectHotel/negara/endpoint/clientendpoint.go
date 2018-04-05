package endpoint

import (
	"context"
	"fmt"

	sv "projectHotel/hotel/negara/server"
)

func (ce NegaraEndpoint) AddNegaraService(ctx context.Context, negara sv.Negara) error {
	_, err := ce.AddNegaraEndpoint(ctx, negara)
	return err
}

func (ce NegaraEndpoint) ReadNegaraService(ctx context.Context) (sv.Negaras, error) {
	resp, err := ce.ReadNegaraEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Negaras), err
}

func (ce NegaraEndpoint) UpdateNegaraService(ctx context.Context, neg sv.Negara) error {
	_, err := ce.UpdateNegaraEndpoint(ctx, neg)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (ce NegaraEndpoint) ReadNegaraByNamaService(ctx context.Context, nama string) (sv.Negara, error) {
	req := sv.Negara{NamaNegara: nama}
	resp, err := ce.ReadNegaraByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	neg := resp.(sv.Negara)
	return neg, err
}
