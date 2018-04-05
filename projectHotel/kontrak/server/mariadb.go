package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addKontrak = `insert into mst_kontrak(id_kontrak,tanggal_mulai,tanggal_selesai,keterangan,status,created_by,
		created_on)values (?,?,?,?,?,?,?)`
	selectKontrakBySelesai = `select id_kontrak,tanggal_mulai,keterangan,status
		from mst_kontrak where tanggal_selesai = ?`
	selectKontrak = `select id_kontrak,tanggal_mulai,tanggal_selesai,keterangan,status 
		from mst_kontrak`
	updateKontrak = `update mst_kontrak set tanggal_mulai=?,tanggal_selesai=?,keterangan=?,status=?,
		updated_by=?,updated_on=? where id_kontrak=?`
	selectKontrakByMulai = `select id_kontrak,tanggal_selesai,keterangan,status
		from mst_kontrak where tanggal_mulai=?`
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

func (rw *dbReadWriter) AddKontrak(kontrak Kontrak) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addKontrak, kontrak.IdKontrak, kontrak.TanggalMulai, kontrak.TanggalSelesai, kontrak.Keterangan,
		OnAdd, OnAdd2, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadKontrakBySelesai(telepon string) (Kontrak, error) {
	kontrak := Kontrak{TanggalSelesai: telepon}
	err := rw.db.QueryRow(selectKontrakBySelesai, telepon).Scan(&kontrak.IdKontrak, &kontrak.TanggalMulai,
		&kontrak.Keterangan, &kontrak.Status)

	if err != nil {
		return Kontrak{}, err
	}

	return kontrak, nil
}

func (rw *dbReadWriter) ReadKontrak() (Kontraks, error) {
	kontrak := Kontraks{}
	rows, _ := rw.db.Query(selectKontrak)
	defer rows.Close()
	for rows.Next() {
		var t Kontrak
		err := rows.Scan(&t.IdKontrak, &t.TanggalMulai, &t.TanggalSelesai, &t.Keterangan, &t.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return kontrak, err
		}
		kontrak = append(kontrak, t)
	}
	//fmt.Println("db nya:", customer)
	return kontrak, nil
}

func (rw *dbReadWriter) UpdateKontrak(ktk Kontrak) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateKontrak, ktk.TanggalMulai, ktk.TanggalSelesai, ktk.Keterangan, ktk.Status,
		OnAdd3, time.Now(), ktk.IdKontrak)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadKontrakByMulai(nama string) (Kontrak, error) {
	kontrak := Kontrak{TanggalMulai: nama}
	err := rw.db.QueryRow(selectKontrakByMulai, nama).Scan(&kontrak.IdKontrak,
		&kontrak.TanggalSelesai, &kontrak.Keterangan, &kontrak.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Kontrak{}, err
	}

	return kontrak, nil
}
