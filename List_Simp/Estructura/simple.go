package Estructura

import (
	"fmt"

	"../Persona"
)

type Nodo struct {
	siguiente *Nodo
	persona   *Persona.Cliente
}

type Lista struct {
	primero  *Nodo
	ultimo   *Nodo
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

	if lista.primero == nil {
		lista.primero = nuevo
		lista.ultimo = nuevo
		lista.Contador += 1
	} else {
		lista.ultimo.siguiente = nuevo
		lista.ultimo = lista.ultimo.siguiente
		lista.Contador += 1
	}
}

func Imprimir(lista *Lista) {
	aux := lista.primero

	for aux != nil {
		fmt.Println("Id:", aux.persona.Id)
		fmt.Println("Nombre:", aux.persona.Nombre)
		fmt.Println("Edad:", aux.persona.Edad)
		aux = aux.siguiente
	}

}
