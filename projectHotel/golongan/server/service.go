package server

import "context"

type Status int32
type CreatedBy string
type UpdatedBy string

const (
	//ServiceID is dispatch service ID
	ServiceID           = "golongan.hotel.id"
	OnAdd     Status    = 1
	OnAdd2    CreatedBy = "Admin"
	OnAdd3    UpdatedBy = "Admin"
)

type Golongan struct {
	IdGolongan   string
	NamaGolongan string
	Status       string
}
type Golongans []Golongan

// interface sebagai parameter
type ReadWriter interface {
	AddGolongan(Golongan) error
	ReadGolongan() (Golongans, error)
	UpdateGolongan(Golongan) error
	ReadGolonganByNama(string) (Golongan, error)
}

//interface sebagai nilai return
type GolonganService interface {
	AddGolonganService(context.Context, Golongan) error
	ReadGolonganService(context.Context) (Golongans, error)
	UpdateGolonganService(context.Context, Golongan) error
	ReadGolonganByNamaService(context.Context, string) (Golongan, error)
}
