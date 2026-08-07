package main

import (
	"assignment2/shared/utils"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// set the number of "pete's"  to 1 (no parallelism)
	runtime.GOMAXPROCS(1)

	// create a slice of colors
	colors := []string{"blue", "yellow", "green"}

	// create the toy shelf
	var shelf []utils.Toy

	startTime := time.Now()

	// create a waitgroup
	var toysDone sync.WaitGroup

	// loop over the colors
	for i, color := range colors {
		toysDone.Add(1)
		go createToy(i+1, color, &shelf, &toysDone)
	}

	// wait for all the toys to be made
	toysDone.Wait()
	fmt.Printf("All done! Total time taken: %v\n", time.Since(startTime))
	fmt.Printf("Final shelf inventory: %+v\n", shelf)
}

func createToy(id int, color string, shelf *[]utils.Toy, wg *sync.WaitGroup) {
	defer wg.Done() // Lower the waitgroup counter by one
	toy := utils.FetchToyFromStorage(id)
	utils.PaintToy(&toy, color)
	utils.DryPaint(&toy)
	*shelf = append(*shelf, toy)
}
