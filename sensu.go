package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	holdPid := pid()
	fmt.Printf("%s", holdPid)

}
func pid() string {
	processName := os.Args[1]

	out, err := exec.Command("pidof", processName).Output()
	if err != nil || len(out) == 0 {
		os.Exit(2)
	}

	return fmt.Sprintf("%s, out")
}
