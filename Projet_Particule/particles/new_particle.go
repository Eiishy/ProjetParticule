package particles 

import(
	"project-particles/config"
	"math/rand"
) 


func newParticle() (Particle){
	//Creation of variables for the color, position and speed of a particle
	var px,py float64

	//Spawn configuration
	if config.General.RandomSpawn{
		px = float64(rand.Intn(config.General.WindowSizeX - 10))
		py = float64(rand.Intn(config.General.WindowSizeY -10))
	}else{
		px = float64(config.General.SpawnX)
		py = float64(config.General.SpawnY)
	}
	//creates and returns a particle
	return Particle{
		PositionX: px,
		PositionY: py,
		Rotation: 0,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1,
	}
}