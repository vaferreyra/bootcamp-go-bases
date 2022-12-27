package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{1, "Heladera", 104000.0, "Heladera grande", "Electrodoméstico"},
	{2, "Celular", 60000.0, "Motorola g42", "Tecnología"},
	{3, "Computadora", 200000.0, "Compu Gamer", "Tecnología"},
	{4, "Televisor", 102000.0, "SmartTV 72''", "Tecnologia"},
	{5, "Lapicera", 30.0, "Bic azul", "Utiles"},
}

func (p Product) Save() {
	Products = append(Products, p)
}

func GetAllProducts() {
	for _, product := range Products {
		fmt.Printf("Producto: %s con precio: %f y categoria: %s \n", product.Name, product.Price, product.Category)
	}
}

func GetById(id int) (producto Product, err error) {
	for _, prod := range Products {
		if prod.ID == id {
			return prod, nil
		}
	}
	return Product{}, errors.New("There is not product with that id")
}

// func main() {
// 	product := Product{6, "Alimento para perro", 2300.0, "Dog chow", "Animales"}

// 	fmt.Println("----------- Productos iniciales -----------")
// 	GetAllProducts()

// 	// Add new product
// 	product.Save()
// 	fmt.Println("----------- Agrego producto alimento -----------")
// 	GetAllProducts()

// 	// Get by id
// 	fmt.Println("----------- Get product by id -----------")
// 	productById, err := GetById(3)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(productById)
// 	}

// 	fmt.Println("----------- Error in get product by id -----------")
// 	productById2, err2 := GetById(100)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 	} else {
// 		fmt.Println(productById2)
// 	}
// }
