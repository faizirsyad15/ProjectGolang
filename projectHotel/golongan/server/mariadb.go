package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addGolongan = `insert into mst_golongan(id_golongan,nama_golongan,status,created_by,
		created_on)values (?,?,?,?,?)`
	selectGolongan = `select id_golongan,nama_golongan,status 
		from mst_golongan`
	updateGolongan = `update mst_golongan set nama_golongan=?,status=?,
		updated_by=?,updated_on=? where id_golongan=?`
	selectGolonganByNama = `select id_golongan,nama_golongan,status
		from mst_golongan where nama_golongan=?`
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

func (rw *dbReadWriter) AddGolongan(golongan Golongan) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addGolongan, golongan.IdGolongan, golongan.NamaGolongan,
		OnAdd, OnAdd2, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadGolongan() (Golongans, error) {
	golongan := Golongans{}
	rows, _ := rw.db.Query(selectGolongan)
	defer rows.Close()
	for rows.Next() {
		var t Golongan
		err := rows.Scan(&t.IdGolongan, &t.NamaGolongan, &t.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return golongan, err
		}
		golongan = append(golongan, t)
	}
	//fmt.Println("db nya:", customer)
	return golongan, nil
}

func (rw *dbReadWriter) UpdateGolongan(gol Golongan) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateGolongan, gol.NamaGolongan, gol.Status,
		OnAdd3, time.Now(), gol.IdGolongan)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadGolonganByNama(nama string) (Golongan, error) {
	golongan := Golongan{NamaGolongan: nama}
	err := rw.db.QueryRow(selectGolonganByNama, nama).Scan(&golongan.IdGolongan,
		&golongan.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Golongan{}, err
	}

	return golongan, nil
}
