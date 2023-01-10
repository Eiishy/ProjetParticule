package particles

import (
	"testing"
	"project-particles/config"
)

func init() {
	/*This function is executed at the start of the tests and allows to recover the data of the config.json*/
	config.Get("../config.json")
}

func Test_Position_Random(t *testing.T) {
	/*This function tests if the coordinates of spawn x and y particles are chosen randomly*/
	if config.General.RandomSpawn {
		p1 := newParticle()
		p2 := newParticle()
		if p1.PositionX == p2.PositionX && p1.PositionY == p2.PositionY{
			t.Error("Particles are not randomly generated")
			t.Log("\n 1 -->",p1.PositionX,":",p1.PositionY,"\n 2 -->",p2.PositionX,":",p2.PositionY)
	}
	}
}

func Test_Spawn_OK(t *testing.T){
	/*This function tests if a particle creates spawn to the chosen coordinates*/
	if !config.General.RandomSpawn{
		p1 := newParticle()
		if p1.PositionX != float64(config.General.WindowSizeX/2) && p1.PositionY != float64(config.General.WindowSizeY/2){
			t.Error("Particles do not spawn at the chosen coordinates")
			t.Log("\n 1 -->",p1.PositionX,":",p1.PositionY)
		}
	}
}

func Test_Particle_Spawn_In_Screen(t *testing.T) {
	/*This function tests if a particle created is located in the screen*/
	p := newParticle()
	if p.PositionX <0 || int(p.PositionX) > config.General.WindowSizeX || p.PositionY < 0 || int(p.PositionY) > config.General.WindowSizeY{
		t.Error("The particles do not spawn in the screen")
		t.Log("\n 1 -->",p.PositionX,":",p.PositionY)
	}
}
