package utils

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Speelgoed struct {
	ID int
	Kleur string
	Droog bool
}

// HaalSpeelgoedUitOpslag haalt een stuk ongeschilderd speelgoed op.
func HaalSpeelgoedUitOpslag(id int) Speelgoed {
	msg := fmt.Sprintf("[Speelgoed #%d] Speelgoed ophalen uit de opslag...", id)
	PrintInKleur(msg,"grijs")
	time.Sleep(300 * time.Millisecond)
	return Speelgoed{ID: id}
}

// SchilderSpeelgoed accepteert een pointer naar een stuk speelgoed en verft dit.
func SchilderSpeelgoed(speelgoed *Speelgoed, kleur string) {
	msg := fmt.Sprintf("[Speelgoed #%d] Speelgoed verven...", speelgoed.ID) 
	DoeWerk(msg, kleur, 6_000_000)
	speelgoed.Kleur = kleur
}

// DroogVerf accepteert een pointer naar een stuk speelgoed en wacht tot deze droog is. 
func DroogVerf(speelgoed *Speelgoed) {
	msg := fmt.Sprintf("[Speelgoed #%d] Wachten to de verf droog is van %s ...", speelgoed.ID, speelgoed.Kleur)
	PrintInKleur(msg, speelgoed.Kleur)
	time.Sleep(800 * time.Millisecond)
	speelgoed.Droog = true
}


// DoeWerk simuleert het doen van CPU intensief werk.
func DoeWerk(tekst string, kleurNaam string, cpuCycles int) {
	PrintInKleur(tekst, kleurNaam)

	if cpuCycles > 0 {
		hash := sha256.Sum256([]byte(tekst))
		for i := 0; i < cpuCycles; i++ {
			hash = sha256.Sum256(hash[:])
		}
	}
}
