package Estructuras

type NodoDirectorio struct{
	Departamento string
	Lista *NodoAdyacente
	Enlace *NodoDirectorio
}

type NodoAdyacente struct{
	Destino *NodoDirectorio
	FactorPeso int
	Enlace *NodoAdyacente
}

func NuevoVertice( grafo ** NodoDirectorio, departamento string){
	nuevoDep := &NodoDirectorio{departamento,nil,nil}

	tmp := *grafo

	if tmp.Enlace == nil{
		tmp.Enlace = nuevoDep
	}else{
		for tmp != nil {
			if tmp.Enlace == nil {
				tmp.Enlace = nuevoDep
				break
			}
			tmp = tmp.Enlace
		}
	}



	//*grafo = nuevoDep

	//nuevoDep = *grafo
}


func Arista(grafo ** NodoDirectorio, oringen string, destino string, peso int){
	v1 := getDirecccion(grafo,oringen)
	v2 := getDirecccion(grafo,destino)

	if v1 != nil && v2 != nil {
		ady := &NodoAdyacente{v2,peso,nil}

		if v1.Lista == nil{
			v1.Lista = ady
		}else{
			tmp := v1.Lista
			for tmp != nil {
				if tmp.Enlace == nil {
					tmp.Enlace = ady
					break
				}
				tmp = tmp.Enlace
			}
		}
		//ady.Enlace = v1.Lista
		//v1.Lista = ady
	}
}

func getDirecccion(grafo ** NodoDirectorio, departamento string) *NodoDirectorio {
	tmp := *grafo
	for tmp != nil {
		if tmp.Departamento == departamento{
			return tmp
		}
		tmp = tmp.Enlace
	}
	return nil
}