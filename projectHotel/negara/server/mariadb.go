package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addNegara = `insert into mst_negara(id_negara,nama_negara,status,created_by,
		created_on)values (?,?,?,?,?)`
	selectNegara = `select id_negara,nama_negara,status 
		from mst_negara`
	updateNegara = `update mst_negara set nama_negara=?,status=?,
		updated_by=?,updated_on=? where id_negara=?`
	selectNegaraByNama = `select id_negara,nama_negara,status
		from mst_negara where nama_negara=?`
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

func (rw *dbReadWriter) AddNegara(negara Negara) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addNegara, negara.IdNegara, negara.NamaNegara,
		OnAdd, OnAdd2, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadNegara() (Negaras, error) {
	negara := Negaras{}
	rows, _ := rw.db.Query(selectNegara)
	defer rows.Close()
	for rows.Next() {
		var t Negara
		err := rows.Scan(&t.IdNegara, &t.NamaNegara, &t.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return negara, err
		}
		negara = append(negara, t)
	}
	//fmt.Println("db nya:", customer)
	return negara, nil
}

func (rw *dbReadWriter) UpdateNegara(neg Negara) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateNegara, neg.NamaNegara, neg.Status,
		OnAdd3, time.Now(), neg.IdNegara)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadNegaraByNama(nama string) (Negara, error) {
	negara := Negara{NamaNegara: nama}
	err := rw.db.QueryRow(selectNegaraByNama, nama).Scan(&negara.IdNegara,
		&negara.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Negara{}, err
	}

	return negara, nil
}
