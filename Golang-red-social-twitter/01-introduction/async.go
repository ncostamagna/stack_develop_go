package main

import (
	"fmt"
	"strings"
	"time"
)


func main(){

	// con go lo ejecuto de manera async
	go nombreLentooo("Nahuel Costamagna")
	fmt.Println("Carajo")
	var x string
	fmt.Scanln(&x) //debo poner el puntero, toco aca y termina
				   // el runtime de go NO se queda esperando hasta que termine 

	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Print("LLegue hasta aqui")

	// como el await, hasta que no se cumplio no sigo con el programa
	msg := <-canal1
	fmt.Print(msg)
}

func nombreLentooo(nombre string){
	letras:=strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}

func bucle(canal chan time.Duration){
	inicio := time.Now()
	for i:=0; i<1000000000000000; i++{

	}

	final:=time.Now()
	canal <- final.Sub(inicio)
}