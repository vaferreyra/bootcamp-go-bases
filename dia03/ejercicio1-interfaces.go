package main

import "fmt"

type Alumno struct {
	Nombre, Apellido string
	DNI              int
	FechaIngreso     string
}

func (a Alumno) String() string {
	return fmt.Sprintf(
		"Nombre: %s\nApellido: %s\nDNI: %v\nFechaIngreso: %s\n",
		a.Nombre,
		a.Apellido,
		a.DNI,
		a.FechaIngreso,
	)
}

// func main() {
// 	var stringAlumno fmt.Stringer = Alumno{
// 		"Valentin", "Ferreyra", 41924003, "19/12/2022",
// 	}

// 	fmt.Println(stringAlumno)
// }
