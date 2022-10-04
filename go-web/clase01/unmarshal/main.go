package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	type product struct {
		Name      string
		Price     float32
		Published bool
	}

	jsonData := []byte(`{"Name": "MacBook Air", "Price": 900, "Published": true}`)
	var p product

	if err := json.Unmarshal(jsonData, &p); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Producto: Name: %s", p.Name)
}
