/*
This program emulates the following :
echo "some input" | cat
*/
package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println(" Trying to Send commands to create the Keys")
	cmd := exec.Command("cat")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Command output: %q\n", out.String())
}
