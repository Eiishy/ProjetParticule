package particles

func (s *System) add(p Particle){
	s.Content.PushFront(&p)
}


