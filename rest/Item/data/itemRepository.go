package data

import (
	"database/sql"
	"day15/Item/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type ItemRepository struct {
	DB *sql.DB
}

// untuk nilai return get all butuh struktur dari pengarang
// 2.a buat model dari Pengarang

func GetAll(db *ItemRepository) []models.Item {
	fmt.Println(db.DB)

	result, err := db.DB.Query("Select * From tblItem")

	if err != nil {
		return nil
	}

	defer result.Close()
	fmt.Println(result)
	item := []models.Item{}
	for result.Next() {
		var p models.Item
		if err := result.Scan(&p.TblItemID, &p.ItemCode, &p.ItemName, &p.BuyingPrice, &p.SellingPrice, &p.ItemAmount, &p.Pieces); err != nil {
			return nil
		}
		item = append(item, p)
	}
	return item
}
