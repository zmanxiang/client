package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	for {
		script := "curl http://test-server:8090"
		var out bytes.Buffer
		cmd := exec.Command("/bin/sh")
		cmd.Stdin = strings.NewReader(script)
		cmd.Stdout = &out
		cmd.Stderr = &out
		errmsgs := ""
		if err := cmd.Run(); err != nil {
			errmsgs += fmt.Sprintf("failed to run: %q, output: %q, error: %s\n", script, out.String(), err)
			log.Fatal(errmsgs)
		}

		fmt.Println(out.String())
		time.Sleep(10 * time.Second)
	}

}
