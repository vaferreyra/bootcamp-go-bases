package main

import (
	"fmt"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	Person   Person
	ID       int
	Position string
}

func (e Employee) PrintEmployee() {
	fmt.Printf("Empleado de nombre: %v, fecha de nacimiento: %v, ID de empleado: %v en posición: %v \n", e.Person.Name, e.Person.DateOfBirth, e.ID, e.Position)
}

func main() {
	personaValentin := Person{
		1, "Valentin", "8/6/1999",
	}

	empleado := Employee{
		Person:   personaValentin,
		ID:       2,
		Position: "Software Developer",
	}

	empleado.PrintEmployee()
}
