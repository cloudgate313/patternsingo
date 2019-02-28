// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	goRoutines := 20
// 	wg.Add(goRoutines)
// 	for i := 0; i < goRoutines; i++ {
// 		go func(goRoutineID int) {
// 			fmt.Println("Go routine executed ", goRoutineID)
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }
