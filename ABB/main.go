package main

import (
	"fmt"
	"./Estructuras"
	"./Grafo"
)

func main(){
	arbol := Estructuras.New_Arbol()
	Estructuras.Insertar_Arbol(arbol, 10) 
	Estructuras.Insertar_Arbol(arbol, 11)
	Estructuras.Insertar_Arbol(arbol, 8)
	Estructuras.Insertar_Arbol(arbol, 6) 
	Estructuras.Insertar_Arbol(arbol, 5)
	Estructuras.Insertar_Arbol(arbol, 9)
	// v:= Estructuras.Insertar_Arbol(arbol, 5)
	// fmt.Println(v)
	// Estructuras.Insertar_Arbol(arbol, 100)

	// Estructuras.Insertar_Arbol(arbol, 35) 
	
	// Estructuras.Insertar_Arbol(arbol, 7)

	Grafo.Generar_Grafo(arbol)
	fmt.Println("Salio exitoso")
}