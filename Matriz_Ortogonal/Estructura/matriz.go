package Estructura

import (
	// "fmt"
)
type Matriz struct {
	EncFila *Encabezado 
	EncColumna *Encabezado
}

func New_Matriz() *Matriz{
	return &Matriz{New_Encabezado(),New_Encabezado()}
}

func Insertar_Matriz(matriz *Matriz,fila int, columna int, valor string){
	
	//Se crea un nuevo nodo con el valor ingresado
	nuevo := New_Nodo(fila, columna, valor)

	//********************** INSERCION EN FILAS ***********************

	//primero se busca el encabezado fila
	encFila := BuscarEncabezado(fila, matriz.EncFila)
	if encFila == nil { //significa que no se encontro
		//se crea el encabezado con su enlace al nuevo Nodo
		encFila = New_NodoEzdo(fila,nuevo)
		InsertarEncabezado(encFila,matriz.EncFila) //se insertar en el encabezado de filas
	}else{

	}

	// ******************** INSERCION EN COLUMNAS **************************
	encColumna := BuscarEncabezado(columna, matriz.EncColumna)

	if encColumna == nil{
		encColumna = New_NodoEzdo(columna,nuevo)
		InsertarEncabezado(encColumna, matriz.EncColumna)
	}
}
