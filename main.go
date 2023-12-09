package main

import (
	"fmt"
	"sync"
)

var buffer *string
var wg sync.WaitGroup

func writer(message string) {
	if buffer == nil {
		buffer = &message
		fmt.Println("Писатель записал:", message)
	}
}

func reader(n int) {
	if buffer != nil {
		fmt.Printf("Читатель №%d: %s\n", n, *buffer)
		buffer = nil
	}
}

func main() {
	numMessages := 5
	numReaders := 2
	writers := [3]string{"A", "B", "C"}

	// Запуск писателей
	go func() {
		for _, w := range writers {
			wg.Add(1)
			go func(w string) {
				defer wg.Done()
				for i := 0; i < numMessages; i++ {
					writer(fmt.Sprintf("%s%d", w, i+1))
				}
			}(w)
		}
	}()

	wg.Wait()

	// Запуск читателей
	for {
		for i := 0; i < numReaders; i++ {
			go reader(i + 1)
		}
	}
}
