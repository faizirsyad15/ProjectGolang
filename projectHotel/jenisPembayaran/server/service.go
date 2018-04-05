package server

import "context"

type Status int32
type CreatedBy string
type UpdatedBy string

const (
	//ServiceID is dispatch service ID
	ServiceID           = "jenisPembayaran.hotel.id"
	OnAdd     Status    = 1
	OnAdd2    CreatedBy = "Admin"
	OnAdd3    UpdatedBy = "Admin"
)

type JenisPembayaran struct {
	IdJenisPembayaran string
	MetodePembayaran  string
	Status            string
}
type JenisPembayarans []JenisPembayaran

// interface sebagai parameter
type ReadWriter interface {
	AddJenisPembayaran(JenisPembayaran) error
	ReadJenisPembayaran() (JenisPembayarans, error)
	UpdateJenisPembayaran(JenisPembayaran) error
	ReadJenisPembayaranByMetode(string) (JenisPembayaran, error)
}

//interface sebagai nilai return
type JenisPembayaranService interface {
	AddJenisPembayaranService(context.Context, JenisPembayaran) error
	ReadJenisPembayaranService(context.Context) (JenisPembayarans, error)
	UpdateJenisPembayaranService(context.Context, JenisPembayaran) error
	ReadJenisPembayaranByMetodeService(context.Context, string) (JenisPembayaran, error)
}
