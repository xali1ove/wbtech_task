package main

import (
	"context"
	"fmt"
	"time"
)

// 1. Использование флага для завершения цикла
func stopUsingFlag() {
	stop := false
	go func() {
		for {
			if stop {
				fmt.Println("Stop using flag")
				return
			}
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(5 * time.Second)
	stop = true
}

// 2. Использование контекста
func stopUsingContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stop using context")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(5 * time.Second)
	cancel()
}

// 3. Использование канала
func stopUsingChannel() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Stop using channel")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	stop <- true
}

// 4. Использование таймера
func stopUsingTimer() {
	timer := time.NewTimer(5 * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("Stop using timer")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	// Если нам нужно остановить горутину раньше, мы можем остановить таймер
	// timer.Stop()
}

func main() {
	stopUsingFlag()
	stopUsingContext()
	stopUsingChannel()
	stopUsingTimer()
}
