package main

import (
	"fmt"
	"./Estructuras"
	"./Grafo"
)


func main(){  
	// var tmp Estructuras.Balance_type
	arbol := Estructuras.New_Binary_Tree()
	Estructuras.Insert(arbol,5)
	Estructuras.Insert(arbol,6)
	Estructuras.Insert(arbol,8)
	Estructuras.Insert(arbol,9)
	Estructuras.Insert(arbol,10)
	Estructuras.Insert(arbol,11)
	// Estructuras.PreOrden(arbol)
	// Estructuras.InOrden(arbol)
	Grafo.Generar_Grafo(arbol)
	fmt.Println("hola mundo")
}