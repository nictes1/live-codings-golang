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
	notebook := product{
		Name:      "MacBook Pro",
		Price:     1399.99,
		Published: true,
	}

	jsonData, err := json.Marshal(notebook)
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Printf("Producto: \n %s \n", string(jsonData))

}
