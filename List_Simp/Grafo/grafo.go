package Grafo

import (
	"os"
	"fmt"
	"strconv"
	"../Estructura"
	"io/ioutil"
	"os/exec"
)

func GenerarGrafo(lista *Estructura.Lista){
	path := "grafo.dot"
	acum := "digraph G{\n rankdir = LR;\nnode [shape=box]; \n compound=true; \n"
	nodo := ""
	enlace := ""
	//Hacemos un recorrido desde el primer nodo
	aux := lista.Primero

	for aux.Siguiente != nil { 
		nodo += strconv.Itoa(aux.Persona.Id) + "[label=\"Nombre: "+ aux.Persona.Nombre + "\n"
		nodo += "Edad : "+ strconv.Itoa(aux.Persona.Edad) + "\"];\n"
		enlace += strconv.Itoa(aux.Persona.Id) + "->" + strconv.Itoa(aux.Siguiente.Persona.Id) + ";\n";
		aux = aux.Siguiente
	}
	//Estamos en el ultimo nodo
	nodo += strconv.Itoa(aux.Persona.Id) + "[label=\"Nombre: "+ aux.Persona.Nombre + "\n"
	nodo += "Edad : "+ strconv.Itoa(aux.Persona.Edad) + "\"];\n"

	//    /sacando enlaces desde la primera posicion a la ultima
	// NodoAvion *tmp2 = ultimo;
   //      while(tmp2->anterior !=NULL)
   //      {
   //          enlaces += tmp2->idNodo+ "->"+tmp2->anterior->idNodo+ ";\n";
   //          tmp2 = tmp2->anterior;
   //      }
	
	acum += nodo + enlace + "\n}\n"


	//creacion del archivo
	var _, err = os.Stat(path)
	if os.IsNotExist(err){
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
		fmt.Println("Se ha creado un archivo")
	}

	//AHORA ESCRIBIEMOS EN EL ARCHIVO

	var file, err2 = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err2) {
		return
	 }
	 defer file.Close()

	 //SE ESCRIBE EN ARCHIVO 
	 _, err = file.WriteString(acum)
	 if existeError(err) {
		return
	 }

	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}

	fmt.Println("Archivo actualizado existosamente.")

	//PARTE EN DONDE GENERO EL GRAFO 
	path2, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path2, "-Tpng","grafo.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("grafo.png",cmd, os.FileMode(mode))
}

func existeError(err error) bool{
	if err != nil{
		fmt.Println(err.Error())
	}
	return (err != nil)
}