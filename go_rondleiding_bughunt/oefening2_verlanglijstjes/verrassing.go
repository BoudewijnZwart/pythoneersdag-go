// Package oefening2 gaat over functies en control flow in Go.
package oefening2

// IsVerrassingsCadeau bepaalt of een binnengekomen brief een verrassingscadeau
// oplevert. Elke 5de brief die binnenkomt (dus brief nummer 5, 10, 15, ...)
// krijgt een extra verrassing bovenop het gewone cadeau.
//
// BUG: de voorwaarde hieronder klopt niet helemaal.
func IsVerrassingsCadeau(briefNummer int) bool {
	if briefNummer%5 == 1 {
		return true
	}
	return false
}
