package server

import "context"

type Status int32
type CreatedBy string
type UpdatedBy string

const (
	//ServiceID is dispatch service ID
	ServiceID           = "jabatan.hotel.id"
	OnAdd     Status    = 1
	OnAdd2    CreatedBy = "Admin"
	OnAdd3    UpdatedBy = "Admin"
)

type Jabatan struct {
	IdJabatan   string
	NamaJabatan string
	Status      string
}
type Jabatans []Jabatan

// interface sebagai parameter
type ReadWriter interface {
	AddJabatan(Jabatan) error
	ReadJabatan() (Jabatans, error)
	UpdateJabatan(Jabatan) error
	ReadJabatanByNama(string) (Jabatan, error)
}

//interface sebagai nilai return
type JabatanService interface {
	AddJabatanService(context.Context, Jabatan) error
	ReadJabatanService(context.Context) (Jabatans, error)
	UpdateJabatanService(context.Context, Jabatan) error
	ReadJabatanByNamaService(context.Context, string) (Jabatan, error)
}
