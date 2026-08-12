package main

import (
	"assignment2/shared/utils"
	"fmt"
	"sync"
	"time"
)

func main() {
	// maak een slice met verschillende kleuren
	kleuren := []string{"blauw", "geel", "groen"}

	// maak een plank voor het speelgoed
	var plank []utils.Speelgoed

	startTijd := time.Now()

	// maak een wait group
	var speelgoedKlaar sync.WaitGroup
	var mu sync.Mutex

	for i, kleuren := range kleuren {
		speelgoedKlaar.Add(1)
		go maakSpeelgoed(i+1, kleuren, &plank, &speelgoedKlaar, &mu)
	}

	// wait for all the toys to be made
	speelgoedKlaar.Wait()
	fmt.Printf("Klaar, tijd verstreken: %v\n", time.Since(startTijd))
	fmt.Printf("Plank: %+v\n", plank)
}

func maakSpeelgoed(id int, kleuren string, plank *[]utils.Speelgoed, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done() // lower the waitgroup counter by one
	speelgoed := utils.HaalSpeelgoedUitOpslag(id)
	utils.SchilderSpeelgoed(&speelgoed, kleuren)
	utils.DroogVerf(&speelgoed)
	mu.Lock()
	*plank = append(*plank, speelgoed)
	mu.Unlock()
}
