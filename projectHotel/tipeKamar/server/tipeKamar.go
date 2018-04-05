package server

import (
	"context"
)

//langkah ke-5
type tipeKamar struct {
	writer ReadWriter
}

func NewTipeKamar(writer ReadWriter) TipeKamarService {
	return &tipeKamar{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (t *tipeKamar) AddTipeKamarService(ctx context.Context, tipeKamar TipeKamar) error {
	//fmt.Println("customer")
	err := t.writer.AddTipeKamar(tipeKamar)
	if err != nil {
		return err
	}

	return nil
}

func (c *tipeKamar) ReadTipeKamarByHargaService(ctx context.Context, tel int32) (TipeKamar, error) {
	tip, err := c.writer.ReadTipeKamarByHarga(tel)
	//fmt.Println(cus)
	if err != nil {
		return tip, err
	}
	return tip, nil
}

func (c *tipeKamar) ReadTipeKamarService(ctx context.Context) (TipeKamars, error) {
	tip, err := c.writer.ReadTipeKamar()
	//fmt.Println("customer", cus)
	if err != nil {
		return tip, err
	}
	return tip, nil
}

func (t *tipeKamar) UpdateTipeKamarService(ctx context.Context, tip TipeKamar) error {
	err := t.writer.UpdateTipeKamar(tip)
	if err != nil {
		return err
	}
	return nil
}

func (t *tipeKamar) ReadTipeKamarByNamaService(ctx context.Context, nama string) (TipeKamar, error) {
	tip, err := t.writer.ReadTipeKamarByNama(nama)
	//fmt.Println("customer:", cus)
	if err != nil {
		return tip, err
	}
	return tip, nil
}
