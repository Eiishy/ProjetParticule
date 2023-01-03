package particles

import "testing"

func TestNew1(t *testing.T) {
	p:= NewParticle()
	if p.PositionX!=1 {
		t.Error("Le mot vide n'est pas avant le mot vide dans l'ordre alphabétique")
	}
	
}