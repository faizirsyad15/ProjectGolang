package server

import "context"

type Status int32
type CreatedBy string
type UpdatedBy string

const (
	//ServiceID is dispatch service ID
	ServiceID           = "tamu.hotel.id"
	OnAdd     Status    = 1
	OnAdd2    CreatedBy = "Admin"
	OnAdd3    UpdatedBy = "Admin"
)

type Tamu struct {
	IdTamu       string
	NamaTamu     string
	NoTelepon    string
	JenisKelamin string
	IdAlamat     string
	Status       string
}
type Tamus []Tamu

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
	AddTamu(Tamu) error
	ReadTamuByTelepon(string) (Tamu, error)
	ReadTamu() (Tamus, error)
	UpdateTamu(Tamu) error
	ReadTamuByNama(string) (Tamu, error)
}

//interface sebagai nilai return
type TamuService interface {
	AddTamuService(context.Context, Tamu) error
	ReadTamuByTeleponService(context.Context, string) (Tamu, error)
	ReadTamuService(context.Context) (Tamus, error)
	UpdateTamuService(context.Context, Tamu) error
	ReadTamuByNamaService(context.Context, string) (Tamu, error)
}
