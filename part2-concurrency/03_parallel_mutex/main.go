package main

import (
	"assignment2/shared/utils"
	"fmt"
	"sync"
	"time"
)

func main() {
	// create a slice of colors
	colors := []string{"blue", "yellow", "green"}

	// create the toy shelf
	var shelf []utils.Toy

	startTime := time.Now()

	// create a waitgroup
	var toysDone sync.WaitGroup
	var mu sync.Mutex

	// loop over the colors
	for i, color := range colors {
		toysDone.Add(1)
		go createToy(i+1, color, &shelf, &toysDone, &mu)
	}

	// wait for all the toys to be made
	toysDone.Wait()
	fmt.Printf("All done! Total time taken: %v\n", time.Since(startTime))
	fmt.Printf("Final shelf inventory: %+v\n", shelf)
}

func createToy(id int, color string, shelf *[]utils.Toy, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done() // lower the waitgroup counter by one
	toy := utils.FetchToyFromStorage(id)
	utils.PaintToy(&toy, color)
	utils.DryPaint(&toy)
	mu.Lock()
	*shelf = append(*shelf, toy)
	mu.Unlock()
}
