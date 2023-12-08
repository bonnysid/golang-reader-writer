package main

import (
	"fmt"
)

var buffer = make(chan string)

func writer(message string) {
	buffer <- message
	fmt.Println("Писатель записал:", message)
}

func reader(n int) {
	fmt.Printf("Читатель №%d: %s\n", n, <-buffer)
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
