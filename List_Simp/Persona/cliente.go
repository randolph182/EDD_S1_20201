package Persona

type Cliente struct {
	Id     int
	Nombre string
	Edad   int
}

func New_Cliente(id int, nombre string, edad int) *Cliente {
	return &Cliente{id, nombre, edad}
}
