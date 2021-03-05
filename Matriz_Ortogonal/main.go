package main

import (
	"fmt"
	"./Estructura"
	"./Grafo"
)
// ptr := fmt.Sprint(&matriz) //para obtener la direccion de un nodo 

func main(){
	fmt.Println("hola mundo ")
	matriz := Estructura.New_Matriz()
	Estructura.Insertar_Matriz(matriz,1,1,"hola")
	Estructura.Insertar_Matriz(matriz,2,2,"hola")
	Estructura.Insertar_Matriz(matriz,3,3,"hola")
	Estructura.Insertar_Matriz(matriz,4,4,"hola")
	Grafo.GenerarMatriz(matriz)
}