package main

import (
	"assignment2/shared/utils"
	"fmt"
	"time"
)

func main() {
	// create a slice of colors
	colors := []string{"blue", "yellow", "green"}

	// create the toy shelf
	var shelf []utils.Toy

	startTime := time.Now()

	// loop over the colors
	for i, color := range colors {
		createToy(i+1, color, &shelf)
	}

	fmt.Printf("All done! Total time taken: %v\n", time.Since(startTime))
	fmt.Printf("Final shelf inventory: %+v\n", shelf)
}

func createToy(id int, color string, shelf *[]utils.Toy) {
	toy := utils.FetchToyFromStorage(id)
	utils.PaintToy(&toy, color)
	utils.DryPaint(&toy)
	*shelf = append(*shelf, toy)
}
