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
	defer wg.Done() // verlaag de teller van de wait group met 1
	speelgoed := utils.HaalSpeelgoedUitOpslag(id)
	utils.SchilderSpeelgoed(&speelgoed, kleur)
	utils.DroogVerf(&speelgoed)
	*plank = append(*plank, speelgoed)
}
