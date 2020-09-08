package main

import (
	"fmt" 
)


func main(){
	paises := make(map[string] string)
	//paises_reservar := make(map[string] string, 5) // mas rapida la asignacion

	fmt.Println(paises)

	paises["Mexico"] ="D.F."
	paises["Argentina"] ="Buenos Aires"

	fmt.Println(paises["Mexico"])
	fmt.Println(paises)

	campeonato := map[string]int{
		"B":134,
		"C":100,
		"D":121}

	campeonato["E"]=29
	delete(campeonato, "B")


	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato{
		fmt.Println(equipo, puntaje )
	}

	valor, exists := campeonato["J"]
	fmt.Println(valor, exists)
	valor, exists = campeonato["E"]
	fmt.Println(valor, exists)
}

