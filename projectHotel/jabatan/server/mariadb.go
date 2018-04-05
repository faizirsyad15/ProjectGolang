package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addJabatan = `insert into mst_jabatan(id_jabatan,nama_jabatan,status,created_by,
		created_on)values (?,?,?,?,?)`
	selectJabatan = `select id_jabatan,nama_jabatan,status 
		from mst_jabatan`
	updateJabatan = `update mst_jabatan set nama_jabatan=?,status=?,
		updated_by=?,updated_on=? where id_jabatan=?`
	selectJabatanByNama = `select id_jabatan,nama_jabatan,status
		from mst_jabatan where nama_jabatan=?`
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

func (rw *dbReadWriter) AddJabatan(jabatan Jabatan) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addJabatan, jabatan.IdJabatan, jabatan.NamaJabatan,
		OnAdd, OnAdd2, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadJabatan() (Jabatans, error) {
	jabatan := Jabatans{}
	rows, _ := rw.db.Query(selectJabatan)
	defer rows.Close()
	for rows.Next() {
		var t Jabatan
		err := rows.Scan(&t.IdJabatan, &t.NamaJabatan, &t.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return jabatan, err
		}
		jabatan = append(jabatan, t)
	}
	//fmt.Println("db nya:", customer)
	return jabatan, nil
}

func (rw *dbReadWriter) UpdateJabatan(jab Jabatan) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateJabatan, jab.NamaJabatan, jab.Status,
		OnAdd3, time.Now(), jab.IdJabatan)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadJabatanByNama(nama string) (Jabatan, error) {
	jabatan := Jabatan{NamaJabatan: nama}
	err := rw.db.QueryRow(selectJabatanByNama, nama).Scan(&jabatan.IdJabatan,
		&jabatan.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Jabatan{}, err
	}

	return jabatan, nil
}
