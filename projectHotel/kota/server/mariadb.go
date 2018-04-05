package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addKota = `insert into mst_kota(id_kota,nama_kota,status,created_by,
		created_on)values (?,?,?,?,?)`
	selectKota = `select id_kota,nama_kota,status 
		from mst_kota`
	updateKota = `update mst_kota set nama_kota=?,status=?,
		updated_by=?,updated_on=? where id_kota=?`
	selectKotaByNama = `select id_kota,nama_kota,status
		from mst_kota where nama_kota=?`
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

func (rw *dbReadWriter) AddKota(kota Kota) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addKota, kota.IdKota, kota.NamaKota,
		OnAdd, OnAdd2, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadKota() (Kotas, error) {
	kota := Kotas{}
	rows, _ := rw.db.Query(selectKota)
	defer rows.Close()
	for rows.Next() {
		var t Kota
		err := rows.Scan(&t.IdKota, &t.NamaKota, &t.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return kota, err
		}
		kota = append(kota, t)
	}
	//fmt.Println("db nya:", customer)
	return kota, nil
}

func (rw *dbReadWriter) UpdateKota(kot Kota) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateKota, kot.NamaKota, kot.Status,
		OnAdd3, time.Now(), kot.IdKota)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadKotaByNama(nama string) (Kota, error) {
	kota := Kota{NamaKota: nama}
	err := rw.db.QueryRow(selectKotaByNama, nama).Scan(&kota.IdKota,
		&kota.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Kota{}, err
	}

	return kota, nil
}
