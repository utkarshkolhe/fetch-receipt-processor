package models

// A Struct to hold individual item on a  reciept
type ItemModel struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// A Struct to hold a reciept
type RecieptModel struct {
	Retailer     string      `json:"retailer"`
	PurchaseDate string      `json:"purchaseDate"`
	PurchaseTime string      `json:"purchaseTime"`
	Items        []ItemModel `json:"items"`
	Total        string      `json:"total"`
}
