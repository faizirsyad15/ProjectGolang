package server

import (
	"context"
)

//langkah ke-5
type jabatan struct {
	writer ReadWriter
}

func NewJabatan(writer ReadWriter) JabatanService {
	return &jabatan{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (t *jabatan) AddJabatanService(ctx context.Context, jabatan Jabatan) error {
	//fmt.Println("customer")
	err := t.writer.AddJabatan(jabatan)
	if err != nil {
		return err
	}

	return nil
}

func (c *jabatan) ReadJabatanService(ctx context.Context) (Jabatans, error) {
	jab, err := c.writer.ReadJabatan()
	//fmt.Println("customer", cus)
	if err != nil {
		return jab, err
	}
	return jab, nil
}

func (t *jabatan) UpdateJabatanService(ctx context.Context, jab Jabatan) error {
	err := t.writer.UpdateJabatan(jab)
	if err != nil {
		return err
	}
	return nil
}

func (t *jabatan) ReadJabatanByNamaService(ctx context.Context, nama string) (Jabatan, error) {
	jab, err := t.writer.ReadJabatanByNama(nama)
	//fmt.Println("customer:", cus)
	if err != nil {
		return jab, err
	}
	return jab, nil
}
