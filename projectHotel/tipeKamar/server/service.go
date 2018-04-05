package server

import "context"

type Status int32
type CreatedBy string
type UpdatedBy string

const (
	//ServiceID is dispatch service ID
	ServiceID           = "tipeKamar.hotel.id"
	OnAdd     Status    = 1
	OnAdd2    CreatedBy = "Admin"
	OnAdd3    UpdatedBy = "Admin"
)

type TipeKamar struct {
	IdTipeKamar   string
	NamaTipeKamar string
	HargaKamar    int32
	Status        string
}
type TipeKamars []TipeKamar

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

// interface sebagai parameter
type ReadWriter interface {
	AddTipeKamar(TipeKamar) error
	ReadTipeKamarByHarga(int32) (TipeKamar, error)
	ReadTipeKamar() (TipeKamars, error)
	UpdateTipeKamar(TipeKamar) error
	ReadTipeKamarByNama(string) (TipeKamar, error)
}

//interface sebagai nilai return
type TipeKamarService interface {
	AddTipeKamarService(context.Context, TipeKamar) error
	ReadTipeKamarByHargaService(context.Context, int32) (TipeKamar, error)
	ReadTipeKamarService(context.Context) (TipeKamars, error)
	UpdateTipeKamarService(context.Context, TipeKamar) error
	ReadTipeKamarByNamaService(context.Context, string) (TipeKamar, error)
}
