package web

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type OrderResponse struct {
	Id          int    `json:"id"`
	Customer_id int    `json:"customer_id"`
	Date        string `json:"date"`
	Status      string `json:"status"`
}
