package utils

import (
	"fmt"
)

// ANSI escape codes voor kleuren in de terminal
const (
	Reset  = "\033[0m"
	Blauw   = "\033[34m"
	Geel = "\033[33m"
	Groen  = "\033[32m"
	Grijs	= "\033[37m"
)

// Helper functie om kleuren te mappen naar ANSI codes
func krijgKleurCode(color string) string {
	switch color {
	case "blauw":
		return Blauw
	case "geel":
		return Geel
	case "groen":
		return Groen
	case "grijs":
		return Grijs
	default:
		return Reset
	}
}

func PrintInKleur(tekst string, kleurNaam string) {
	kleurCode := krijgKleurCode(kleurNaam)
	fmt.Printf("%s %s %s\n", kleurCode, tekst, Reset)
}

