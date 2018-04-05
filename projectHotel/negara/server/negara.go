package server

import (
	"context"
)

//langkah ke-5
type negara struct {
	writer ReadWriter
}

func NewNegara(writer ReadWriter) NegaraService {
	return &negara{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (t *negara) AddNegaraService(ctx context.Context, negara Negara) error {
	//fmt.Println("customer")
	err := t.writer.AddNegara(negara)
	if err != nil {
		return err
	}

	return nil
}

func (c *negara) ReadNegaraService(ctx context.Context) (Negaras, error) {
	neg, err := c.writer.ReadNegara()
	//fmt.Println("customer", cus)
	if err != nil {
		return neg, err
	}
	return neg, nil
}

func (t *negara) UpdateNegaraService(ctx context.Context, neg Negara) error {
	err := t.writer.UpdateNegara(neg)
	if err != nil {
		return err
	}
	return nil
}

func (t *negara) ReadNegaraByNamaService(ctx context.Context, nama string) (Negara, error) {
	neg, err := t.writer.ReadNegaraByNama(nama)
	//fmt.Println("customer:", cus)
	if err != nil {
		return neg, err
	}
	return neg, nil
}
