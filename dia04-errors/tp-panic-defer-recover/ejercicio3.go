package main

import "fmt"

type Client struct {
	Legajo         int
	Nombre         string
	DNI            int
	NumeroTelefono int
	Domicilio      string
}

var (
	ErrInvalidLegajo      = fmt.Errorf("El legajo es inválido.")
	ErrInvalidDNI         = fmt.Errorf("El DNI es inválido.")
	ErrInvalidName        = fmt.Errorf("El nombre es inválido.")
	ErrInvalidPhoneNumber = fmt.Errorf("El número de teléfono es inválido.")
	ErrInvalidAddress     = fmt.Errorf("El domicilio es inválido.")
)

var Clientes = []Client{
	{
		0001, "Valentín", 42049202, 1523532323, "Roque 3333",
	},
	{
		0002, "Rodolfo", 30403943, 1523342323, "Boca 2042",
	},
	{
		0003, "Enrique", 23959194, 1539050029, "Mendoza 1234",
	},
}

func (c *Client) addToStorage() {
	_, err := checkNewClient(c)
	if err != nil {
		panic(err)
	}

	if existsClient(c.Legajo) {
		panic("Error: el cliente ya existe")
	}
}

func checkNewClient(c *Client) (result interface{}, err error) {
	if c.Legajo < 0 {
		result = c.Legajo
		err = ErrInvalidLegajo
		return
	}

	if c.DNI <= 0 {
		result = c.DNI
		err = ErrInvalidDNI
		return
	}

	if c.Domicilio == "" {
		result = c.Domicilio
		err = ErrInvalidAddress
		return
	}

	if c.Nombre == "" {
		result = c.Nombre
		err = ErrInvalidName
		return
	}

	if c.NumeroTelefono <= 0 {
		result = c.NumeroTelefono
		err = ErrInvalidPhoneNumber
		return
	}
	return
}

func existsClient(legajo int) (result bool) {
	for _, client := range Clientes {
		if client.Legajo == legajo {
			result = !result
			break
		}
	}
	return
}

// func main3_1() {
// 	newClient := Client{
// 		0001, "Mariano", 30504244, 1524025459, "Italia 234",
// 	}

// 	// file, err := os.ReadFile("customers.txt")

// 	defer func() {
// 		fmt.Println("Ejecución finalizada")
// 	}()

// 	newClient.addToStorage()
// }

func main() {
	// invalidLegajoClient := Client{
	// 	-2, "Almendra", 34252134, 152525324, "Bolivia 123",
	// }

	invalidNameClient := Client{
		23232, "", 34252134, 152525324, "Bolivia 123",
	}

	// validClient := Client{
	// 	0010, "Mariano", 30504244, 1524025459, "Italia 234",
	// }

	defer func() {
		fmt.Println("Ejecución finalizada")
		if err := recover(); err != nil {
			fmt.Printf("Hubo un error en la ejecución. \nMensaje de error: %v\n", err)
		}
	}()

	// invalidLegajoClient.addToStorage()
	invalidNameClient.addToStorage()
	// validClient.addToStorage()
}
