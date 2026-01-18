package repositories

import (
	"context"
	"database/sql"
	"golang-restful-api/models/domains"
)

type CategoryRepository interface {
	Store(ctx context.Context, tx *sql.Tx, category domains.Category) domains.Category
	Update(ctx context.Context, tx *sql.Tx, category domains.Category) domains.Category
	Delete(ctx context.Context, tx *sql.Tx, category domains.Category)
	FindAll(ctx context.Context, tx *sql.Tx) []domains.Category
	FindById(ctx context.Context, tx *sql.Tx, id int) (domains.Category, error)
}
