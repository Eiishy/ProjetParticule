package particles

import "project-particles/config"


func (p *Particle) update(){
	//deplacement horizontale
	if p.SpeedX > 0{
		if p.PositionX < float64(config.General.WindowSizeX)-10*p.ScaleX{
			p.PositionX += p.SpeedX
		}else{
			p.SpeedX = -p.SpeedX 
		}
	}
	if p.PositionX > 0+ 2*p.ScaleX {
		p.PositionX += p.SpeedX
	}else{
		p.SpeedX = -p.SpeedX
	}
	//deplacement verticale
	if p.SpeedY > 0{
		if p.PositionY < float64(config.General.WindowSizeY)-10*p.ScaleY{
			p.PositionY += p.SpeedY
		}else{
			p.SpeedY = -p.SpeedY
		}
	}
	if p.PositionY > 0+ 2*p.ScaleY {
		p.PositionY += p.SpeedY
	}else{
		p.SpeedY = -p.SpeedY
	}
}