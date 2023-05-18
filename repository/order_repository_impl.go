package repository

import (
	"assessment-golang-restful-api/helper"
	"assessment-golang-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, Order domain.Order) domain.Order {
	SQL := "insert into orders(customer_id, date, status) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, Order.Customer_id, Order.Date, Order.Status)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	Order.Id = int(id)
	return Order
}

func (repository *OrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, Order domain.Order) domain.Order {
	SQL := "update orders set customer_id = ?, date = ?, status = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, Order.Customer_id, Order.Date, Order.Status, Order.Id)
	helper.PanicIfError(err)

	return Order
}

func (repository *OrderRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, Order domain.Order) {
	SQL := "delete from orders where id = ?"
	_, err := tx.ExecContext(ctx, SQL, Order.Id)
	helper.PanicIfError(err)
}

func (repository *OrderRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, OrderId int) (domain.Order, error) {
	SQL := "select id, name from orders where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, OrderId)
	helper.PanicIfError(err)
	defer rows.Close()

	Order := domain.Order{}
	if rows.Next() {
		err := rows.Scan(&Order.Id, &Order.Customer_id, &Order.Date, &Order.Status)
		helper.PanicIfError(err)
		return Order, nil
	} else {
		return Order, errors.New("Order is not found")
	}
}

func (repository *OrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Order {
	SQL := "select * from orders"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		Order := domain.Order{}
		err := rows.Scan(&Order.Id, &Order.Customer_id, &Order.Date, &Order.Status)
		helper.PanicIfError(err)
		orders = append(orders, Order)
	}
	return orders
}
