package repositories

import (
	"database/sql"
	"errors"
)

func CreateOperator(db *sql.DB) {
}

// Produk
// menambahkan produk
func CreateProduct(db *sql.DB, sku, name, productType, unit string, cost float64) error {

	if productType != "finished" && productType != "raw" && productType != "semi-finished" {
		return errors.New("type harus: finished, raw, atau semi-finished")
	}

	query := `
	INSERT INTO products (sku, name, type, unit, standard_cost)
	VALUES (?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query, sku, name, productType, unit, cost)
	if err != nil {
		return err
	}

	return nil
}

type Product struct {
	ID        int
	SKU       string
	Name      string
	Type      string
	Unit      string
	Cost      float64
	CreatedAt string
	UpdatedAt string
}

func ListProducts(db *sql.DB) ([]Product, error) {
	query := `
		SELECT product_id, sku, name, type, unit, standard_cost, created_at, updated_at
		FROM products
		ORDER BY product_id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product

		err := rows.Scan(
			&p.ID,
			&p.SKU,
			&p.Name,
			&p.Type,
			&p.Unit,
			&p.Cost,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, rows.Err()
}

// update product
func UpdateProduct(db *sql.DB, id int, sku, name, productType, unit string, cost float64) error {
	query := `
		UPDATE products 
		SET sku = ?, name = ?, type = ?, unit = ?, standard_cost = ?, updated_at = CURRENT_TIMESTAMP
		WHERE product_id = ?
	`

	_, err := db.Exec(query, sku, name, productType, unit, cost, id)
	return err
}

// delete product
func DeleteProduct(db *sql.DB, id int) (int64, error) {
	query := `DELETE FROM products WHERE product_id = ?`

	result, err := db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
