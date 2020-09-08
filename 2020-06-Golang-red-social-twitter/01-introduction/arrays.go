package main

import (
	"fmt" 
)

var tabla [10]int // de 10 items, no lo puedo ampliar mas
var matriz [5][10]int 
var matriz_slice []int // cuando no le asigno logitud

func main(){
	tabla[0]=1
	tabla[5]=15
	fmt.Println(tabla)

	tabla2 := [10]int{4,3,1,2,3,4,2,1}
	fmt.Println(tabla2, len(tabla2)) 

	variante3()

}

func variante2(){
	elementos := [5]int{1,2,3,4,5}
	porcion := elementos[3:] //de la posicion 3 hasta el ultimo
	fmt.Println(elementos, porcion) 
}

func variante3(){
	elementos := make([]int,5,20) //tipo, largo elemento, capacidad (en memmoria)
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

	nums := make([]int,0,0)

	for i:=0; i < 100; i++ {

		// lo ideal es usar solo append cuando me quede sin espacio
		// gastara su tiempo en volver a incrementar memoria
		nums = append(nums,i)
	}

	fmt.Println(nums)
	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
	//capacidad 128 porque se guia por la secuencia binaria
}