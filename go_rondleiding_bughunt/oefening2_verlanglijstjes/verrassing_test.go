package oefening2

import "testing"

func TestIsVerrassingsCadeau(t *testing.T) {
	gevallen := map[int]bool{
		1:  false,
		4:  false,
		5:  true,
		6:  false,
		10: true,
		11: false,
		15: true,
	}

	for briefNummer, verwacht := range gevallen {
		resultaat := IsVerrassingsCadeau(briefNummer)
		if resultaat != verwacht {
			t.Errorf("IsVerrassingsCadeau(%d) = %v, verwacht %v", briefNummer, resultaat, verwacht)
		}
	}
}
