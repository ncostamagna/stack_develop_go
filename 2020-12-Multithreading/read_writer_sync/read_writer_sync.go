package main

import (
	"sync"
	"time"
)

var (
	money = 100
	lock  = sync.RWMutex{}
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10

		time.Sleep(1 * time.Millisecond)
		lock.Unlock()
	}
	println("Stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money -= 10

		time.Sleep(1 * time.Millisecond)
		lock.Unlock()
	}
	println("Spendy Done")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(money)
}
