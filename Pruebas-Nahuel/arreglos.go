package main

import "fmt"

func main()  {
	arreglo := [10]int{4,3,1,2,3,4,2,1,3,4}

	fmt.Println(arreglo[5:])
	fmt.Println(arreglo[2:6])
	fmt.Println(arreglo[:5])
	fmt.Println(arreglo[:])
}
