// Package oefening5 gaat over interfaces in Go.
package oefening5

import "fmt"

// Bezorgmethode is de interface die elke manier van bezorgen moet implementeren.
type Bezorgmethode interface {
	Bezorg(pakket string) string
}

// Stoomboot is de traditionele manier waarop Sinterklaas zelf aankomt.
type Stoomboot struct{}

// BUG: Stoomboot voldoet niet aan de Bezorgmethode interface. Zoek uit waarom
// niet (het compileert momenteel niet als je Stoomboot probeert te gebruiken
// als Bezorgmethode).
func (s Stoomboot) bezorg(pakket string) string {
	return fmt.Sprintf("%s wordt per stoomboot bezorgd", pakket)
}

// Paard is hoe Amerigo de pakjes langs de daken brengt. Deze klopt al helemaal.
type Paard struct{}

func (p Paard) Bezorg(pakket string) string {
	return fmt.Sprintf("%s wordt per paard over de daken bezorgd", pakket)
}
