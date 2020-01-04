package spawn

import (
	"os"
	"testing"
	// "time"
)

func TestSpawn(t *testing.T) {
	if sp := os.Getenv(SPAWNTIME); sp != "3" {
		t.Log("spawning: current spawn", sp)
		if err := Spawn(); err != nil {
			t.Error(err)
			return
		}
		t.SkipNow()
	} else {
		t.Log("spawned test passed!")
	}
}
