package main

import "fmt"

type Domicilio struct {
	Calle  string `json:"calle"`
	Altura int    `json:"altura"`
}

type Persona struct {
	Nombre     string    `json:"nombre"`
	Edad       int       `json:"edad"`
	Residencia Domicilio `json:"domicilio"`
}

func (p *Persona) cumplirAños() {
	p.Edad++
}

func main() {

	var persona Persona

	persona.Nombre = "Valentin"
	persona.Edad = 23
	persona.Residencia = Domicilio{Calle: "Brahms", Altura: 5732}

	fmt.Println(persona.Edad)
	persona.cumplirAños()
	fmt.Println(persona.Edad)

}
