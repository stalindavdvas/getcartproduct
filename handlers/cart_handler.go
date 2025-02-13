// handlers/cart_handler.go
package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func GetCart(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := "user:1" // Clave única para el carrito del usuario

		// Obtener todos los productos del carrito
		cartItems, err := client.HGetAll(ctx, userID).Result()
		if err != nil {
			http.Error(w, "Error al obtener los datos del carrito", http.StatusInternalServerError)
			return
		}

		// Parsear los datos del carrito
		items := make(map[string]map[string]interface{})
		for productID, productJSON := range cartItems {
			var productData map[string]interface{}
			err := json.Unmarshal([]byte(productJSON), &productData)
			if err != nil {
				http.Error(w, "Error al procesar los datos del carrito", http.StatusInternalServerError)
				return
			}
			items[productID] = productData
		}

		// Construir la respuesta
		response := map[string]interface{}{
			"items": items,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
