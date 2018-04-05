package server

import "context"

type Status int32
type CreatedBy string
type UpdatedBy string

const (
	//ServiceID is dispatch service ID
	ServiceID           = "negara.hotel.id"
	OnAdd     Status    = 1
	OnAdd2    CreatedBy = "Admin"
	OnAdd3    UpdatedBy = "Admin"
)

type Negara struct {
	IdNegara   string
	NamaNegara string
	Status     string
}
type Negaras []Negara

// interface sebagai parameter
type ReadWriter interface {
	AddNegara(Negara) error
	ReadNegara() (Negaras, error)
	UpdateNegara(Negara) error
	ReadNegaraByNama(string) (Negara, error)
}

//interface sebagai nilai return
type NegaraService interface {
	AddNegaraService(context.Context, Negara) error
	ReadNegaraService(context.Context) (Negaras, error)
	UpdateNegaraService(context.Context, Negara) error
	ReadNegaraByNamaService(context.Context, string) (Negara, error)
}
