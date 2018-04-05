package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addJenisPembayaran = `insert into mst_jenis_pembayaran(id_jenis_pembayaran,metode_pembayaran,status,created_by,
		created_on)values (?,?,?,?,?)`
	selectJenisPembayaran = `select id_jenis_pembayaran,metode_pembayaran,status 
		from mst_jenis_pembayaran`
	updateJenisPembayaran = `update mst_jenis_pembayaran set metode_pembayaran=?,status=?,
		updated_by=?,updated_on=? where id_jenis_pembayaran=?`
	selectJenisPembayaranByMetode = `select id_jenis_pembayaran,metode_pembayaran,status
		from mst_jenis_pembayaran where metode_pembayaran=?`
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

func (rw *dbReadWriter) AddJenisPembayaran(jenispembayaran JenisPembayaran) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addJenisPembayaran, jenispembayaran.IdJenisPembayaran, jenispembayaran.MetodePembayaran,
		OnAdd, OnAdd2, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadJenisPembayaran() (JenisPembayarans, error) {
	jenispembayaran := JenisPembayarans{}
	rows, _ := rw.db.Query(selectJenisPembayaran)
	defer rows.Close()
	for rows.Next() {
		var t JenisPembayaran
		err := rows.Scan(&t.IdJenisPembayaran, &t.MetodePembayaran, &t.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return jenispembayaran, err
		}
		jenispembayaran = append(jenispembayaran, t)
	}
	//fmt.Println("db nya:", customer)
	return jenispembayaran, nil
}

func (rw *dbReadWriter) UpdateJenisPembayaran(jen JenisPembayaran) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateJenisPembayaran, jen.MetodePembayaran, jen.Status,
		OnAdd3, time.Now(), jen.IdJenisPembayaran)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadJenisPembayaranByMetode(metode string) (JenisPembayaran, error) {
	jenispembayaran := JenisPembayaran{MetodePembayaran: metode}
	err := rw.db.QueryRow(selectJenisPembayaranByMetode, metode).Scan(&jenispembayaran.IdJenisPembayaran,
		&jenispembayaran.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return JenisPembayaran{}, err
	}

	return jenispembayaran, nil
}
