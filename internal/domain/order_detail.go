package domain

// OrderDetail represents the full JSON response from the "get order" endpoint of the Montink API.
// It maps the hierarchical structure where customer data is separated from product items.
type OrderDetail struct {
	Success       bool        `json:"success"`
	Message       string      `json:"msg"`
	PaymentMethod string      `json:"forma_pagamento"`
	TrackingLink  string      `json:"linkRastreio"`
	Data          OrderData   `json:"carrinho"`         // Maps the "carrinho" object containing customer info
	Items         []OrderItem `json:"carrinhoProdutos"` // Maps the list of items
}

// OrderData represents the customer and general order information found in the "carrinho" object.
type OrderData struct {
	ID        string `json:"id"`
	FirstName string `json:"nome"`
	LastName  string `json:"sobrenome"`
	Email     string `json:"email"`
	Phone     string `json:"telefone"`
	Status    string `json:"status"` // e.g., "Devolução", "Pago"

	// Address information
	Address string `json:"endereco"`
	Number  string `json:"numero"`
	City    string `json:"cidade"`
	State   string `json:"estado"`
	ZipCode string `json:"cep"`

	// Financial details (returned as strings by the API)
	TotalValue    string `json:"valor_venda"`
	ShippingValue string `json:"valor_frete"`
	PaymentStatus string `json:"pagamento_status"`
	OrderDate     string `json:"data_pedido_realizado"`
}

// OrderItem represents a single item within the order list.
// The API separates the static product info ("produto") from the specific selection ("carrinho_produto").
type OrderItem struct {
	Product   ProductInfo   `json:"produto"`
	Selection SelectionInfo `json:"carrinho_produto"`
}

// ProductInfo contains the static details of the product (Name, Image, Base ID).
type ProductInfo struct {
	ID     string `json:"product_id"`
	Name   string `json:"product_name"` // Sometimes "nomeShopify"
	Image  string `json:"img"`
	Handle string `json:"handle"`
}

// SelectionInfo contains the specific choices made by the customer for this item (Size, Color, Quantity).
type SelectionInfo struct {
	Quantity string `json:"quantidade"`
	Size     string `json:"variant1"` // Maps to "variant1" (e.g., "GG")
	Color    string `json:"variant2"` // Maps to "variant2" (e.g., "Preto")
	Price    string `json:"valor"`
}
