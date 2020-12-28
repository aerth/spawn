package spawn

import (
	"os"
	"testing"
	"time"
	// "time"
)

func TestSpawn(t *testing.T) {
	if sp := os.Getenv(SPAWNED_ENV); sp != "10" {
		t.Logf("spawning: current spawn #%s", sp)
		<-time.After(time.Second)
		if err := Spawn(); err != nil {
			t.Error(err)
			return
		}
		t.SkipNow()
	} else {
		t.Log("spawned test passed!")
	}
}
