// Tiene que llamarse igual que el archivo
package main

import (
	"fmt"
)

var Calculo2 func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main(){
	fmt.Println(uno(2))
	fmt.Println(dos(4))
	numero, estado := dos(4)
	fmt.Println(numero, estado)
	fmt.Println(Calculo2(5,3))

	carajo:=func() int {
		return 1
	}

	fmt.Println(carajo())

	tabla:=2
	MiTabla:=Tabla(tabla) // Se ejecuta solo Tabla

	for l4:=1; l4<11; l4++ {
		fmt.Println(MiTabla()) // se ejecuta solo la funcion dentro de tabla
	}
}

func uno(numero int) int {
	return numero+2
}

func dos(numero int) (int, bool) {
	return numero+2, true
}

func uno2(numero int) (z int){
	z = numero+2
	return
}

// Parametros dinamicos, voy a recibir 1 o mas
func Calculo(numero ...int) int {
	total:=0

	// _ -> indice del elemento
	// utlilizamos _ para alojar una variable que no vamos a usar, no reserva memoria
	//
	for _, num:= range numero{
		total = total + num
	}

	return total
}

// FUNCIONES ANONIMAS

func Tabla(valor int) func() int{

	// CLOSURES -> valores que solo van a vivir en la funcion
	numero:=valor
	secuencia:=0

	return func() int {
		secuencia+=1
		return numero+secuencia
	}
}