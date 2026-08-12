package main

import (
	"assignment2/shared/utils"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// huur maar 1 piet in
	runtime.GOMAXPROCS(1)

	// maak een slice met kleuren
	kleuren := []string{"blauw", "geel", "groen"}

	// maak een plank aan
	var plank []utils.Speelgoed

	startTijd := time.Now()

	// maak een wait group
	var speelgoedKlaar sync.WaitGroup

	// maak een stuk speelgoed in elke kleur
	for i, kleur := range kleuren {
		speelgoedKlaar.Add(1)
		go maakSpeelgoed(i+1, kleur, &plank, &speelgoedKlaar)
	}

	// wacht tot al het speelgoed klaar is
	speelgoedKlaar.Wait()
	fmt.Printf("All done! Total time taken: %v\n", time.Since(startTijd))
	fmt.Printf("Final shelf inventory: %+v\n", plank)
}

func maakSpeelgoed(id int, kleur string, plank *[]utils.Speelgoed, wg *sync.WaitGroup) {
	defer wg.Done() // Verlaag de  wait group teller met 1
	speelgoed := utils.HaalSpeelgoedUitOpslag(id)
	utils.SchilderSpeelgoed(&speelgoed, kleur)
	utils.DroogVerf(&speelgoed)
	*plank = append(*plank, speelgoed)
}
