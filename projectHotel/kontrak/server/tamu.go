package server

import (
	"context"
)

//langkah ke-5
type kontrak struct {
	writer ReadWriter
}

func NewKontrak(writer ReadWriter) KontrakService {
	return &kontrak{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (t *kontrak) AddKontrakService(ctx context.Context, kontrak Kontrak) error {
	//fmt.Println("customer")
	err := t.writer.AddKontrak(kontrak)
	if err != nil {
		return err
	}

	return nil
}

func (c *kontrak) ReadKontrakBySelesaiService(ctx context.Context, tel string) (Kontrak, error) {
	ktk, err := c.writer.ReadKontrakBySelesai(tel)
	//fmt.Println(cus)
	if err != nil {
		return ktk, err
	}
	return ktk, nil
}

func (c *kontrak) ReadKontrakService(ctx context.Context) (Kontraks, error) {
	ktk, err := c.writer.ReadKontrak()
	//fmt.Println("customer", cus)
	if err != nil {
		return ktk, err
	}
	return ktk, nil
}

func (t *kontrak) UpdateKontrakService(ctx context.Context, ktk Kontrak) error {
	err := t.writer.UpdateKontrak(ktk)
	if err != nil {
		return err
	}
	return nil
}

func (t *kontrak) ReadKontrakByMulaiService(ctx context.Context, nama string) (Kontrak, error) {
	ktk, err := t.writer.ReadKontrakByMulai(nama)
	//fmt.Println("customer:", cus)
	if err != nil {
		return ktk, err
	}
	return ktk, nil
}
