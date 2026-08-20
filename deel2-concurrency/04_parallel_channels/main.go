package main

import (
	"assignment2/shared/utils"
	"fmt"
	"sync"
	"time"
)

func main() {
	var plank []utils.Speelgoed
	var speelgoedKlaar sync.WaitGroup
	var consumentKlaar sync.WaitGroup
	channel := make(chan utils.Speelgoed)
	kleuren := []string{"blauw", "geel", "groen"}

	startTijd := time.Now()

	consumentKlaar.Add(1)
	go verplaatsVanChannelNaarPlank(&plank, channel, &consumentKlaar)

	for i, kleur := range kleuren {
		speelgoedKlaar.Add(1)
		go maakSpeelgoed(i+1, kleur, &speelgoedKlaar, channel)
	}

	// wacht op de producenten
	speelgoedKlaar.Wait()
	close(channel)

	// wacht op de consument
	consumentKlaar.Wait()

	fmt.Printf("Klaar, tijd verstreken: %v\n", time.Since(startTijd))
	fmt.Printf("Plank: %+v\n", plank)
}

func maakSpeelgoed(id int, kleur string, wg *sync.WaitGroup, uit chan<- utils.Speelgoed) {
	defer wg.Done() // Lower the waitgroup counter by one

	speelgoed := utils.HaalSpeelgoedUitOpslag(id)
	utils.SchilderSpeelgoed(&speelgoed, kleur)
	utils.DroogVerf(&speelgoed)
	uit <- speelgoed
}

func verplaatsVanChannelNaarPlank(plank *[]utils.Speelgoed, input <-chan utils.Speelgoed, klaar *sync.WaitGroup){
	defer klaar.Done()
	for speelgoed := range input {
		*plank = append(*plank, speelgoed)
	}
}
