package model

type InvoiceRequest struct {
	Id            string  `json:"id"`
	OrderID       string  `json:"orderId"`
	TransactionID string  `json:"transactionId"`
	Paid          Paid    `json:"paid"`
	Ship          Ship    `json:"ship"`
	Items         []Item  `json:"items"`
	TotalAmount   float64 `json:"totalAmount"`
	Currency      string  `json:"currency"`
}

type Paid struct {
	Card struct {
		Name   string `json:"name"`
		Number string `json:"number"`
	} `json:"card"`
	Name string `json:"name"`
}

type Ship struct {
	Name    string `json:"name"`
	Address struct {
		Line1   string `json:"line1"`
		Line2   string `json:"line2"`
		Country string `json:"country"`
	} `json:"address"`
}

type Item struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}
