package main

import "fmt" 




func main(){
	var x,y *int
	entero:=5
	x = &entero
	y = &entero

	printDir(x,y,&entero)
	printVal(*x, *y, entero)

	entero = 10

	printDir(x,y,&entero)
	printVal(*x, *y, entero)

	*x = 20

	printDir(x,y,&entero)
	printVal(*x, *y, entero)

}

func printDir(x *int, y *int, entero *int){
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(entero)
}

func printVal(x int, y int, entero int){
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(entero)
}
