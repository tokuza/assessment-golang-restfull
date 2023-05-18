package repository

import (
	"assessment-golang-restful-api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}

type OrderRepository interface {
	Save(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	Update(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	Delete(ctx context.Context, tx *sql.Tx, order domain.Order)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Order, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Order
}
