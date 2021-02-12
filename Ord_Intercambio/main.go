package main

import(
	"fmt"
)

func main(){

	vec := [4]int{8,4,6,2}
	n := len(vec)
	fmt.Println("el vector sin ordenar: ",vec)
	for i := 0; i <= n-2 ; i++ {
		for j := i+1; j <= n - 1; j++ {
			if vec[i] > vec[j]{ //se hace intercambio
				aux := vec[i]
				vec[i] = vec[j]
				vec[j] = aux
			}
		}
	}
	fmt.Println("vector ordenado: ",vec)

}