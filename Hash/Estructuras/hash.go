package Estructuras

type NodoHash struct {
	Carnet int
	Nombre string
	Llv  int
	Siguiente *NodoHash
}

type TablaHash struct{
	Tamanio int
	Primero *NodoHash
	FactorCarga int
	Id int
}



func InsertarHash(th *TablaHash, carnet int, nombre string){
	//se saca el porcentaje de carga
	factorCarga := (float32(th.Id) / float32(th.Tamanio)) * 100;

	if factorCarga < 56 && factorCarga >= 0 { //significa que posee el 55%

		llv := carnet % th.Tamanio


		insertar(th, carnet,nombre,llv)

	} else{ //aplicacion del Rehashing

		//aumentamos el tamanio a un numero impar

		th.Tamanio += 2
		//Corregir , ver video de clase
		//InsertarHash(th, carnet,nombre) //recursivo



	}
}

func insertar(th *TablaHash, carnet int, nombre string, llv int){
	nuevo := &NodoHash{carnet,nombre,llv,nil}

	if th.Primero == nil {
		th.Primero = nuevo
		th.Primero.Siguiente = nil
		th.Id++
		return
	}

	//APLICANDO LA COLICION, para saber si la llave ya existe

	if buscarLlave(th, nuevo.Llv){
		//se usa la funcion s(llv,i) - ( llv % 7 + 1 ) * i
		i := 0
		pos := buscarPos(th,nuevo,i)
		insertar(th,carnet,nombre,pos)

	} else{
		tmp := th.Primero

		if nuevo.Llv < tmp.Llv{ //agreagando a inicio
			nuevo.Siguiente = tmp
			th.Primero = nuevo
			th.Id += 1
		} else { //agreagando al medio

			for tmp.Siguiente != nil{
				tmp2 := tmp.Siguiente
				if nuevo.Llv < tmp2.Llv {
					tmp.Siguiente = nuevo
					nuevo.Siguiente = tmp2
					th.Id += 1
					break // saliendo del for
				}
				tmp = tmp.Siguiente
			}

			if (tmp.Siguiente == nil){ //INSERTANDO AL FINAL
				tmp.Siguiente = nuevo
				th.Id += 1
			}
		}
	}

}

func buscarPos(th *TablaHash, actual *NodoHash, i int) int {
	//usando la funcion  s(llv,i) = (llv % 7 +1) *i
	//a := **i
	pos := (actual.Llv % 7 + 1)* i

	if buscarLlave(th, pos){
		i += 1
		return buscarPos(th, actual,i)
	}

	return pos

}

func buscarLlave(th *TablaHash, llv int) bool{
	tmp := th.Primero
	for tmp != nil {
		if llv == tmp.Llv{
			return true
		}
		tmp = tmp.Siguiente
	}
	return false
}
