package endpoint

import (
	"context"
	"fmt"

	sv "projectHotel/hotel/jabatan/server"
)

func (ce JabatanEndpoint) AddJabatanService(ctx context.Context, jabatan sv.Jabatan) error {
	_, err := ce.AddJabatanEndpoint(ctx, jabatan)
	return err
}

func (ce JabatanEndpoint) ReadJabatanService(ctx context.Context) (sv.Jabatans, error) {
	resp, err := ce.ReadJabatanEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Jabatans), err
}

func (ce JabatanEndpoint) UpdateJabatanService(ctx context.Context, jab sv.Jabatan) error {
	_, err := ce.UpdateJabatanEndpoint(ctx, jab)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (ce JabatanEndpoint) ReadJabatanByNamaService(ctx context.Context, nama string) (sv.Jabatan, error) {
	req := sv.Jabatan{NamaJabatan: nama}
	resp, err := ce.ReadJabatanByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	jab := resp.(sv.Jabatan)
	return jab, err
}
