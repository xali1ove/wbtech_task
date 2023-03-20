package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func worker(id int, dataCh chan int) {
	// Бесконечный цикл чтения из канала и вывода в stdout
	for {
		data := <-dataCh
		fmt.Printf("Worker %d: %d\n", id, data)
	}
}

func main() {
	// Создаем канал для передачи данных между главным потоком и воркерами
	dataCh := make(chan int)

	// Запрос ввода количества воркеров
	var numWorkers int
	fmt.Print("Enter number of workers: ")
	_, err := fmt.Scanln(&numWorkers)
	if err != nil {
		fmt.Printf("Error reading number of workers: %v\n", err)
		os.Exit(1)
	}

	// Создаем воркеров и запускаем их
	for i := 1; i <= numWorkers; i++ {
		go worker(i, dataCh)
	}

	// Записываем произвольные данные в канал
	for i := 1; ; i++ {
		dataCh <- i
	}

	// Ждем сигнала о завершении программы (Ctrl+C)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	// Завершаем работу воркеров
	close(dataCh)
	fmt.Println("Exiting program")
}
