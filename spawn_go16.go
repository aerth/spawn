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

// +build !go1.7

// Package spawn a process like a salmon
package spawn

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/kardianos/osext"
)

// Exe is exported only for convenience
func Exe() (self string, dir string, args []string) {
	self, _ = osext.Executable()
	dir, _ = osext.ExecutableFolder()
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	return self, dir, args
}

// Spawn better than a salmon!
func Spawn() error {

	// Count increment (new process gets our env)
	i, _ := strconv.Atoi(os.Getenv(SPAWNTIME))
	i++
	os.Setenv(SPAWNTIME, strconv.Itoa(i))

	// Spawned process has new environmental variable: SPAWNED=true
	os.Setenv("SPAWNED", "true")

	me, medir, args := Exe()

	cmd := exec.Command(me, args...)
	cmd.Dir = medir
	return cmd.Start()
}

// GitPull fetches newest source code.
func GitPull(cd string, cmd ...string) (output string, ok bool) {
	if cmd == nil {
		cmd = []string{"git", "pull", "origin", "master"}
	}
	return execute(cd, cmd...)
}

// Rebuild attemps to rebuild the running go program
func Rebuild(cd string, cmd ...string) (output string, ok bool) {
	if cmd == nil {
		cmd = []string{"make"}
	}
	return execute(cd, cmd...)
}

// execute runs a program
func execute(cd string, cmd ...string) (output string, ok bool) {
	if cmd == nil {
		cmd = []string{"fortune"}
	}
	var c *exec.Cmd
	c = exec.Command(cmd[0])
	if len(cmd) > 1 {
		c = exec.Command(cmd[0], cmd[1:]...)
	}
	c.Dir = cd
	cmdlog, err := c.CombinedOutput() // block & execute
	if err != nil {
		return string(cmdlog) + "\n\n" + err.Error(), false
	}
	return string(cmdlog), true
}

// Destroy is the same as os.Exit(0) for now.
func Destroy() {
	//runtime.Gosched()
	os.Exit(0)
}
