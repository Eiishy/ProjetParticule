package particles


func (p *Particle) update(){
	/*
	This function is a method of a particle
	It updates its position according to its speed
	*/
 	p.PositionX += p.SpeedX
	p.PositionY += p.SpeedY 
	p.gravity() //The vertical speed of the particle changes with gravity. 
}