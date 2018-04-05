package endpoint

import (
	"context"
	"fmt"

	sv "projectHotel/hotel/kontrak/server"
)

func (ce KontrakEndpoint) AddKontrakService(ctx context.Context, kontrak sv.Kontrak) error {
	_, err := ce.AddKontrakEndpoint(ctx, kontrak)
	return err
}

func (ce KontrakEndpoint) ReadKontrakBySelesaiService(ctx context.Context, selesai string) (sv.Kontrak, error) {
	req := sv.Kontrak{TanggalSelesai: selesai}
	fmt.Println(req)
	resp, err := ce.ReadKontrakBySelesaiEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	ktk := resp.(sv.Kontrak)
	return ktk, err
}

func (ce KontrakEndpoint) ReadKontrakService(ctx context.Context) (sv.Kontraks, error) {
	resp, err := ce.ReadKontrakEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Kontraks), err
}

func (ce KontrakEndpoint) UpdateKontrakService(ctx context.Context, ktk sv.Kontrak) error {
	_, err := ce.UpdateKontrakEndpoint(ctx, ktk)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (ce KontrakEndpoint) ReadKontrakByMulaiService(ctx context.Context, nama string) (sv.Kontrak, error) {
	req := sv.Kontrak{TanggalMulai: nama}
	resp, err := ce.ReadKontrakByMulaiEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	ktk := resp.(sv.Kontrak)
	return ktk, err
}
