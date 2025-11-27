package main

import (
	"fmt"
	"log"
	"os"

	"github.com/brunoapolinariodev/montink_erp/internal/montink"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar arquivo .env")
	}

	token := os.Getenv("MONTINK_TOKEN")
	if token == "" {
		fmt.Println("Token vazio")
	}

	fmt.Println("Iniciando teste de integração com a Montink...")

	client := montink.NewClient(token)

	orders, err := client.GetOrders()
	if err != nil {
		log.Fatalf("Erro ao buscar pedidos: %v", err)
	}

	fmt.Printf("Sucesso! Encontrados %d pedidos.\n", len(orders))
	fmt.Println("------------------------------------------------")

	for _, order := range orders {
		fmt.Printf("Pedido #%s | Cliente: %s %s | Status: %s | Valor: %s\n",
			order.ID,
			order.FirstName,
			order.LastName,
			order.Status,
			order.SalesValueStr,
		)
	}
}
