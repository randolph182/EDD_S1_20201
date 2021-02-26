package main

import (
	"fmt"
	"./Estructuras"
	"./Grafo"
)

func main(){
	arbol := Estructuras.New_Arbol()
	Estructuras.Insertar_Arbol(arbol, 34) 
	Estructuras.Insertar_Arbol(arbol, 13)
	Estructuras.Insertar_Arbol(arbol, 21)
	Estructuras.Insertar_Arbol(arbol, 109) 
	Estructuras.Insertar_Arbol(arbol, 10293)
	v:= Estructuras.Insertar_Arbol(arbol, 5)
	fmt.Println(v)
	Estructuras.Insertar_Arbol(arbol, 100)

	// Estructuras.Insertar_Arbol(arbol, 35) 
	
	// Estructuras.Insertar_Arbol(arbol, 7)

	Grafo.Generar_Grafo(arbol)
	fmt.Println("Salio exitoso")
}