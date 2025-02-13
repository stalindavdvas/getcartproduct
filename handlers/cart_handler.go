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
		// User Default
		userID := "user:1"

		// Get products
		cartItems, err := client.HGetAll(ctx, userID).Result()
		if err != nil {
			http.Error(w, "Failure to get cart", http.StatusInternalServerError)
			return
		}

		// Parsear carts data
		var items []map[string]interface{}
		for productID, productJSON := range cartItems {
			// Decode JSON
			var productData map[string]interface{}
			err := json.Unmarshal([]byte(productJSON), &productData)
			if err != nil {
				http.Error(w, "Failure process data cart", http.StatusInternalServerError)
				return
			}

			// Add id cart
			productData["product_id"] = productID
			items = append(items, productData)
		}

		// Build response
		response := map[string]interface{}{
			"items": items,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
