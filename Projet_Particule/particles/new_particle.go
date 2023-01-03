package particles 

import(
	"project-particles/config"
	"math/rand"
	"fmt"
) 


func NewParticle() (Particle){
	px, py := float64(rand.Intn(config.General.WindowSizeX)), float64(rand.Intn(config.General.WindowSizeY))
	r,g,b := rand.Float64() * (1 - 0), rand.Float64() * (1 - 0), rand.Float64() * (1 - 0)
	vx, vy, coef:= rand.Float64() * (2 - 0), rand.Float64() * (2 - 0),rand.Intn(2)
	//Determine le sens des mouvement des particules 
	fmt.Println("jdjdjd")
	if coef == 0{
		coef = -1
	}
	//Retourne une Particule
	return Particle{
		PositionX: px,
		PositionY: py,
		ScaleX:    1, ScaleY: 1,
		ColorRed: r, ColorGreen: g, ColorBlue: b,
		Opacity: 1,
		SpeedX : vx*float64(coef),
		SpeedY : vy*float64(coef),
	}
}