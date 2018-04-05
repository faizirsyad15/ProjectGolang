package models

type Item struct {
	TblItemID    int    `json:"TblItemID"`
	ItemCode     string `json:"ItemCode"`
	ItemName     string `json:"ItemName"`
	BuyingPrice  string `json:"BuyingPrice"`
	SellingPrice string `json:"SellingPrice"`
	ItemAmount   int    `json:"ItemAmount"`
	Pieces       string `json:"Pieces"`
}
