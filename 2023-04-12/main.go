package main

func main() {
	// var wg sync.WaitGroup
	// foo := make(chan int)
	// bar := make(chan int)
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	select {
	// 	case foo <- <-bar:
	// 	default:
	// 		println("default")
	// 	}
	// }()
	// wg.Wait()

	foo := make(chan int, 2)
	foo <- 1
	bar := make(chan int)
	foo <- <-bar
}
