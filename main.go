package main

import (
	"fmt"
	"sync"
)

const (
	numMessages   = 10 // Число сообщений
	messageLength = 5  // Длина сообщения
)

var (
	buffer       = make([]byte, messageLength)
	readersMutex = sync.Mutex{}
	writersMutex = sync.Mutex{}
	readersCond  = sync.NewCond(&readersMutex)
	writersCond  = sync.NewCond(&writersMutex)
)

func writer(id int) {
	for i := 0; i < numMessages; i++ {
		writersMutex.Lock()
		for buffer[0] != 0 {
			writersCond.Wait()
		}
		// Запись сообщения в буфер
		for j := 0; j < messageLength; j++ {
			buffer[j] = byte('A' + id*100 + i) // Просто пример данных для записи
		}
		readersCond.Broadcast() // Сообщаем читателям, что буфер готов
		writersMutex.Unlock()
	}
}

func reader(id int) {
	for i := 0; i < numMessages; i++ {
		readersMutex.Lock()
		for buffer[0] == 0 {
			readersCond.Wait()
		}
		// Чтение сообщения из буфера
		message := string(buffer[:messageLength])
		fmt.Printf("Reader %d read: %s\n", id, message)
		// Сбрасываем буфер
		for j := 0; j < messageLength; j++ {
			buffer[j] = 0
		}
		writersCond.Broadcast() // Сообщаем писателям, что буфер освободился
		readersMutex.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5) // 3 писателя + 2 читателя

	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			writer(id)
		}(i)
	}

	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			reader(id)
		}(i)
	}

	wg.Wait()
}
