package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
