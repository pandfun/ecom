package order

import (
	"database/sql"

	"github.com/pandfun/ecom/types"
)

type OrderStore struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *OrderStore {
	return &OrderStore{db: db}
}

func (s *OrderStore) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec("INSERT INTO orders (userId, total, status, address) VALUES (?, ?, ?, ?)", order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (s *OrderStore) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (orderId, productID, quantity, price) VALUES (?, ?, ?, ?)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}
