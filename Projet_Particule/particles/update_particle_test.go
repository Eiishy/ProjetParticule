package particles 

import (
	"testing"
	"project-particles/config"
)

func init() {
	//This function is executed at the start of the tests and allows to recover the data of the config.json
	config.Get("../config.json")
}

func Test_Particle_In_Movement(t *testing.T) {
	//This function tests if the correct number of particles is added to the system at each call of the update function
	p := newParticle()
	px_init,py_init := p.PositionX,p.PositionY
	p.update()
	if px_init == p.PositionX && py_init == p.PositionY {
		t.Error("La particule ne se deplace pas")
		t.Log("\nBefore --> ",px_init, " : ", py_init,"\nAfter --> ",p.PositionX," : ",p.PositionY)
	}
}