package main

import (
	"getcartproduct/database"
	"getcartproduct/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors" // Middleware para CORS
)

func main() {
	// Inicializar la conexión a Redis
	client := database.InitRedis()
	defer client.Close()

	// Crear un nuevo router
	r := mux.NewRouter()

	// Ruta para listar productos del carrito
	r.HandleFunc("/api/cart", handlers.GetCart(client)).Methods("GET")

	// Configurar CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://3.229.231.204:3000"},               // Permite solicitudes desde el frontend
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Métodos HTTP permitidos
		AllowedHeaders: []string{"Content-Type", "Authorization"},           // Encabezados permitidos
	})

	// Envolver el router con el middleware de CORS
	handler := corsHandler.Handler(r)

	// Iniciar el servidor en el puerto 8082
	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
