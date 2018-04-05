package server

import "context"

type Status int32
type CreatedBy string
type UpdatedBy string

const (
	//ServiceID is dispatch service ID
	ServiceID           = "alamat.hotel.id"
	OnAdd     Status    = 1
	OnAdd2    CreatedBy = "Admin"
	OnAdd3    UpdatedBy = "Admin"
)

type Alamat struct {
	IdAlamat    string
	AlamatRumah string
	RtRw        string
	NoRumah     string
	IdKota      string
	IdNegara    string
	Status      string
}
type Alamats []Alamat

// interface sebagai parameter
type ReadWriter interface {
	AddAlamat(Alamat) error
	ReadAlamatByNoRumah(string) (Alamat, error)
	ReadAlamat() (Alamats, error)
	UpdateAlamat(Alamat) error
}

//interface sebagai nilai return
type AlamatService interface {
	AddAlamatService(context.Context, Alamat) error
	ReadAlamatByNoRumahService(context.Context, string) (Alamat, error)
	ReadAlamatService(context.Context) (Alamats, error)
	UpdateAlamatService(context.Context, Alamat) error
}
