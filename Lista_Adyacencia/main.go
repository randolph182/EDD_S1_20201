package main

import (
	"./Estructuras"
)

var INFINITO int = 9999

func main(){

	gra := &Estructuras.NodoDirectorio{"v1",nil,nil}
	Estructuras.NuevoVertice(&gra,"v2")
	Estructuras.NuevoVertice(&gra,"v3")
	Estructuras.NuevoVertice(&gra,"v4")
	Estructuras.NuevoVertice(&gra,"v5")
	Estructuras.NuevoVertice(&gra,"v6")



	Estructuras.Arista(&gra,"v1","v2",3)
	Estructuras.Arista(&gra,"v1","v5",8)
	Estructuras.Arista(&gra,"v1","v3",4)
	Estructuras.Arista(&gra,"v2","v5",5)
	Estructuras.Arista(&gra,"v3","v5",3)
	Estructuras.Arista(&gra,"v5","v4",7)
	Estructuras.Arista(&gra,"v5","v6",3)
	Estructuras.Arista(&gra,"v6","v4",2)

	D:= make([]*EstadoVertice,6)
	CaminoMinimos(&D,&gra,6)
}

type EstadoVertice struct {
	Ultimo *Estructuras.NodoDirectorio
	refND *Estructuras.NodoDirectorio
	Distancia int
}

func CaminoMinimos(D *[]*EstadoVertice, nodoDirectorio **Estructuras.NodoDirectorio, tamanio int){

	F := make([]*Estructuras.NodoDirectorio,tamanio)
	F[0] = *nodoDirectorio
	tmp := (*nodoDirectorio).Lista
	i := 1
	for tmp != nil{
		F[i] = nil
		nuevo := &EstadoVertice{nil,tmp.Destino,tmp.FactorPeso}
		(*D)[i] = nuevo
		i = i + 1
		tmp = tmp.Enlace
	}

 	for i := 1; i < tamanio; i++{
		v := Minimo(&F,D,tamanio)
		F[v] = (*D)[v].refND
		lstNd := (*D)[v].refND.Lista

		for lstNd != nil {
			val1 := (*D)[v].Distancia + lstNd.FactorPeso
			//for i := 1;(*D)[i] != nil && i< len(*D); i++{

				//val2 := (*D)[i].Distancia //Es para buscar si hay un camino menor que el actual
				val2 := lstNd.FactorPeso
				if val1 < val2{
					(*D)[i].Distancia = val1
					(*D)[i].Ultimo = (*D)[i].refND
				}
			//}
			//val2 := lstNd.FactorPe1

			lstNd = lstNd.Enlace
		}

		//F[v] = (*D)[v].refND
		//
		//tmp := F[v]
		//
		//for tmp != nil

		//
		//for w := 1; w < tamanio ; w++ {
		//	if F[w] == nil {
		//		if (*D)[w].Distancia  +
		//	}
		//}


	}
}

func Minimo(F *[]*Estructuras.NodoDirectorio, D *[]*EstadoVertice, tamanio int ) int {

	mx:= INFINITO
	v := 0
	for j := 1; j < tamanio; j++{
		if (*F)[j] == nil && (*D)[j] != nil && mx >= (*D)[j].Distancia {
			mx = (*D)[j].Distancia
			v = j
		}
	}
	return v
}



//func RecorrerAnchura(grafo ** Estructuras.NodoDirectorio, numeroNodos int, nodoOrigen * Estructuras.NodoDirectorio, nodosMarcados [] int){
//	//i := 0
//	//u := 0
//	//w := 0
//	queue := list.New()
//
//	//Inicializando los Nodos como no marcados
//
//	for i:= 1; i < n; i++ {
//		nodosMarcados[i] = 9999
//	}
//
//	//marcamos el Nodo de Origen
//	//nodosMarcados[0]
//	queue.PushBack(nodoOrigen)
//
//	for queue.Len() != 0 {
//		w := queue.Front()
//		fmt.Println(w.Departamento)
//
//	}
//}

//
//func Insertar(queue ** list.List,nodoOrigen *Estructuras.NodoDirectorio){
//
//}