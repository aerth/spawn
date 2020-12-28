// The MIT License (MIT)
//
// Copyright (c) 2016,2017  aerth <aerth@riseup.net>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package spawn a process like a salmon
package spawn

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
)

const (
	// SPAWNED_ENV ("$SPAWNED") is an environmental variable that is set for new processes created with Spawn()
	// Your application may use it to count spawn depth.
	// SPAWNED_ENV of 1 means first spawn, SPAWNTIME=2 means was spawned from a spawn,
	// SPAWNED_ENV=3 means this program was spawned from a spawned program that was spawned from another instance of the program.
	SPAWNED_ENV = "SPAWNED"
)

// GetEnviron can be used, but if using Destroy() then its easier to just os.Setenv() before running Spawn() before Destroy()
var GetEnviron func() []string = os.Environ

// Spawn better than a salmon!
func Spawn() error {
	myproc, myworkingdir, args, err := Exe()
	if err != nil {
		return fmt.Errorf("couldn't find our own process: %v", err)
	}
	// Increment SPAWNTIME count
	// (we dont care about errors, because it returns 0 if empty)
	var (
		origSpawntime     = os.Getenv(SPAWNED_ENV)
		spawntime     int = 0
	)
	if origSpawntime != "" { // spawned from a spawn
		spawntime, err = strconv.Atoi(origSpawntime)
		if err != nil {
			println("this program may be using the spawn library incorrectly")
		}
	}
	prev := spawntime
	newnum := spawntime + 1
	// Spawned process has new environmental variable: SPAWNED=1, or SPAWNED=2, or maybe even higher numbers
	os.Setenv(SPAWNED_ENV, strconv.Itoa(newnum))
	// fmt.Println("spawning:", me, medir, args)
	// path to myproc may not exist anymore, for example in go tests
	_, staterr := os.Stat(myproc)
	if staterr != nil {
		err = fmt.Errorf("trying to stat file %q ran into error: %v", myproc, err)
		if runtime.GOOS != "linux" { // TODO: ...
			return err
		}
		mypid := os.Getpid()
		tmpfile, err := ioutil.TempFile("", "spawned")
		if err != nil {
			err = fmt.Errorf("trying to stat file %q ran into error: %v", myproc, err)
			return err
		}
		procfilething, err := os.Open(filepath.Join("/", "proc", strconv.Itoa(mypid), "exe"))
		if err != nil {
			err = fmt.Errorf("trying to stat file %q ran into error: %v", myproc, err)
			return err
		}
		n, err := io.Copy(tmpfile, procfilething)
		if err != nil {
			err = fmt.Errorf("trying to stat file %q ran into error: %v", myproc, err)
			return err
		}
		if err := tmpfile.Close(); err != nil {
			return err
		}
		os.Chmod(tmpfile.Name(), 0700)
		if n == 0 {
			err = fmt.Errorf("trying to stat file %q ran into error: %v", myproc, err)
			return err
		}
		myproc = tmpfile.Name()
		_, staterr = os.Stat(myproc)
		if staterr != nil {
			staterr = fmt.Errorf("trying to stat file %q ran into error: %v", myproc, staterr)
			return staterr
		}
	}
	cmd := exec.Command(myproc, args...)
	cmd.Dir = myworkingdir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = GetEnviron() // copy the current (with incremented SPAWNED var)
	// reset current processes SPAWN variables
	os.Setenv(SPAWNED_ENV, strconv.Itoa(prev))
	return cmd.Start()
}

// Destroy is the same as os.Exit(0) for now.
func Destroy() {
	os.Exit(0)
}
