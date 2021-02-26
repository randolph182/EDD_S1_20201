package Estructuras

type Arbol struct {
	Raiz *Nodo
}

func New_Arbol() *Arbol{
	return &Arbol{nil}
}


func Insertar_Arbol(arbol *Arbol, valor int) bool{
	return Insertar_Nodo(&arbol.Raiz, valor)
}

func Insertar_Nodo(nodo **Nodo, valor int ) bool {

		if *nodo == nil {

			*nodo = New_Nodo(valor)

			return true

		} else if  valor < (*nodo).Valor { // insertar en la rama izquierda

			return Insertar_Nodo( &(*nodo).Izquierda, valor)

		} else if valor > (*nodo).Valor   { // insertar en la rama derecha

			return Insertar_Nodo( &(*nodo).Derecha, valor)

		}
		return false
}