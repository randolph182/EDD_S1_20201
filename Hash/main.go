package main

// ## Estructuras de datos  Seccion A
// Randolph Estuardo Muy

import (
	"./Estructuras"
)

func main(){

	tablaHash := &Estructuras.TablaHash{11,nil, 0,0}

	Estructuras.InsertarHash(tablaHash,1234,"randolph1")
	Estructuras.InsertarHash(tablaHash,1234,"randolph2")
	Estructuras.InsertarHash(tablaHash,1234,"randolph3")
	Estructuras.InsertarHash(tablaHash,1234,"randolph4")
	Estructuras.InsertarHash(tablaHash,1234,"randolph5")
	Estructuras.InsertarHash(tablaHash,1234,"randolph6")
	Estructuras.InsertarHash(tablaHash,1234,"randolph7")
	Estructuras.InsertarHash(tablaHash,1234,"randolph8")
}