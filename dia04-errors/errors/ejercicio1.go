package main

import "fmt"

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return "Error: el salario ingresado no alcanza el mínimo imponible"
}

func main1() {
	err := &MyError{
		msg: "Error message",
	}

	var input float64
	fmt.Scan(&input)

	if input < 150000 {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Debe pagar impuesto")
}
