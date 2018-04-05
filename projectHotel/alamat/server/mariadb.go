package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addAlamat = `insert into mst_alamat(id_alamat,alamat_rumah,rt_rw,no_rumah,id_kota,id_negara,status,created_by,
		created_on)values (?,?,?,?,?,?,?,?,?)`
	selectAlamatByNoRumah = `select id_alamat,alamat_rumah,rt_rw,id_kota,id_negara,status
		from mst_alamat where no_rumah = ?`
	selectAlamat = `select id_alamat,alamat_rumah,rt_rw,no_rumah,id_kota,id_negara,status 
		from mst_alamat`
	updateAlamat = `update mst_alamat set alamat_rumah=?,rt_rw=?,no_rumah=?,id_kota=?,id_negara=?,status=?,
		updated_by=?,updated_on=? where id_alamat=?`
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

func (rw *dbReadWriter) AddAlamat(alamat Alamat) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addAlamat, alamat.IdAlamat, alamat.AlamatRumah, alamat.RtRw, alamat.NoRumah, alamat.IdKota,
		alamat.IdNegara, OnAdd, OnAdd2, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadAlamatByNoRumah(norumah string) (Alamat, error) {
	alamat := Alamat{NoRumah: norumah}
	err := rw.db.QueryRow(selectAlamatByNoRumah, norumah).Scan(&alamat.IdAlamat, &alamat.AlamatRumah,
		&alamat.RtRw, &alamat.IdKota, &alamat.IdNegara, &alamat.Status)

	if err != nil {
		return Alamat{}, err
	}

	return alamat, nil
}

func (rw *dbReadWriter) ReadAlamat() (Alamats, error) {
	alamat := Alamats{}
	rows, _ := rw.db.Query(selectAlamat)
	defer rows.Close()
	for rows.Next() {
		var a Alamat
		err := rows.Scan(&a.IdAlamat, &a.AlamatRumah, &a.RtRw, &a.NoRumah, &a.IdKota, &a.IdNegara, &a.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return alamat, err
		}
		alamat = append(alamat, a)
	}
	//fmt.Println("db nya:", customer)
	return alamat, nil
}

func (rw *dbReadWriter) UpdateAlamat(amt Alamat) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateAlamat, amt.AlamatRumah, amt.RtRw, amt.NoRumah, amt.IdKota, amt.IdNegara,
		amt.Status, OnAdd3, time.Now(), amt.IdAlamat)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}
