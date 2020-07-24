// Tiene que llamarse igual que el archivo
package main

import (
	"fmt" // mostrar por consila, ingresar por teclado, etc..
)

var numero int
var texto string
var status bool

/*
float32 float64
uint -> enteros sin signos
*/
func main(){
	var numero2 int

	// Toma directamente el tipo de dato int con el valor que le meto
	// tiene que ser una nueva variable
	numero3 := 4 
	// asigno un valor asi -> numero3 = 15

	// Asignar varios valores creando nuevas variables
	numero4, numero5, numero6 := 4, 4,4
	numero6, numero7, numero8, numero9 := 4, 12, 5, "Texto"

	texto = fmt.Sprintf("%d", 23)
	
	fmt.Println(numero, texto, status, numero2, numero3)
	fmt.Println(numero4, numero5, numero6)
	fmt.Println(numero6, numero7, numero8, numero9)
}