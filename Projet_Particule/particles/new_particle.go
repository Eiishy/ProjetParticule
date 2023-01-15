package particles 

import(
	"project-particles/config"
	"math/rand"
	//"math"
) 


func newParticle() (Particle){

	//Creation of a BasicParticle 
	p := Particle{
		PositionX: float64(config.General.WindowSizeX/2), //the particle spawn in the center of the screen if the randomspawn if not activacted
		PositionY: float64(config.General.WindowSizeY/2) ,
		Rotation: 0,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1,
		SpeedX : rand.NormFloat64()-rand.NormFloat64(),
		SpeedY : rand.NormFloat64()-rand.NormFloat64(),
		Alive : true,
		LiveSpan : config.General.LiveSpan,
	}
	//Particle configuration
	p.randomColor() // choose random color for the particle if the option is set on true
	p.randomOpacity() //choose a random opacity between 30% and 100% if the option is set on true
	p.randomScale() //choose a random scale x and y between 75% and 100% if the option is set on true
	p.randomSpawn() // Spawn choose randomly if the option is set on true
	p.randomRotation() //choose a random Rotation if the option is set on true
	//Return the particle
	return p

	/* }else if config.General.Gamemod == 2 {
		//to create a circle but need RandomSpawn false and initNumParticles more than 500 and the Gamemod on 2 set in the config.json file 
		a :=rand.Float64() * 2 * math.Pi
		sx = math.Cos(a)  * 2
		sy = math.Sin(a)  * 2 
	} */
}