package server

import "context"

type Status int32
type CreatedBy string
type UpdatedBy string

const (
	//ServiceID is dispatch service ID
	ServiceID           = "kota.hotel.id"
	OnAdd     Status    = 1
	OnAdd2    CreatedBy = "Admin"
	OnAdd3    UpdatedBy = "Admin"
)

type Kota struct {
	IdKota   string
	NamaKota string
	Status   string
}
type Kotas []Kota

// interface sebagai parameter
type ReadWriter interface {
	AddKota(Kota) error
	ReadKota() (Kotas, error)
	UpdateKota(Kota) error
	ReadKotaByNama(string) (Kota, error)
}

//interface sebagai nilai return
type KotaService interface {
	AddKotaService(context.Context, Kota) error
	ReadKotaService(context.Context) (Kotas, error)
	UpdateKotaService(context.Context, Kota) error
	ReadKotaByNamaService(context.Context, string) (Kota, error)
}
