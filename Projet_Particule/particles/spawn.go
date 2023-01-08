package particles 

import(
	"math"
	"project-particles/config"
)

func (s *System) spawnUpdate() {
	/*
	spawnUpdate is a system method that updates the number of particles in real time based on the spawnrate
	For this it adds the spawrate to the spawnvalue which is an attribute of the system
	If the spawnvalue is greater than one then we create particles and we keep only the decimal part
	This method is called in the update function which is run 60 times per second
	*/
	s.SpawnValue += config.General.SpawnRate
	if s.SpawnValue >= 1{
		s.add_number(int(s.SpawnValue))							
		_, s.SpawnValue = math.Modf(s.SpawnValue)
	}
}