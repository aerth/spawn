# spawn

### like a salmon

spawn the current process

[![Build Status](https://travis-ci.org/aerth/spawn.svg?branch=master)](https://travis-ci.org/aerth/spawn)

as seen in:

  * server daemons
  * http servers
  * chat bots
  * long running processes of all kinds

"use *anywhere*!"

  * go routines
  * main function
  * "click to run the program"
  * hook it up to your `signal.Notify()` ?

### Important notes

  * this library sets the environmental variable SPAWNED=N where N is the number of levels into this we are
  * this means if `$SPAWNED` is empty, we know its safe to spawn another process and exit (not an infinite loop).
  * the arguments, environment are carried over to the new process, so consider setting those before calling spawn.Spawn()
  * if your database (or something) is locked you may want to wait until after spawn check to lock the db
  
### Usage

```
import "github.com/aerth/spawn"
if os.Getenv(spawn.SPAWNED) == "" {
  err := spawn.Spawn()
  if err == nil {
    os.Exit(0)
  }
}
```
