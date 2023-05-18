package web

type CategoryCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

type OrderCreateRequest struct {
	Customer_id int    `validate:"required" json:"customer_id"`
	Date        string `validate:"required" json:"date"`
	Status      string `validate:"required" json:"status"`
}
