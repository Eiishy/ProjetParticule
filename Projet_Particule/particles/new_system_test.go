package particles

import (
	"project-particles/config"
	"testing"
)

func Test_New_System(t *testing.T) {
	//This function tests if when a system is created it has the right number of particles
	s:= NewSystem()
	if s.Content.Len() != config.General.InitNumParticles{
		t.Error("Le systeme n'as pas le nombre de particules initiale lors de sa création")
	}
} 