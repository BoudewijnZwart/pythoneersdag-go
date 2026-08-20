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

	// wacht tot al het speelgoed klaar is
	speelgoedKlaar.Wait()
	fmt.Printf("Klaar, tijd verstreken: %v\n", time.Since(startTijd))
	fmt.Printf("Plank: %+v\n", plank)
}

func maakSpeelgoed(id int, kleuren string, plank *[]utils.Speelgoed, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done() // verlaag de wait group teller met 1
	speelgoed := utils.HaalSpeelgoedUitOpslag(id)
	utils.SchilderSpeelgoed(&speelgoed, kleuren)
	utils.DroogVerf(&speelgoed)
	mu.Lock()
	*plank = append(*plank, speelgoed)
	mu.Unlock()
}
