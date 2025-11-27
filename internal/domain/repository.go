package domain

import "context"

// OrderRepository represents the contract for the Order persistence layer
type OrderRepository interface {
	Save(ctx context.Context, order *Order)
	List(ctx context.Context) ([]Order, error)
	GetByID(ctx context.Context, id string) (*Order, error)
}
