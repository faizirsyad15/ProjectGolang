package endpoint

import (
	"context"
	"fmt"

	sv "projectHotel/hotel/karyawan/server"
)

func (ce KaryawanEndpoint) AddKaryawanService(ctx context.Context, karyawan sv.Karyawan) error {
	_, err := ce.AddKaryawanEndpoint(ctx, karyawan)
	return err
}

func (ce KaryawanEndpoint) ReadKaryawanByTeleponService(ctx context.Context, telepon string) (sv.Karyawan, error) {
	req := sv.Karyawan{NoTelepon: telepon}
	fmt.Println(req)
	resp, err := ce.ReadKaryawanByTeleponEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	kar := resp.(sv.Karyawan)
	return kar, err
}

func (ce KaryawanEndpoint) ReadKaryawanService(ctx context.Context) (sv.Karyawans, error) {
	resp, err := ce.ReadKaryawanEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Karyawans), err
}

func (ce KaryawanEndpoint) UpdateKaryawanService(ctx context.Context, kar sv.Karyawan) error {
	_, err := ce.UpdateKaryawanEndpoint(ctx, kar)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (ce KaryawanEndpoint) ReadKaryawanByNamaService(ctx context.Context, nama string) (sv.Karyawan, error) {
	req := sv.Karyawan{NamaKaryawan: nama}
	resp, err := ce.ReadKaryawanByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	kar := resp.(sv.Karyawan)
	return kar, err
}

func (ce KaryawanEndpoint) ReadKaryawanByKeteranganService(ctx context.Context, keterangan string) (sv.Karyawans, error) {
	req := sv.Karyawan{Keterangan: keterangan}
	resp, err := ce.ReadKaryawanByKeteranganEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	kar := resp.(sv.Karyawans)
	return kar, err
}
