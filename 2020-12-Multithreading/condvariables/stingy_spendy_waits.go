package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	money          = 100
	lock           = sync.Mutex{}
	moneyDeposited = sync.NewCond(&lock) // Conditional variable
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 3
		fmt.Println("Stingy sees balance of ", money)
		moneyDeposited.Signal() // cuando incrementamos mandamos una señal para
		// al wait para que siga ejecutando
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		for money-20 < 0 {
			fmt.Println("Espera el otro thread")
			// La operacion sera bloqueada si el valor da negativo
			// Espera a ver el valor que le asigna el otro thread,
			// si es negativo sigue el bucle, sino sale
			moneyDeposited.Wait()
			fmt.Println("Termino el otro thread")
		}
		fmt.Println("Decrementa!")
		money -= 20
		fmt.Println("Spendy sees balance of ", money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Spendy Done")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(money)
}
