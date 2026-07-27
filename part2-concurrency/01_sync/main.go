package main

import (
	"fmt"
	"time"
	"assignment2/shared/utils"
)


type Toy struct {
	ID int
	Color string
	Dry bool
}


func main() {
	colors := []string{"blue", "yellow", "green"}
	var shelf []Toy

	startTime := time.Now()

	for i, color := range colors {
		toy := createToy(i+1, color)
		shelf = append(shelf, toy)
	}
	fmt.Printf("All done! Total time taken: %v\n", time.Since(startTime))
	fmt.Printf("Final shelf inventory: %+v\n", shelf)
}

func createToy(id int, color string) Toy {
	toy := fetchToyFromStorage(id)
	paintToy(&toy, color)
	dryPaint(&toy)
	return toy
}

func fetchToyFromStorage(id int) Toy {
	fmt.Printf("[Toy #%d] Fetching from storage...\n", id)
	time.Sleep(300 * time.Millisecond)
	return Toy{ID: id}
}

func paintToy(toy *Toy, color string) {
	msg := fmt.Sprintf("[Toy #%d] Painting the toy.", toy.ID) 
	utils.DoWork(msg, color, 5_000_000)
	toy.Color = color
}

func dryPaint(toy *Toy) {
	msg := fmt.Sprintf("[Toy #%d] Waiting for %s paint to dry...", toy.ID, toy.Color)
	utils.PrintInColor(msg, toy.Color)
	time.Sleep(800 * time.Millisecond)
	toy.Dry = true
}
