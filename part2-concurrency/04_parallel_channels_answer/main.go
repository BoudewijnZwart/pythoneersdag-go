package main

import (
	"assignment2/shared/utils"
	"fmt"
	"sync"
	"time"
)

func main() {
	var shelf []utils.Toy
	var toysDone sync.WaitGroup
	var consumerDone sync.WaitGroup
	channel := make(chan utils.Toy)
	colors := []string{"blue", "yellow", "green"}

	startTime := time.Now()

	consumerDone.Add(1)
	go emptyChannelPutOnShelf(&shelf, channel, &consumerDone)

	for i, color := range colors {
		toysDone.Add(1)
		go createToy(i+1, color, &toysDone, channel)
	}

	// wait for all the producer, then close the channel
	toysDone.Wait()
	close(channel)

	// wait for the consumer to finish
	consumerDone.Wait()

	fmt.Printf("All done! Total time taken: %v\n", time.Since(startTime))
	fmt.Printf("Final shelf inventory: %+v\n", shelf)
}

func createToy(id int, color string, wg *sync.WaitGroup, out chan<- utils.Toy) {
	defer wg.Done() // Lower the waitgroup counter by one

	toy := utils.FetchToyFromStorage(id)
	utils.PaintToy(&toy, color)
	utils.DryPaint(&toy)
	out <- toy
}

func emptyChannelPutOnShelf(shelf *[]utils.Toy, input <-chan utils.Toy, done *sync.WaitGroup){
	defer done.Done()
	for toy := range input {
		*shelf = append(*shelf, toy)
	}
}
