package main

import (
	"fmt" 
)
 var estado bool
func main(){

	estado = true

	if estado == true {
		fmt.Println(estado)
	}

	// Puedo asignar una variable en medio de un if
	if estado=false; estado == true {
		fmt.Println(estado)
	}

	switch numero:=5; numero {
	case 1:
		//algo
	case 2:
		//algo
	default:
		//algo
		
	}

	i := 1
	for i < 10{
		fmt.Println(i)
		i++
	}

	for i2:=1;i2 < 10;i2++{
		fmt.Println(i2)
	}

	for {
		// podemos usar Printf tambien
		fmt.Println(i)
		i++
		if i > 20 {
			break
		}
	}


	var i3 = 0
	RUTINA: //Indicamos un inicio del codigo, una etiqueta
		i3++
		fmt.Printf("Entra rutina")
	
	fmt.Printf("Ejecuto esto")
	if i3 < 5{
		goto RUTINA
	}
	
	fmt.Println("Hola Mundo")
}