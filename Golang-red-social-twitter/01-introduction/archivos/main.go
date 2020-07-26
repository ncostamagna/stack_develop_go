package main

import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
)

func main(){
	graboArchivo()
	graboArchivo2()

	leoArchivo()
	leoArchivo2()
}

func graboArchivo(){
	archivo, err := os.Create("./archivo.txt")

	// nil es null
	if err != nil{
		fmt.Println("Hubo un error")
		return
	}
	
	fmt.Fprintln(archivo, "Esta es una linea nueva")
	fmt.Fprintln(archivo, "Otra")

	archivo.Close()
}

func graboArchivo2(){
	// adicionar texto
	fileName := "./archivo.txt"
	if Append(fileName, "\nEsta es otra linea") == false {
		fmt.Println("Hubo un error")
	}
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)

	// nil es null
	if err != nil{
		fmt.Println("Hubo un error")
		return false
	}
	
	_, err = arch.WriteString(texto)

	if err != nil{
		fmt.Println("Hubo un error")
		return false
	}
	
	return true
}

func leoArchivo(){
	// Lee todo el archivo de un intento
	archivo, err := ioutil.ReadFile("./archivo.txt")
	// nil es null
	if err != nil{
		fmt.Println("Hubo un error")
	}else{
		fmt.Println(string(archivo))
	}
}

func leoArchivo2(){
	// leo linea por linea
	archivo, err := os.Open("./archivo.txt")
	// nil es null

	if err != nil{
		fmt.Println("Hubo un error")
	}else{
		scanner := bufio.NewScanner(archivo)

		for scanner.Scan(){
			registro:=scanner.Text()
			fmt.Println("Linea > " + registro)
		}
	}
	archivo.Close()
}