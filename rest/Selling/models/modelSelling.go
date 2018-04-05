package models

type Selling struct {
	TblSellingID int    `json:"TblSellingID"`
	Invoice      string `json:"Invoice"`
	InvoiceDate  string `json:"InvoiceDate"`
	Item         int    `json:"Item"`
	Total        int    `json:"Total"`
	Paid         int    `json:"Paid"`
	Pengembalian int    `json:"Pengembalian"`
	OfficerCode  string `json:"OfficerCode"`
}
