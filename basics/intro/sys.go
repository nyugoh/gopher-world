package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func sysCalls() {
	cmd, args := "ls", "-las"

	if runtime.GOOS == "Windows" {
		fmt.Println("Can't run command on windows")
		return
	} else {
		fmt.Printf("Nice %s box, it's running Go version %s\n", runtime.GOOS, runtime.Version())
	}

	bs, err := exec.Command(cmd, args).Output()
	if err != nil {
		log.Fatal(err)
		panic("Out")
	}

	fmt.Println(string(bs[:]))
}
