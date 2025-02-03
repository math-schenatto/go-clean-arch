package database

import (
	"database/sql"
	"fmt"

	"github.com/math-schenatto/go-clean-arch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) FindAll(page, limit int) ([]entity.Order, error) {
	fmt.Println("está no database")
	var orders []entity.Order
	query := "SELECT id, price, tax, final_price FROM orders"
	var args []interface{}

	if page > 0 && limit > 0 {
		query += " LIMIT ? OFFSET ?"
		args = append(args, limit, (page-1)*limit)
	}

	rows, err := r.Db.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order entity.Order
		if err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
