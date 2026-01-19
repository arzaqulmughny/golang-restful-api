package repositories

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/helpers"
	"golang-restful-api/models/domains"
)

type CategoryImplementation struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryImplementation{}
}

func (repository *CategoryImplementation) Store(ctx context.Context, tx *sql.Tx, category domains.Category) domains.Category {
	sql := "INSERT INTO categories (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)
	helpers.PanicIfError(err)

	id, err := result.LastInsertId()

	helpers.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryImplementation) Update(ctx context.Context, tx *sql.Tx, id int, category domains.Category) domains.Category {
	sql := "UPDATE categories set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Name, id)

	helpers.PanicIfError(err)

	return category
}

func (repository *CategoryImplementation) Delete(ctx context.Context, tx *sql.Tx, category domains.Category) {
	sql := "DELETE FROM categories where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Id)

	helpers.PanicIfError(err)
}

func (repository *CategoryImplementation) FindAll(ctx context.Context, tx *sql.Tx) []domains.Category {
	sql := "SELECT id, name from categories"
	rows, err := tx.QueryContext(ctx, sql)
	helpers.PanicIfError(err)
	defer rows.Close()

	var categories []domains.Category
	for rows.Next() {
		category := domains.Category{}
		err := rows.Scan(&category.Id, &category.Name)

		if err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func (repository *CategoryImplementation) FindById(ctx context.Context, tx *sql.Tx, id int) (domains.Category, error) {
	sql := "SELECT id, name FROM categories where id = ?"
	row, err := tx.QueryContext(ctx, sql, id)
	helpers.PanicIfError(err)
	defer row.Close()

	category := domains.Category{}

	if !row.Next() {
		return category, errors.New("category is not found")
	} else {
		err := row.Scan(&category.Id, &category.Name)

		if err != nil {
			panic(err)
		}

		return category, nil
	}
}
