// Spawn a process like a fork
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
	i, _ := strconv.Atoi(os.Getenv("SPAWNTIME"))
	i++
	os.Setenv("SPAWNTIME", strconv.Itoa(i))

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
