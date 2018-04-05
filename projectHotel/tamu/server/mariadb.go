package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addTamu = `insert into mst_tamu(id_tamu,nama_tamu,no_telepon,jenis_kelamin,id_alamat,status,created_by,
		created_on)values (?,?,?,?,?,?,?,?)`
	selectTamuByTelepon = `select id_tamu,nama_tamu,jenis_kelamin,id_alamat,status
		from mst_tamu where no_telepon = ?`
	selectTamu = `select id_tamu,nama_tamu,no_telepon,jenis_kelamin,id_alamat,status 
		from mst_tamu`
	updateTamu = `update mst_tamu set nama_tamu=?,no_telepon=?,jenis_kelamin=?,id_alamat=?,status=?,
		updated_by=?,updated_on=? where id_tamu=?`
	selectTamuByNama = `select id_tamu,no_telepon,jenis_kelamin,id_alamat,status
		from mst_tamu where nama_tamu=?`
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

func (rw *dbReadWriter) AddTamu(tamu Tamu) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addTamu, tamu.IdTamu, tamu.NamaTamu, tamu.NoTelepon, tamu.JenisKelamin, tamu.IdAlamat,
		OnAdd, OnAdd2, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadTamuByTelepon(telepon string) (Tamu, error) {
	tamu := Tamu{NoTelepon: telepon}
	err := rw.db.QueryRow(selectTamuByTelepon, telepon).Scan(&tamu.IdTamu, &tamu.NamaTamu,
		&tamu.JenisKelamin, &tamu.IdAlamat, &tamu.Status)

	if err != nil {
		return Tamu{}, err
	}

	return tamu, nil
}

func (rw *dbReadWriter) ReadTamu() (Tamus, error) {
	tamu := Tamus{}
	rows, _ := rw.db.Query(selectTamu)
	defer rows.Close()
	for rows.Next() {
		var t Tamu
		err := rows.Scan(&t.IdTamu, &t.NamaTamu, &t.NoTelepon, &t.JenisKelamin, &t.IdAlamat, &t.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return tamu, err
		}
		tamu = append(tamu, t)
	}
	//fmt.Println("db nya:", customer)
	return tamu, nil
}

func (rw *dbReadWriter) UpdateTamu(tam Tamu) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateTamu, tam.NamaTamu, tam.NoTelepon, tam.JenisKelamin, tam.IdAlamat, tam.Status,
		OnAdd3, time.Now(), tam.IdTamu)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadTamuByNama(nama string) (Tamu, error) {
	tamu := Tamu{NamaTamu: nama}
	err := rw.db.QueryRow(selectTamuByNama, nama).Scan(&tamu.IdTamu,
		&tamu.NoTelepon, &tamu.JenisKelamin, &tamu.IdAlamat, &tamu.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Tamu{}, err
	}

	return tamu, nil
}
