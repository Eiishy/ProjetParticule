package particles

import (
	"project-particles/config"
)

func (p *Particle) update(){
	/*
	This function is a method of a particle
	It updates its position according to its speed and its livespan
	If the particle is Alive 
	And kill the particle if it go out of the screen
	*/
	if p.Alive{
		p.LiveSpan--
		p.PositionX += p.SpeedX
		p.PositionY += p.SpeedY 

 		//The vertical speed of the particle changes with gravity. 
		p.gravity()

		//kill the particles if their LiveSpan is finish
		if config.General.LiveSpanActivator && p.LiveSpan < 0 {
			p.Alive = false
		}
	}
	
	//Verify if the particle go out of the screen 
	if p.PositionX > float64(config.General.WindowSizeX) +10 || p.PositionX < -10  ||p.PositionY > float64(config.General.WindowSizeY) +10  || p.PositionY < -10 {
		p.Alive = false //Set the Alive parameter to false
	}
}