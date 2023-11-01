package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	fmt.Println("Nucleos", runtime.NumCPU())
	fmt.Println("Routines", runtime.NumGoroutine())
	wg.Add(2)
	go saludo1()
	go saludo2()
	fmt.Println("Routines", runtime.NumGoroutine())
	wg.Wait()

}

func saludo1() {
	fmt.Println("Saludo 1")
	wg.Done()
}

func saludo2() {
	fmt.Println("Saludo 2")
	wg.Done()
}
