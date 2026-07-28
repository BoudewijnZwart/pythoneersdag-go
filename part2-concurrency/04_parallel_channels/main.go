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

	var toysDone sync.WaitGroup
	channel := make(chan utils.Toy)

	// start the consumer
	go emptyChannelPutOnShelf(&shelf, channel)

	// loop over the colors
	for i, color := range colors {
		toysDone.Add(1)
		go createToy(i+1, color, &shelf, &toysDone, channel)
	}

	// wait for all the toys to be made and close the channel
	toysDone.Wait()
	close(channel)

	fmt.Printf("All done! Total time taken: %v\n", time.Since(startTime))
	fmt.Printf("Final shelf inventory: %+v\n", shelf)
}

func createToy(id int, color string, shelf *[]utils.Toy, wg *sync.WaitGroup, out chan<- utils.Toy) {
	defer wg.Done() // Lower the waitgroup counter by one

	toy := utils.FetchToyFromStorage(id)
	utils.PaintToy(&toy, color)
	utils.DryPaint(&toy)
	out <- toy
}

func emptyChannelPutOnShelf(shelf *[]utils.Toy, input <-chan utils.Toy){
	for toy := range input {
		*shelf = append(*shelf, toy)
	}
}
