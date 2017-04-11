# spawn

### like a salmon

spawn the current process

[![Build Status](https://travis-ci.org/aerth/spawn.svg?branch=master)](https://travis-ci.org/aerth/spawn)

as seen in:

  * server daemons
  * http servers
  * irc bots

"use *anywhere*!"

  * go routines
  * main function


```
import "github.com/aerth/spawn"
if os.Getenv("SPAWN") == "" {
  spawn.Spawn()
  spawn.Destroy()
}
```
