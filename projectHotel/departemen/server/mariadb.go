package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addDepartemen = `insert into mst_departemen(id_departemen,nama_departemen,status,created_by,
		created_on)values (?,?,?,?,?)`
	selectDepartemen = `select id_departemen,nama_departemen,status 
		from mst_departemen`
	updateDepartemen = `update mst_departemen set nama_departemen=?,status=?,
		updated_by=?,updated_on=? where id_departemen=?`
	selectDepartemenByNama = `select id_departemen,nama_departemen,status
		from mst_departemen where nama_departemen=?`
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

func (rw *dbReadWriter) AddDepartemen(departemen Departemen) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addDepartemen, departemen.IdDepartemen, departemen.NamaDepartemen,
		OnAdd, OnAdd2, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadDepartemen() (Departemens, error) {
	departemen := Departemens{}
	rows, _ := rw.db.Query(selectDepartemen)
	defer rows.Close()
	for rows.Next() {
		var t Departemen
		err := rows.Scan(&t.IdDepartemen, &t.NamaDepartemen, &t.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return departemen, err
		}
		departemen = append(departemen, t)
	}
	//fmt.Println("db nya:", customer)
	return departemen, nil
}

func (rw *dbReadWriter) UpdateDepartemen(dep Departemen) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateDepartemen, dep.NamaDepartemen, dep.Status,
		OnAdd3, time.Now(), dep.IdDepartemen)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadDepartemenByNama(nama string) (Departemen, error) {
	departemen := Departemen{NamaDepartemen: nama}
	err := rw.db.QueryRow(selectDepartemenByNama, nama).Scan(&departemen.IdDepartemen,
		&departemen.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Departemen{}, err
	}

	return departemen, nil
}
