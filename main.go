package main

import (
	"fmt"
	"os/exec"

	"github.com/shirou/gopsutil/v4/process"
)

func main() {
	pids, err := process.Processes()

	if err != nil {
		fmt.Println("Error fetching processes:", err)
		return
	}

	// convert to JSON. String() is also implemented
	for _, pid := range pids[1:] {
		fmt.Println("Process ID:", pid.Pid)
		p, _ := process.NewProcess(pid.Pid)
		exe, _ := pid.Exe()
		test, _ := pid.Connections()
		parent, _ := pid.Parent()
		fmt.Println(p.Pid, exe, test, parent)
	}

	cmd := exec.Command("/usr/bin/postgres", "--version")
	output, err := cmd.Output()

	fmt.Println(output)
}
