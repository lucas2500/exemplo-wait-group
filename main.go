package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// O WaitGroup permite que seja especificada a quantidade de goroutines a serem
	// executadas, fazendo assim com que a goroutine principal aguardade a finalização das demais
	// antes de ser encerrada

	var wt sync.WaitGroup

	for i := 0; i < 5; i++ {

		wt.Add(1)

		go func() {
			defer wt.Done()
			PrintText(2 * time.Second)
		}()
	}

	wt.Wait()
}

func PrintText(t time.Duration) {

	fmt.Println("Starting execution...")

	time.Sleep(t)

	fmt.Println("Ending execution...")
}
