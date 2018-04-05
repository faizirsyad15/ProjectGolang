package server

import "context"

type Status int32
type CreatedBy string
type UpdatedBy string

const (
	//ServiceID is dispatch service ID
	ServiceID           = "kamar.hotel.id"
	OnAdd     Status    = 1
	OnAdd2    CreatedBy = "Admin"
	OnAdd3    UpdatedBy = "Admin"
)

type Kamar struct {
	IdKamar     string
	IdTipeKamar string
	IdMenu      string
	Status      string
}
type Kamars []Kamar

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
	AddKamar(Kamar) error
	ReadKamar() (Kamars, error)
	UpdateKamar(Kamar) error
}

//interface sebagai nilai return
type KamarService interface {
	AddKamarService(context.Context, Kamar) error
	ReadKamarService(context.Context) (Kamars, error)
	UpdateKamarService(context.Context, Kamar) error
}
