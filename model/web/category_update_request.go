package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}

type OrderUpdateRequest struct {
	Id          int    `validate:"required"`
	Customer_id int    `validate:"required" json:"customer_id"`
	Date        string `validate:"required" json:"date"`
	Status      string `validate:"required" json:"status"`
}
