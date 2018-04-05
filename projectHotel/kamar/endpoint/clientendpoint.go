package endpoint

import (
	"context"
	"fmt"

	sv "projectHotel/hotel/kamar/server"
)

func (ce KamarEndpoint) AddKamarService(ctx context.Context, kamar sv.Kamar) error {
	_, err := ce.AddKamarEndpoint(ctx, kamar)
	return err
}

func (ce KamarEndpoint) ReadKamarService(ctx context.Context) (sv.Kamars, error) {
	resp, err := ce.ReadKamarEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Kamars), err
}

func (ce KamarEndpoint) UpdateKamarService(ctx context.Context, tam sv.Kamar) error {
	_, err := ce.UpdateKamarEndpoint(ctx, tam)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
