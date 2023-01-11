package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"os"
)

func (s *System) Update() {

	//update each particle of the system 
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p:=e.Value.(*Particle)
		p.update()
	}

	//Generate more particles
    s.spawnUpdate()
	
	//to close the window with the escape key
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape){
		os.Exit(0)
	}
}
