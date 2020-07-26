package main

import (
	"fmt"
	"os"
	"log"

)


func main(){

	//f, _ := os.Open("./archivo.txt")

	// esta instruccion no se ejecuta secuencialmente, se ejecuta recien cuando salga
	// cuando haga el Exit, se ejecuta siempre a lo ultimo
	// defer f.Close()

	// panic("Se encontro un error") // es como el throw

	ejemploPanic()
	fmt.Println("Pasa aca")
	os.Exit(1)
}

func ejemploPanic(){
	defer func(){
		reco := recover() // resultado del ultimo panic
		if reco != nil{

			// Me hace un printf y un exit al mismo tiempo
			log.Fatalf("ourrio un error que genero Panic\n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro un error")
	}
}