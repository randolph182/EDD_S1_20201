package Grafo
import (
	"../Estructuras"
	"fmt"
	"strconv"
	"os"
	"io/ioutil"
	"os/exec"
)

func Generar_Grafo(arbol *Estructuras.Arbol){
	estructuraDot := "digraph G{\nnode [shape=circle];\n"
	acum := ""

	if arbol.Raiz != nil{
		Recorrer_Arbol(&arbol.Raiz,&acum)
	}

	estructuraDot += acum + "\n}\n"



	path := "grafo.dot"
	//SE ESCRIBE EL ARCHIVO .DOT
	var _, err = os.Stat(path)
	if os.IsNotExist(err){
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
		fmt.Println("Se ha creado un archivo")
	}

	var file, err2 = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err2) {
		return
	 }
	 defer file.Close()

	 //SE ESCRIBE EN ARCHIVO 
	 _, err = file.WriteString(estructuraDot)
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


func Recorrer_Arbol(actual ** Estructuras.Nodo, acum *string){

	if *actual != nil{

		//SE OBTIENE INFORMACION DEL NODO ACTUAL
		*acum += "\"" + fmt.Sprint(&(*actual)) + "\"[label=\"" + strconv.Itoa((*actual).Valor) + "\"];\n"

		//VIAJAMOS A LA SUBRAMA IZQ
		if (*actual).Izquierda != nil{
			*acum += "\"" +  fmt.Sprint(&(*actual)) + "\" -> \"" + fmt.Sprint(&(*actual).Izquierda) + "\";\n"
		}
		//VIAJAMOS A LA SUBRAMA DER
		if (*actual).Derecha != nil{
			*acum += "\"" +  fmt.Sprint(&(*actual)) + "\" -> \"" + fmt.Sprint(&(*actual).Derecha) + "\";\n"
		}

		Recorrer_Arbol(&(*actual).Izquierda,acum)
		Recorrer_Arbol(&(*actual).Derecha, acum)

	}
}

func existeError(err error) bool{
	if err != nil{
		fmt.Println(err.Error())
	}
	return (err != nil)
}