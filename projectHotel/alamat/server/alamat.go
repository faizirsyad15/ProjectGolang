package server

import (
	"context"
)

//langkah ke-5
type alamat struct {
	writer ReadWriter
}

func NewAlamat(writer ReadWriter) AlamatService {
	return &alamat{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (t *alamat) AddAlamatService(ctx context.Context, alamat Alamat) error {
	//fmt.Println("customer")
	err := t.writer.AddAlamat(alamat)
	if err != nil {
		return err
	}

	return nil
}

func (c *alamat) ReadAlamatByNoRumahService(ctx context.Context, nor string) (Alamat, error) {
	amt, err := c.writer.ReadAlamatByNoRumah(nor)
	//fmt.Println(cus)
	if err != nil {
		return amt, err
	}
	return amt, nil
}

func (c *alamat) ReadAlamatService(ctx context.Context) (Alamats, error) {
	amt, err := c.writer.ReadAlamat()
	//fmt.Println("customer", cus)
	if err != nil {
		return amt, err
	}
	return amt, nil
}

func (t *alamat) UpdateAlamatService(ctx context.Context, amt Alamat) error {
	err := t.writer.UpdateAlamat(amt)
	if err != nil {
		return err
	}
	return nil
}
