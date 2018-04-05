package server

import "context"

type Status int32
type CreatedBy string
type UpdatedBy string

const (
	//ServiceID is dispatch service ID
	ServiceID           = "karyawan.hotel.id"
	OnAdd     Status    = 1
	OnAdd2    CreatedBy = "Admin"
	OnAdd3    UpdatedBy = "Admin"
)

type Karyawan struct {
	IdKaryawan   string
	NamaKaryawan string
	Alamat       string
	NoTelepon    string
	Keterangan   string
	IdJabatan    string
	IdDepartemen string
	IdGolongan   string
	IdKontrak    string
	Status       string
}
type Karyawans []Karyawan

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
	AddKaryawan(Karyawan) error
	ReadKaryawanByTelepon(string) (Karyawan, error)
	ReadKaryawan() (Karyawans, error)
	UpdateKaryawan(Karyawan) error
	ReadKaryawanByNama(string) (Karyawan, error)
	ReadKaryawanByKeterangan(string) (Karyawans, error)
}

//interface sebagai nilai return
type KaryawanService interface {
	AddKaryawanService(context.Context, Karyawan) error
	ReadKaryawanByTeleponService(context.Context, string) (Karyawan, error)
	ReadKaryawanService(context.Context) (Karyawans, error)
	UpdateKaryawanService(context.Context, Karyawan) error
	ReadKaryawanByNamaService(context.Context, string) (Karyawan, error)
	ReadKaryawanByKeteranganService(context.Context, string) (Karyawans, error)
}
