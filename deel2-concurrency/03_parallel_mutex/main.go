package main

import (
	"assignment2/shared/utils"
	"fmt"
	"sync"
	"time"
)

func main() {
	// maak een slice met kleuren
	kleuren := []string{"blauw", "geel", "groen"}

	// maak een plank
	var plank []utils.Speelgoed

	startTijd := time.Now()

	// maak een wait group
	var speelgoedKlaar sync.WaitGroup

	// maak speelgoed in meerdere kleuren
	for i, kleur := range kleuren {
		speelgoedKlaar.Add(1)
		go maakSpeelgoed(i+1, kleur, &plank, &speelgoedKlaar)
	}


	// wacht to al the speelgoed klaar is
	speelgoedKlaar.Wait()
	fmt.Printf("Klaar, tijd verstreken: %v\n", time.Since(startTijd))
	fmt.Printf("Plank: %+v\n", plank)
}

func maakSpeelgoed(id int, kleur string, plank *[]utils.Speelgoed, wg *sync.WaitGroup) {
	defer wg.Done() // lower the waitgroup counter by one
	speelgoed := utils.HaalSpeelgoedUitOpslag(id)
	utils.SchilderSpeelgoed(&speelgoed, kleur)
	utils.DroogVerf(&speelgoed)
	*plank = append(*plank, speelgoed)
}
