package particles

import(
	"project-particles/config"
) 

func (p *Particle) gravity(){
	/*
	This function is a method of a particle
	It updates its verticale speed if the gravity option is true
	*/
	if config.General.Gravity{
		if p.SpeedY < 0 {
			p.SpeedY += 0.025
		}else{
			p.SpeedY += 0.045
		}
	}
}