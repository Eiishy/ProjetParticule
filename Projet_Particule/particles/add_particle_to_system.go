package particles


//this function allows to add a Particle to the system
func (s *System) add(){
		p1 := newParticle()
		s.Content.PushFront(&p1)
		s.AliveParticleCounter += 1 	//Update the counter ou particles 
	}

//This function allows you to add a finite number of particles to the system
func (s *System) add_number(number int){
	for i:=0;i<number;i++{
		s.add()
	}
}
