package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"os"
)

func (s *System) Update() {

	//update each alive particle of the system and delete the others
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p, ok :=e.Value.(*Particle)
		if !ok {
			continue
		}
		if p.Alive {
			p.update()
		}else {
			go s.Content.Remove(e)
		}
	}

	//Generate more particles
    s.spawnUpdate()
	
	//to close the window with the escape key
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape){
		os.Exit(0)
	}
}
