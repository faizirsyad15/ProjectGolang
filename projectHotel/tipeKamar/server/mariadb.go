package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addTipeKamar = `insert into mst_tipe_kamar(id_tipe_kamar,nama_tipe_kamar,harga_kamar,status,created_by,
		created_on)values (?,?,?,?,?,?)`
	selectTipeKamarByHarga = `select id_tipe_kamar,nama_tipe_kamar,status
		from mst_tipe_kamar where harga_kamar = ?`
	selectTipeKamar = `select id_tipe_kamar,nama_tipe_kamar,harga_kamar,status 
		from mst_tipe_kamar`
	updateTipeKamar = `update mst_tipe_kamar set nama_tipe_kamar=?,harga_kamar=?,status=?,
		updated_by=?,updated_on=? where id_tipe_kamar=?`
	selectTipeKamarByNama = `select id_tipe_kamar,harga_kamar,status
		from mst_tipe_kamar where nama_tipe_kamar=?`
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

func (rw *dbReadWriter) AddTipeKamar(kamar TipeKamar) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addTipeKamar, kamar.IdTipeKamar, kamar.NamaTipeKamar, kamar.HargaKamar,
		OnAdd, OnAdd2, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadTipeKamarByHarga(harga int32) (TipeKamar, error) {
	kamar := TipeKamar{HargaKamar: harga}
	err := rw.db.QueryRow(selectTipeKamarByHarga, harga).Scan(&kamar.IdTipeKamar, &kamar.NamaTipeKamar,
		&kamar.Status)

	if err != nil {
		return TipeKamar{}, err
	}

	return kamar, nil
}

func (rw *dbReadWriter) ReadTipeKamar() (TipeKamars, error) {
	kamar := TipeKamars{}
	rows, _ := rw.db.Query(selectTipeKamar)
	defer rows.Close()
	for rows.Next() {
		var t TipeKamar
		err := rows.Scan(&t.IdTipeKamar, &t.NamaTipeKamar, &t.HargaKamar, &t.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return kamar, err
		}
		kamar = append(kamar, t)
	}
	//fmt.Println("db nya:", customer)
	return kamar, nil
}

func (rw *dbReadWriter) UpdateTipeKamar(tip TipeKamar) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateTipeKamar, tip.NamaTipeKamar, tip.HargaKamar, tip.Status,
		OnAdd3, time.Now(), tip.IdTipeKamar)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadTipeKamarByNama(nama string) (TipeKamar, error) {
	kamar := TipeKamar{NamaTipeKamar: nama}
	err := rw.db.QueryRow(selectTipeKamarByNama, nama).Scan(&kamar.IdTipeKamar,
		&kamar.HargaKamar, &kamar.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return TipeKamar{}, err
	}

	return kamar, nil
}
