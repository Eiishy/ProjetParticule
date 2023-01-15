package particles

import(
	"project-particles/config"
	"math/rand"
) 

//5.1
func (p *Particle) gravity(){
	/*
	This function is a method of a particle
	It updates its verticale speed if the gravity option is true
	*/
	if config.General.Gravity{
		if p.SpeedY < 0 {
			p.SpeedY += 0.035
		}else{
			p.SpeedY += 0.045
		}
	}
}

//5.3
func (p *Particle) updateLiveSpan(){
	//Update the particle LiveSpan 
	p.LiveSpan--
	//kill the particles if their LiveSpan is finish and the LiveSpanActivator is set on true
	if config.General.LiveSpanActivator && p.LiveSpan < 0 {
		p.Alive = false
	}
}
//5.4
//s to change the specs of a particle at its creation
func (p *Particle) randomColor(){
	if config.General.RandomColor{
		p.ColorRed,p.ColorBlue,p.ColorGreen  = rand.Float64(),rand.Float64(),rand.Float64()
	}
}
func (p *Particle) randomScale(){
	if config.General.RandomScale{
		p.ScaleX = 0.75 + rand.Float64() * (0.25)
		p.ScaleY = 0.75 + rand.Float64() * (0.25)
	}
}
func (p *Particle) randomOpacity(){
	if config.General.RandomOpacity{
		p.Opacity = 0.3 + rand.Float64() * 0.7
	}
}
func (p *Particle) randomSpawn(){
	if config.General.RandomSpawn{
		p.PositionX = float64(rand.Intn(config.General.WindowSizeX - 10))
		p.PositionY = float64(rand.Intn(config.General.WindowSizeY -10))
	}
}
func (p *Particle) randomRotation(){
	//Rotation configuration 
	if  config.General.RandomRotation{
		p.Rotation = rand.Float64() 
	}
}
func (p *Particle) SpeedMode(){}

//5.4
//Methods to uptade the specs of a particle during its livespan
func (p *Particle) updatePosition(){
	//Update the position of a particle by adding the speed 
	p.PositionX += p.SpeedX
	p.PositionY += p.SpeedY
}
func (p *Particle) updateRotation(){
	p.Rotation += 0.08
}
func (p *Particle) updateOpacity(){
	if config.General.OpacityManagementMode == 1 {
		//increase Opacity
		if p.Opacity < 1 {
			p.Opacity +=0.001
		}
	}
	if config.General.OpacityManagementMode == 2 {
		//decrease Opacity
		if p.Opacity > 0 {
			p.Opacity -=0.003
		}else{
			p.Alive = false
		}
	}
}

func (p *Particle) updateColor(){
	
}




