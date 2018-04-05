package server

import "context"

type Status int32
type CreatedBy string
type UpdatedBy string

const (
	//ServiceID is dispatch service ID
	ServiceID           = "kontrak.hotel.id"
	OnAdd     Status    = 1
	OnAdd2    CreatedBy = "Admin"
	OnAdd3    UpdatedBy = "Admin"
)

type Kontrak struct {
	IdKontrak      string
	TanggalMulai   string
	TanggalSelesai string
	Keterangan     string
	Status         string
}
type Kontraks []Kontrak

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
	AddKontrak(Kontrak) error
	ReadKontrakBySelesai(string) (Kontrak, error)
	ReadKontrak() (Kontraks, error)
	UpdateKontrak(Kontrak) error
	ReadKontrakByMulai(string) (Kontrak, error)
}

//interface sebagai nilai return
type KontrakService interface {
	AddKontrakService(context.Context, Kontrak) error
	ReadKontrakBySelesaiService(context.Context, string) (Kontrak, error)
	ReadKontrakService(context.Context) (Kontraks, error)
	UpdateKontrakService(context.Context, Kontrak) error
	ReadKontrakByMulaiService(context.Context, string) (Kontrak, error)
}
