package server

import (
	"context"
)

//langkah ke-5
type golongan struct {
	writer ReadWriter
}

func NewGolongan(writer ReadWriter) GolonganService {
	return &golongan{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (t *golongan) AddGolonganService(ctx context.Context, golongan Golongan) error {
	//fmt.Println("customer")
	err := t.writer.AddGolongan(golongan)
	if err != nil {
		return err
	}

	return nil
}

func (c *golongan) ReadGolonganService(ctx context.Context) (Golongans, error) {
	gol, err := c.writer.ReadGolongan()
	//fmt.Println("customer", cus)
	if err != nil {
		return gol, err
	}
	return gol, nil
}

func (t *golongan) UpdateGolonganService(ctx context.Context, gol Golongan) error {
	err := t.writer.UpdateGolongan(gol)
	if err != nil {
		return err
	}
	return nil
}

func (t *golongan) ReadGolonganByNamaService(ctx context.Context, nama string) (Golongan, error) {
	gol, err := t.writer.ReadGolonganByNama(nama)
	//fmt.Println("customer:", cus)
	if err != nil {
		return gol, err
	}
	return gol, nil
}
