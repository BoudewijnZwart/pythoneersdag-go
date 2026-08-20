package main

import (
	"assignment2/shared/utils"
	"fmt"
	"runtime"
	"time"
)

func main() {
	// huur maar 1 piet in (geen parallelisme)
	runtime.GOMAXPROCS(1)

	// maak een slice met kleuren
	kleuren := []string{"blauw", "geel", "groen"}

	// maak een plank voor de opslag
	var plank []utils.Speelgoed

	startTime := time.Now()

	// maak een stuk speelgoed in elke kleur
	for i, color := range kleuren {
		maakSpeelgoed(i+1, color, &plank)
	}

	fmt.Printf("Klaar, tijd vertreken: %v\n", time.Since(startTime))
	fmt.Printf("Plank: %+v\n", plank)
}

func maakSpeelgoed(id int, kleur string, plank *[]utils.Speelgoed) {
	speelgoed := utils.HaalSpeelgoedUitOpslag(id)
	utils.SchilderSpeelgoed(&speelgoed, kleur)
	utils.DroogVerf(&speelgoed)
	*plank = append(*plank, speelgoed)
}
