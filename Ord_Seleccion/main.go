package main

import "fmt"

func main(){
	vec := [5]int{51,21,39,80,36}
	fmt.Println("vector sin ordenar:",vec)
	n := len(vec)
	for i := 0; i < n-1; i++ {
		indice_menor := i //mantenemos el valor de i como una constancia 
		for j := i+1; j < n; j++ { //recorro la sublista
			if vec[j] < vec[indice_menor]{
				indice_menor = j
			}
		}

		//compruebo si hubo alteracion en el indice menor
		if indice_menor != i {
			aux := vec[i]
			vec[i] = vec[indice_menor]
			vec[indice_menor] = aux
		}
	}

	fmt.Println("vector ordenado: ", vec)
}