package particles

import (
	"container/list"
	"project-particles/config"
	"math/rand"
	"time"
)

func NewSystem() System {
	/*
	NewSystem is a function that initializes a particle system
	And adds the chosen number of particles (InitNumParticles)
	And returns it to the main function of the project, which will display it
	*/
	rand.Seed(time.Now().UnixNano())
	s := System{Content: list.New()}
	s.add_number(config.General.InitNumParticles)
	return s
}
