package main

import (
	"demo"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	fmt.Println("summing:", os.Args[1:])
	sum := demo.SumStrings(os.Args[1:]...)
	fmt.Println("total:", sum)

	fmt.Println("sleeping")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(2 * time.Second)
		wg.Done()
	}()

	fmt.Println("waiting for sleep to finish")
	wg.Wait()
	fmt.Println("done")
}
