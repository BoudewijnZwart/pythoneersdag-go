package oefening2

import "testing"

func TestKeurAlleZakjesGoed(t *testing.T) {
	zakjes := []Zakje{
		{Smaak: "Kaneel"},
		{Smaak: "Honing"},
		{Smaak: "Naturel"},
	}

	KeurAlleZakjesGoed(zakjes)

	for i, z := range zakjes {
		if !z.Gecontroleerd {
			t.Errorf("zakjes[%d].Gecontroleerd = false, verwacht true (smaak: %s)", i, z.Smaak)
		}
	}
}
