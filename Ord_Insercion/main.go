package main

import "fmt"

func main(){
	vec := [5]int{1,4,6,5,2}

	var j int
	var aux int

	n := len(vec)
	fmt.Println("vectores de entrada: ",vec)
	
	for i := 1; i < n; i++ {
		j = i
		aux = vec[i]
		
		for j > 0 && aux < vec[j - 1]{
			vec[j] = vec[j-1]
			j--
		}
		vec[j] = aux
	}
	
	fmt.Println("vector ordenado: ",vec)

}