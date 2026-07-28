package utils

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Toy represents a toy that might be painted.
type Toy struct {
	ID int
	Color string
	Dry bool
}

// FetchToyFromStorage fetches a toy from the storage room.
func FetchToyFromStorage(id int) Toy {
	msg := fmt.Sprintf("[Toy #%d] Fetching from storage...", id)
	PrintInColor(msg,"grey")
	time.Sleep(300 * time.Millisecond)
	return Toy{ID: id}
}

// PaintToy takes a pointer to a toy and gives it the specified color.
func PaintToy(toy *Toy, color string) {
	msg := fmt.Sprintf("[Toy #%d] Painting the toy...", toy.ID) 
	DoWork(msg, color, 10_000_000)
	toy.Color = color
}

// DryPaint takes a pointer to a toy and waits for the paint to dry.
func DryPaint(toy *Toy) {
	msg := fmt.Sprintf("[Toy #%d] Waiting for %s paint to dry...", toy.ID, toy.Color)
	PrintInColor(msg, toy.Color)
	time.Sleep(800 * time.Millisecond)
	toy.Dry = true
}


// DoWork simulates doing CPU intensive work.
func DoWork(message string, colorName string, cpuCycles int) {
	PrintInColor(message, colorName)

	if cpuCycles > 0 {
		// hash is intentionally discarded; the loop exists only to burn CPU.
		hash := sha256.Sum256([]byte(message))
		for i := 0; i < cpuCycles; i++ {
			hash = sha256.Sum256(hash[:])
		}
	}
}
