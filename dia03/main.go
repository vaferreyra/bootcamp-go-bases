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

type Author struct {
	FirstName, LastName string
}

func (a *Author) String() string {
	return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
}

type Book struct {
	Title, Content string
	Author         Author
}

func (b Book) String() string {
	return fmt.Sprintf(
		"El libro %s fue escrito por %s y contiene %d caracteres",
		b.Title,
		b.Author.String(),
		len(b.Content),
	)
}

func increment(value *int) {
	*value++
}

// func main() {
// 	author := Author{
// 		"John", "Doe",
// 	}

// 	var textDescriptionGenerator fmt.Stringer = Book{
// 		Title:   "Lord of the Rings",
// 		Content: "Lorem ipsum",
// 		Author:  author,
// 	}

// 	fmt.Println(textDescriptionGenerator.String())
// 	// number := 10
// 	// fmt.Println(number)

// 	// increment(&number)

// 	// fmt.Println(number)

// 	// author := Author{
// 	// 	"John", "Doe",
// 	// }

// 	// book := Book{
// 	// 	Title:   "Lord of the Rings",
// 	// 	Content: "Lorem ipsum",
// 	// 	Author:  author,
// 	// }

// 	// book.PrintInformation()

// 	// var persona Persona

// 	// persona.Nombre = "Valentin"
// 	// persona.Edad = 23
// 	// persona.Residencia = Domicilio{Calle: "Brahms", Altura: 5732}

// 	// fmt.Println(persona.Edad)
// 	// persona.cumplirAños()
// 	// fmt.Println(persona.Edad)

// }
