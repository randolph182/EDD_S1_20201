package main

import (
	"fmt"
)

func main(){
	
	a := [10]int{8,1,4,9,6,3,5,2,7,0}
	fmt.Println("Vector Desordenado: ", a)
	QuickSort(&a,0,len(a)-1)
	fmt.Println("Vector Ordenado: ", a)
}


func QuickSort(a *[10]int, primero int, ultimo int){
	var i,j,central int
	var pivote int

	central = (primero + ultimo) / 2  //obtengo el resultado aproximada de la posicion central
	pivote = a[central] 
	i = primero
	j = ultimo
	//Do
	for{
		for a[i] < pivote { i += 1} //hasta que encuentro un valor menor al pivote me detengo
		for a[j] > pivote { j -= 1}	// hasta que encuentre un valor mayor que el pivote me detengo

		if i <= j { // aqui es la fase de intercambio
			var tmp int
			tmp = a[i];
			a[i] = a[j]
			a[j] = tmp
			i += 1
			j -= 1
		}

		//While
		if !(i <= j){
			break
		}
	}

	if primero < j{
		QuickSort(a, primero, j)  // Apliacion de Quicksort con la sublistas Izquierda
	}

	if i < ultimo{
		QuickSort(a,i,ultimo) //  Aplicacion de Quicksort con la sublista Derecha
	}
}