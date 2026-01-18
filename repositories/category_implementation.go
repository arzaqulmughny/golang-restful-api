package repositories

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/models/domains"
)

type CategoryImplementation struct {
}

func (repository *CategoryImplementation) Store(ctx context.Context, tx *sql.Tx, category domains.Category) domains.Category {
	sql := "INSERT INTO categories (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	category.Id = int(id)
	return category
}

func (repository *CategoryImplementation) Update(ctx context.Context, tx *sql.Tx, category domains.Category) domains.Category {
	sql := "UPDATE categories set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)

	if err != nil {
		panic(err)
	}

	return category
}

func (repository *CategoryImplementation) Delete(ctx context.Context, tx *sql.Tx, category domains.Category) {
	sql := "DELETE FROM categories where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Id)

	if err != nil {
		panic(err)
	}
}

func (repository *CategoryImplementation) FindAll(ctx context.Context, tx *sql.Tx) []domains.Category {
	sql := "SELECT id, name from categories"
	rows, err := tx.QueryContext(ctx, sql)

	if err != nil {
		panic(err)
	}

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
	sql := "SELECT id, name where id = ?"
	row, err := tx.QueryContext(ctx, sql, id)

	category := domains.Category{}

	if err != nil {
		panic(err)
	}

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
