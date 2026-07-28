package utils

import (
	"fmt"
	"time"
	"crypto/sha256"
)

type Toy struct {
	ID int
	Color string
	Dry bool
}

func FetchToyFromStorage(id int) Toy {
	fmt.Printf("[Toy #%d] Fetching from storage...\n", id)
	time.Sleep(300 * time.Millisecond)
	return Toy{ID: id}
}

func PaintToy(toy *Toy, color string) {
	msg := fmt.Sprintf("[Toy #%d] Painting the toy.", toy.ID) 
	DoWork(msg, color, 10_000_000)
	toy.Color = color
}

func DryPaint(toy *Toy) {
	msg := fmt.Sprintf("[Toy #%d] Waiting for %s paint to dry...", toy.ID, toy.Color)
	PrintInColor(msg, toy.Color)
	time.Sleep(800 * time.Millisecond)
	toy.Dry = true
}

func DoWork(message string, colorName string, cpuCycles int) {

	PrintInColor(message, colorName)

	if cpuCycles > 0 {
		hash := sha256.Sum256([]byte(message))
		for i := 0; i < cpuCycles; i++ {
			hash = sha256.Sum256(hash[:])
		}
	}
}
