package Estructura

import (
	"fmt"

	"../Persona"
)

type Nodo struct {
	Siguiente *Nodo
	Persona   *Persona.Cliente
}

type Lista struct {
	Primero  *Nodo
	Ultimo   *Nodo
	Contador int
}

func New_Nodo(persona *Persona.Cliente) *Nodo {
	return &Nodo{nil, persona}
}

func New_Lista() *Lista {
	return &Lista{nil, nil, 0}
}

func Insertar(cliente *Persona.Cliente, lista *Lista) {

	var nuevo *Nodo = New_Nodo(cliente)

	if lista.Primero == nil {
		lista.Primero = nuevo
		lista.Ultimo = nuevo
		lista.Contador += 1
	} else {
		lista.Ultimo.Siguiente = nuevo
		lista.Ultimo = lista.Ultimo.Siguiente
		lista.Contador += 1
	}
}

func Imprimir(lista *Lista) {
	aux := lista.Primero

	for aux != nil {
		fmt.Println("Id:", aux.Persona.Id)
		fmt.Println("Nombre:", aux.Persona.Nombre)
		fmt.Println("Edad:", aux.Persona.Edad)
		aux = aux.Siguiente
	}

}
