package server

import (
	"context"
)

//langkah ke-5
type kota struct {
	writer ReadWriter
}

func NewKota(writer ReadWriter) KotaService {
	return &kota{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (t *kota) AddKotaService(ctx context.Context, kota Kota) error {
	//fmt.Println("customer")
	err := t.writer.AddKota(kota)
	if err != nil {
		return err
	}

	return nil
}

func (c *kota) ReadKotaService(ctx context.Context) (Kotas, error) {
	kot, err := c.writer.ReadKota()
	//fmt.Println("customer", cus)
	if err != nil {
		return kot, err
	}
	return kot, nil
}

func (t *kota) UpdateKotaService(ctx context.Context, kot Kota) error {
	err := t.writer.UpdateKota(kot)
	if err != nil {
		return err
	}
	return nil
}

func (t *kota) ReadKotaByNamaService(ctx context.Context, nama string) (Kota, error) {
	kot, err := t.writer.ReadKotaByNama(nama)
	//fmt.Println("customer:", cus)
	if err != nil {
		return kot, err
	}
	return kot, nil
}
