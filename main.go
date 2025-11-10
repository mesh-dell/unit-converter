package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mesh-dell/unit-converter/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.LengthHandler)
	http.HandleFunc("/weight", handlers.WeightHandler)
	http.HandleFunc("/temperature", handlers.TemperatureHandler)
	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
