package Estructuras

type Nodo struct{
	Valor int
	Izquierda *Nodo
	Derecha *Nodo
}

func New_Nodo(valor int) *Nodo{
	return &Nodo{valor,nil,nil}
}