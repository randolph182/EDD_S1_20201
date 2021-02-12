package main

import "fmt"

func main(){
	vec := [17]int{1,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59}

	valor := BusquedaBinaria(vec,17,37)
	fmt.Println("El valor que me devolvio fue: ", valor)

}

func BusquedaBinaria(lista [17]int, n int, clave int) int{
	var central, bajo, alto int
	var valorCentral int

	bajo =0 
	alto = n-1
	for bajo <= alto{
		central = (bajo + alto) / 2
		valorCentral = lista[central]
		if clave == valorCentral{
			return central
		} else if clave < valorCentral { //Aca nos bamos a la sublista inferior
			alto = central - 1
		}else{ // Sino vamos a la lista Superior
			bajo = central + 1
		}
	}
	return -1  //no se encontro el valor
}