package server

import (
	"context"
)

//langkah ke-5
type tamu struct {
	writer ReadWriter
}

func NewTamu(writer ReadWriter) TamuService {
	return &tamu{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (t *tamu) AddTamuService(ctx context.Context, tamu Tamu) error {
	//fmt.Println("customer")
	err := t.writer.AddTamu(tamu)
	if err != nil {
		return err
	}

	return nil
}

func (c *tamu) ReadTamuByTeleponService(ctx context.Context, tel string) (Tamu, error) {
	tam, err := c.writer.ReadTamuByTelepon(tel)
	//fmt.Println(cus)
	if err != nil {
		return tam, err
	}
	return tam, nil
}

func (c *tamu) ReadTamuService(ctx context.Context) (Tamus, error) {
	tam, err := c.writer.ReadTamu()
	//fmt.Println("customer", cus)
	if err != nil {
		return tam, err
	}
	return tam, nil
}

func (t *tamu) UpdateTamuService(ctx context.Context, tam Tamu) error {
	err := t.writer.UpdateTamu(tam)
	if err != nil {
		return err
	}
	return nil
}

func (t *tamu) ReadTamuByNamaService(ctx context.Context, nama string) (Tamu, error) {
	tam, err := t.writer.ReadTamuByNama(nama)
	//fmt.Println("customer:", cus)
	if err != nil {
		return tam, err
	}
	return tam, nil
}
