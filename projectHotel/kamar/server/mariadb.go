package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addKamar = `insert into mst_kamar(id_kamar,id_tipe_kamar,id_menu,status,created_by,
		created_on)values (?,?,?,?,?,?)`
	selectKamar = `select id_kamar,id_tipe_kamar,id_menu,status 
		from mst_kamar`
	updateKamar = `update mst_kamar set id_tipe_kamar=?,id_menu=?,status=?,
		updated_by=?,updated_on=? where id_kamar=?`
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

func (rw *dbReadWriter) AddKamar(kamar Kamar) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addKamar, kamar.IdKamar, kamar.IdTipeKamar, kamar.IdMenu,
		OnAdd, OnAdd2, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadKamar() (Kamars, error) {
	kamar := Kamars{}
	rows, _ := rw.db.Query(selectKamar)
	defer rows.Close()
	for rows.Next() {
		var t Kamar
		err := rows.Scan(&t.IdKamar, &t.IdTipeKamar, &t.IdMenu, &t.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return kamar, err
		}
		kamar = append(kamar, t)
	}
	//fmt.Println("db nya:", customer)
	return kamar, nil
}

func (rw *dbReadWriter) UpdateKamar(tip Kamar) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateKamar, tip.IdTipeKamar, tip.IdMenu, tip.Status,
		OnAdd3, time.Now(), tip.IdKamar)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}
