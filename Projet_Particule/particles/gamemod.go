package particles

import (
	"project-particles/config"
)

func circle(){
	if config.General.Gamemod == "circle"{
		config.General.LiveSpanActivator = false
		config.General.RandomRotation = false
		config.General.Gravity = false
		config.General.InitNumParticles = 0
		config.General.NumMaxParticles = 10000
		config.General.SpawnRate = 0
		config.General.OnClick = true
	}
}