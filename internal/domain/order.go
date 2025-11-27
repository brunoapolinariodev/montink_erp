package domain

import "strconv"

// Order represents a purchase made in the Montink Store
// It maps the inconsistent JSON fields returned by the API
type Order struct {
	ID            string  `json:"referencia"`
	Status        string  `json:"status"`
	FirstName     string  `json:"nome"`
	LastName      string  `json:"sobrenome"`
	OrderDate     string  `json:"data_pedido_realizado"`
	SalesValueStr string  `json:"valor_venda"`
	ProfitStr     string  `json:"lojista_lucro"`
	PaymentStatus string  `json:"pagamento_status"`
	PaymentMethod string  `json:"forma_pagamento"`
	Cost          float64 `json:"custo"`
	CustomerName  string  `json:"nome_cliente"`
}

func (o Order) SalesValue() float64 {
	val, err := strconv.ParseFloat(o.SalesValueStr, 64)
	if err != nil {
		return 0.0
	}
	return val
}

// MontinkOrderResponse represents the JSON response from the Montink API and wraps order on Orders slice
type MontinkOrderResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"msg"`
	Orders  []Order `json:"pedidos"`
	Filters struct {
		PageLimit int `json:"limite"`
	} `json:"filtros"`
}
