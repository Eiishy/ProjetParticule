package particles 

import(
	"project-particles/config"
	"math/rand"
	"math"
) 


func newParticle() (Particle){
	//Creation of variables position and speed of a particle
	var px,py,sx,sy float64

	//Spawn configuration
	//Spawn choose randomly 
	if config.General.RandomSpawn{
		px = float64(rand.Intn(config.General.WindowSizeX - 10))
		py = float64(rand.Intn(config.General.WindowSizeY -10))
	//Spawn in the center of the screen 
	}else{
		px = float64(config.General.WindowSizeX/2)
		py = float64(config.General.WindowSizeY/2)
	}

	//Speed configuration
	if config.General.Gamemod == 1 {
		//choose a random number between -1 and 1 for the particle speed 
		sx = rand.NormFloat64()-rand.NormFloat64()
		sy = rand.NormFloat64()-rand.NormFloat64()  
	}else if config.General.Gamemod == 2 {
		//to create a circle but need RandomSpawn false and initNumParticles more than 500 and the Gamemod on 2 set in the config.json file 
		a :=rand.Float64() * 2 * math.Pi
		sx = math.Cos(a)  * 2
		sy = math.Sin(a)  * 2 
	}
	

	//creates and returns a particle
	return Particle{
		PositionX: px,
		PositionY: py,
		Rotation: 0,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1,
		SpeedX : sx,
		SpeedY : sy,
		Alive : true,
		LiveSpan : config.General.LiveSpan,
	}
}