package main

import (
	"fmt"
	"sync"
	"time"
)

var buffer string
var mutex = &sync.Mutex{}

func writer(message string) {
	mutex.Lock()
	buffer = message
	mutex.Unlock()
}

func reader() {
	mutex.Lock()
	fmt.Println("Read:", buffer)
	mutex.Unlock()
}

func main() {
	numWriters := 3
	numReaders := 2

	for i := 0; i < numWriters; i++ {
		go func(id int) {
			for j := 0; j < 3; j++ {
				writer(fmt.Sprintf("Writer %d: Message %d", id, j))
				time.Sleep(time.Millisecond * 100) // Задержка для имитации работы
			}
		}(i)
	}

	for i := 0; i < numReaders; i++ {
		go func(id int) {
			for {
				reader()
				time.Sleep(time.Millisecond * 200) // Задержка для имитации работы
			}
		}(i)
	}

	time.Sleep(time.Second * 2)
}
