package main

import (
	"./Estructura"
	"./Persona"
)

func main() {
	var cl1 *Persona.Cliente = Persona.New_Cliente(1, "cliente1", 20)
	var cl2 *Persona.Cliente = Persona.New_Cliente(2, "cliente2", 22)
	var cl3 *Persona.Cliente = Persona.New_Cliente(3, "cliente3", 23)

	var lista *Estructura.Lista = Estructura.New_Lista()

	Estructura.Insertar(cl1, lista)
	Estructura.Insertar(cl2, lista)
	Estructura.Insertar(cl3, lista)

	Estructura.Imprimir(lista)

}
