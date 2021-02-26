package main

import "fmt"

func main(){
	fmt.Println("La potencia de 4 ^2 es: ",potencia(4,16))

}

func potencia(x int, n int) int{
	if n == 1{
		return x
	}
	if x%2 == 0 { //  utlizamos modulo para saber si un numero es par o impar ya que un numero par tiene residuo 0
		return potencia(x,n/2) * potencia(x,n/2)
	} else{
		return x * potencia(x, (n-1) / 2)* potencia(x, (n-1)/2)
	}
	return 0
}