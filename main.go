package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/Users/amado/sandbox/play/blink-diff/bin/blink-diff", "--output result.png", "/Users/amado/sandbox/play/react/react-test-sample/old.png /Users/amado/sandbox/play/react/react-test-sample/out.png")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}
