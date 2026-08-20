package main

import (
	"assignment2/shared/utils"
	"fmt"
	"time"
)

func main() {
	// maak een slice met kleuren
	kleuren := []string{"blauw", "geel", "groen"}

	// maak een plank om speelgoed op te slaan
	var plank []utils.Speelgoed

	startTijd := time.Now()

	// maak een stuk speelgoed in elke kleur
	for i, kleur := range kleuren {
		maakSpeelgoed(i+1, kleur, &plank)
	}

	fmt.Printf("Klaar! Tijd verstreken: %v\n", time.Since(startTijd))
	fmt.Printf("Plank: %+v\n", plank)
}

func maakSpeelgoed(id int, kleur string, plank *[]utils.Speelgoed) {
	speelgoed := utils.HaalSpeelgoedUitOpslag(id)
	utils.SchilderSpeelgoed(&speelgoed, kleur)
	utils.DroogVerf(&speelgoed)
	*plank = append(*plank, speelgoed)
}
