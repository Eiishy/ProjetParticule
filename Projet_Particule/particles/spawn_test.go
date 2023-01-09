package particles 

import (
	"testing"
	"project-particles/config"
)

func init() {
	//this function is executed at the start of the tests and allows to recover the data of the config.json
	config.Get("../config.json")
}

func Test_Spawn_Rate_Ok(t *testing.T) {
	//This function tests if the correct number of particles is added to the system at each call of the update function
	s := NewSystem()
	before:= s.Content.Len()
	s.spawnUpdate()
	after :=  before+int(config.General.SpawnRate)
	if after != s.Content.Len() {
		t.Error("Le spawn regulier de particules ne marche pas")
		t.Log("\nBefore : ",before,"\n + ",config.General.SpawnRate,"\nAfter : ",after)
	}

}