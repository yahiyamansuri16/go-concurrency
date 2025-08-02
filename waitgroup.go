// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1) // Add 1 goroutine to wait for
// 	go func() {
// 		defer wg.Done() // Signal completion when this goroutine finishes
// 		defer wg.Done()
// 		fmt.Println("Hello from GoRoutine")
// 	}()
// 	wg.Wait() // wait for goroutines to complete
// }
