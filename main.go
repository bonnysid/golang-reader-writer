package main

import (
	"fmt"
	"sync"
)

var (
	buffer string
	mutex  sync.Mutex
)

func writer(message string) {
	mutex.Lock()
	buffer = message
	fmt.Println("Писатель записал:", buffer)
	mutex.Unlock()
}

func reader(n int) {
	mutex.Lock()
	fmt.Printf("Читатель №%d: %s\n", n, buffer)
	mutex.Unlock()
}

func main() {
	numMessages := 5
	numReaders := 2
	writers := [3]string{"A", "B", "C"}

	// Запуск писателей
	go func() {
		for _, w := range writers {
			go func(w string) {
				for i := 0; i < numMessages; i++ {
					writer(fmt.Sprintf("%s%d", w, i+1))
				}
			}(w)
		}
	}()

	// Запуск читателей
	for {
		for i := 0; i < numReaders; i++ {
			go reader(i + 1)
		}
	}
}
