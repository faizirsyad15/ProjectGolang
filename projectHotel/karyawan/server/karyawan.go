package server

import (
	"context"
)

//langkah ke-5
type karyawan struct {
	writer ReadWriter
}

func NewKaryawan(writer ReadWriter) KaryawanService {
	return &karyawan{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (t *karyawan) AddKaryawanService(ctx context.Context, karyawan Karyawan) error {
	//fmt.Println("customer")
	err := t.writer.AddKaryawan(karyawan)
	if err != nil {
		return err
	}

	return nil
}

func (c *karyawan) ReadKaryawanByTeleponService(ctx context.Context, tel string) (Karyawan, error) {
	kar, err := c.writer.ReadKaryawanByTelepon(tel)
	//fmt.Println(cus)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (c *karyawan) ReadKaryawanService(ctx context.Context) (Karyawans, error) {
	kar, err := c.writer.ReadKaryawan()
	//fmt.Println("customer", cus)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (t *karyawan) UpdateKaryawanService(ctx context.Context, kar Karyawan) error {
	err := t.writer.UpdateKaryawan(kar)
	if err != nil {
		return err
	}
	return nil
}

func (t *karyawan) ReadKaryawanByNamaService(ctx context.Context, nama string) (Karyawan, error) {
	kar, err := t.writer.ReadKaryawanByNama(nama)
	//fmt.Println("customer:", cus)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (c *karyawan) ReadKaryawanByKeteranganService(ctx context.Context, keterangan string) (Karyawans, error) {
	kar, err := c.writer.ReadKaryawanByKeterangan(keterangan)
	//fmt.Println("customer", cus)
	if err != nil {
		return kar, err
	}
	return kar, nil
}
