package main

import "fmt"

type Dog struct {
	Name string
}

func (d *Dog) WoofWoof() string {
	return fmt.Sprintf("Woof %s Woof\n", d.Name)
}

func main() {

	perrito := &Dog{"Hernán"}

	fmt.Println(perrito.WoofWoof())

	perrito = nil

	fmt.Println(perrito.WoofWoof())

}
