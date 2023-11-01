package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	objects = make(map[string]any)
	mutex   sync.Mutex
)

func main() {
	go addObject("Key 1", "Valor 1")
	go addObject("Key 2", "Valor 2")
	go addObject("Key 3", "Valor 3")
	go addObject("Key 4", "Valor 4")
	go addObject("Key 5", "Valor 5")
	go addObject("Key 6", "Valor 1")
	go addObject("Key 7", "Valor 1")
	go addObject("Key 8", "Valor 1")
	go addObject("Key 9", "Valor 1")
	go addObject("Key 10", "Valor 1")
	go addObject("Key 11", "Valor 1")
	go addObject("Key 12", "Valor 1")
	go addObject("Key 13", "Valor 1")
	go addObject("Key 14", "Valor 1")
	go printObject("Key 1")
	go addObject("Key 15", "Valor 1")
	go addObject("Key 16", "Valor 1")

	time.Sleep(time.Second * 2)

}

func addObject(key string, value any) {
	mutex.Lock()
	objects[key] = value
	mutex.Unlock()
}

func printObject(key string) {
	mutex.Lock()
	value, ok := objects[key]
	mutex.Unlock()
	if ok {
		fmt.Println(value)
	}
}
