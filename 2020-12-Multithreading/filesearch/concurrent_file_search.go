package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	matches    []string
	waitgroup  = sync.WaitGroup{}
	waitgroup2 = sync.WaitGroup{}
	lock       = sync.Mutex{}
)

func fileSearch(root string, filename string) {
	fmt.Println("Searching in", root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			// bloqueamos variable para no pisarla
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			// Agregamos un nuevo thread hijo y sumamos uno
			waitgroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}

	// finalizamos thread
	waitgroup.Done()
}

func main() {

	// indicamos que vamos a tener un thread
	waitgroup.Add(1)
	go fileSearch("/media/ncostamagna/Costamagna/gopath/src/github.com/ncostamagna/stack_develop_go/2020-12-Multithreading", "README.md")

	// esperamos hasta que terminen en Done TODOS LOS THREADS
	waitgroup.Wait()
	for _, file := range matches {
		fmt.Println("Matched", file)
	}
	fmt.Println()
	othesTest()
}

func othesTest() {
	waitgroup2.Add(3)

	go correrProceso(1000, "Proceso 1")
	go correrProceso(1700, "Proceso 2")
	go correrProceso(500, "Proceso 3")

	waitgroup2.Wait()
	fmt.Println("Finalizo")
}

func correrProceso(t int, process string) {
	fmt.Printf("%s - Begin\n", process)
	time.Sleep(time.Duration(t) * time.Millisecond)
	fmt.Printf("%s - End\n", process)
	waitgroup2.Done()
}
