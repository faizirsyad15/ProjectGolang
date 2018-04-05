package server

import (
	"context"
)

//langkah ke-5
type kamar struct {
	writer ReadWriter
}

func NewKamar(writer ReadWriter) KamarService {
	return &kamar{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (t *kamar) AddKamarService(ctx context.Context, kamar Kamar) error {
	//fmt.Println("customer")
	err := t.writer.AddKamar(kamar)
	if err != nil {
		return err
	}

	return nil
}

func (c *kamar) ReadKamarService(ctx context.Context) (Kamars, error) {
	tip, err := c.writer.ReadKamar()
	//fmt.Println("customer", cus)
	if err != nil {
		return tip, err
	}
	return tip, nil
}

func (t *kamar) UpdateKamarService(ctx context.Context, tip Kamar) error {
	err := t.writer.UpdateKamar(tip)
	if err != nil {
		return err
	}
	return nil
}
