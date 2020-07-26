package main

import "fmt"

func main(){
	var numero int
	numero = operacionesMidd(sumar)(2,3)

	fmt.Println(numero)
}

func sumar(a,b int) int{
	return a+b
}

func operacionesMidd(f func(int, int) int) func(int,int) int{
	return func(a,b int) int {
		fmt.Println("Carajo mierda")
		return f(a,b)
	}
}