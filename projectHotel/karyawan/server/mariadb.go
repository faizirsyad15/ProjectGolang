package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addKaryawan = `insert into mst_karyawan(id_karyawan,nama_karyawan,alamat,no_telepon,keterangan,id_jabatan,id_departemen,id_golongan,id_kontrak,status,created_by,
		created_on)values (?,?,?,?,?,?,?,?,?,?,?,?)`
	selectKaryawanByTelepon = `select id_karyawan,nama_karyawan,alamat,keterangan,id_jabatan,id_departemen,id_golongan,id_kontrak,status
		from mst_karyawan where no_telepon = ?`
	selectKaryawan = `select id_karyawan,nama_karyawan,alamat,no_telepon,keterangan,id_jabatan,id_departemen,id_golongan,id_kontrak,status 
		from mst_karyawan`
	updateKaryawan = `update mst_karyawan set nama_karyawan=?,alamat=?,no_telepon=?,keterangan=?,id_jabatan=?,id_departemen=?,id_golongan=?,id_kontrak=?,status=?,
		updated_by=?,updated_on=? where id_karyawan=?`
	selectKaryawanByNama = `select id_karyawan,alamat,no_telepon,keterangan,id_jabatan,id_departemen,id_golongan,id_kontrak,status
		from mst_karyawan where nama_karyawan=?`
	selectKaryawanByKeterangan = `select id_karyawan,nama_karyawan,alamat,no_telepon,keterangan,id_jabatan,id_departemen,id_golongan,id_kontrak,status 
		from mst_karyawan where keterangan like ?`
)

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddKaryawan(karyawan Karyawan) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addKaryawan, karyawan.IdKaryawan, karyawan.NamaKaryawan, karyawan.Alamat,
		karyawan.NoTelepon, karyawan.Keterangan, karyawan.IdJabatan,
		karyawan.IdDepartemen, karyawan.IdGolongan, karyawan.IdKontrak, OnAdd, OnAdd2, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadKaryawanByTelepon(telepon string) (Karyawan, error) {
	karyawan := Karyawan{NoTelepon: telepon}
	err := rw.db.QueryRow(selectKaryawanByTelepon, telepon).Scan(&karyawan.IdKaryawan, &karyawan.NamaKaryawan,
		&karyawan.Alamat, &karyawan.Keterangan, &karyawan.IdJabatan, &karyawan.IdDepartemen, &karyawan.IdGolongan,
		&karyawan.IdKontrak, &karyawan.Status)

	if err != nil {
		return Karyawan{}, err
	}

	return karyawan, nil
}

func (rw *dbReadWriter) ReadKaryawan() (Karyawans, error) {
	karyawan := Karyawans{}
	rows, _ := rw.db.Query(selectKaryawan)
	defer rows.Close()
	for rows.Next() {
		var t Karyawan
		err := rows.Scan(&t.IdKaryawan, &t.NamaKaryawan, &t.Alamat, &t.NoTelepon,
			&t.Keterangan, &t.IdJabatan,
			&t.IdDepartemen, &t.IdGolongan, &t.IdKontrak, &t.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return karyawan, err
		}
		karyawan = append(karyawan, t)
	}
	//fmt.Println("db nya:", customer)
	return karyawan, nil
}

func (rw *dbReadWriter) UpdateKaryawan(kar Karyawan) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateKaryawan, kar.NamaKaryawan, kar.Alamat, kar.NoTelepon,
		kar.Keterangan, kar.IdJabatan,
		kar.IdDepartemen, kar.IdGolongan, kar.IdKontrak, kar.Status,
		OnAdd3, time.Now(), kar.IdKaryawan)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadKaryawanByNama(nama string) (Karyawan, error) {
	karyawan := Karyawan{NamaKaryawan: nama}
	err := rw.db.QueryRow(selectKaryawanByNama, nama).Scan(&karyawan.IdKaryawan,
		&karyawan.Alamat, &karyawan.NoTelepon, &karyawan.Keterangan, &karyawan.IdJabatan,
		&karyawan.IdDepartemen, &karyawan.IdGolongan, &karyawan.IdKontrak, &karyawan.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Karyawan{}, err
	}

	return karyawan, nil
}

func (rw *dbReadWriter) ReadKaryawanByKeterangan(keterangan string) (Karyawans, error) {
	karyawan := Karyawans{}
	rows, _ := rw.db.Query(selectKaryawanByKeterangan, keterangan)
	defer rows.Close()
	for rows.Next() {
		var t Karyawan
		err := rows.Scan(&t.IdKaryawan, &t.NamaKaryawan, &t.Alamat, &t.NoTelepon,
			&t.Keterangan, &t.IdJabatan,
			&t.IdDepartemen, &t.IdGolongan, &t.IdKontrak, &t.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return karyawan, err
		}
		karyawan = append(karyawan, t)
	}
	//fmt.Println("db nya:", customer)
	return karyawan, nil
}
