package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/brunoapolinariodev/montink_erp/internal/domain"
	_ "modernc.org/sqlite"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(dsn string) (*SQLiteRepository, error) {

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id TEXT PRIMARY KEY,
		status TEXT,
		first_name TEXT,
		last_name TEXT,
		order_date TEXT,
		sales_value REAL,
		cost REAL
	);`

	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &SQLiteRepository{db: db}, nil
}

func (r *SQLiteRepository) Save(ctx context.Context, order *domain.Order) error {
	query := `
	INSERT OR REPLACE INTO orders (id, status, first_name, last_name, order_date, sales_value, cost)
	VALUES (?, ?, ?, ?, ?, ?, ?);
	`

	_, err := r.db.ExecContext(ctx, query,
		order.ID,
		order.Status,
		order.FirstName,
		order.LastName,
		order.OrderDate,
		order.SalesValue(), // Salvamos o número float já convertido!
		order.Cost,
	)
	return err
}

func (r *SQLiteRepository) List(ctx context.Context) ([]domain.Order, error) {
	query := `SELECT id, status, first_name, last_name, order_date, sales_value, cost FROM orders`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var o domain.Order
		var salesValue float64

		err := rows.Scan(
			&o.ID, &o.Status, &o.FirstName, &o.LastName, &o.OrderDate, &salesValue, &o.Cost,
		)
		if err != nil {
			return nil, err
		}

		// Truque: Como o banco guarda float, mas nossa struct original recebia String da API,
		// precisamos preencher o campo string manualmente para manter compatibilidade,
		// ou apenas aceitar que o dado veio do banco.
		// Para simplificar, vou formatar de volta para string caso precisemos exibir.
		o.SalesValueStr = fmt.Sprintf("%.2f", salesValue)

		orders = append(orders, o)
	}
	return orders, nil
}

func (r *SQLiteRepository) GetByID(ctx context.Context, id string) (*domain.Order, error) {
	query := `SELECT id, status, first_name, last_name, order_date, sales_value, cost FROM orders WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)

	var o domain.Order
	var salesValue float64

	err := row.Scan(
		&o.ID, &o.Status, &o.FirstName, &o.LastName, &o.OrderDate, &salesValue, &o.Cost,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}
	o.SalesValueStr = fmt.Sprintf("%.2f", salesValue)

	return &o, nil
}
