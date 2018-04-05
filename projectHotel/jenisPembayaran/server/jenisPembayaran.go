package server

import (
	"context"
)

//langkah ke-5
type jenisPembayaran struct {
	writer ReadWriter
}

func NewJenisPembayaran(writer ReadWriter) JenisPembayaranService {
	return &jenisPembayaran{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (t *jenisPembayaran) AddJenisPembayaranService(ctx context.Context, jenisPembayaran JenisPembayaran) error {
	//fmt.Println("customer")
	err := t.writer.AddJenisPembayaran(jenisPembayaran)
	if err != nil {
		return err
	}

	return nil
}

func (c *jenisPembayaran) ReadJenisPembayaranService(ctx context.Context) (JenisPembayarans, error) {
	jen, err := c.writer.ReadJenisPembayaran()
	//fmt.Println("customer", cus)
	if err != nil {
		return jen, err
	}
	return jen, nil
}

func (t *jenisPembayaran) UpdateJenisPembayaranService(ctx context.Context, jen JenisPembayaran) error {
	err := t.writer.UpdateJenisPembayaran(jen)
	if err != nil {
		return err
	}
	return nil
}

func (t *jenisPembayaran) ReadJenisPembayaranByMetodeService(ctx context.Context, metode string) (JenisPembayaran, error) {
	jen, err := t.writer.ReadJenisPembayaranByMetode(metode)
	//fmt.Println("customer:", cus)
	if err != nil {
		return jen, err
	}
	return jen, nil
}
