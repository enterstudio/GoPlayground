/**
* Experiment FAILED !

Could not set up the OpenSSL command to be run on Windows
This was due to the problem with `Spaces in the Path`

E.g.:
C:\Program Files\Git\mingw64\bin\openssl.exe
The space in he `Program Files` causes the command to fail

**/
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const keygenCmd = "openssl"

//const keygenArgs = ` req -x509 -newkey rsa:2048 -keyout rsa_private.pem -nodes -out rsa_cert.pem -subj "//CN=unused"`
var keygenArgs = []string{
	"/C",
	"openssl",
	"req",
	"-x509",
	"-newkey rsa:2048",
	"-keyout rsa_private.pem",
	"-nodes",
	"-out rsa_cert.pem",
	`-subj "//CN=unused"`,
}

func main() {
	fmt.Println(" Trying to Send commands to create the Keys")
	path, err := exec.LookPath(keygenCmd)
	if err != nil {
		log.Fatal("Open SSL not found in Path")
		os.Exit(1)
	}
	fmt.Printf("openssl is available at %s\n", path)
	/*
		args := keygenArgs
		cmd := exec.Command(path, args...)
		fmt.Printf(" Executing Command \n%s\n", strings.Join(cmd.Args, " "))
		out, err := cmd.CombinedOutput()
		fmt.Printf(" Command Output: \n%s\n", string(out))
		if err != nil {
			log.Fatal("Error 1: ", err)
		}
		//fmt.Printf(" Command Output: \n%s\n", string(out))
	*/
	cmd2 := exec.Command("pwd")
	fmt.Printf(" Executing Command \n%s\n", strings.Join(cmd2.Args, " "))
	out2, _ := cmd2.CombinedOutput()
	fmt.Printf(" Command Output: \n%s\n", string(out2))
	path = strings.Replace(path, "\\", "/", 1)
	path = strings.Replace(path, " ", "\\ ", 1)
	// openssl req -x509 -nodes -keyout rsa_private.pem -out rsa_cert.pem -newkey rsa:2048 -subj "//CN=unused"
	//cmd3 := exec.Command("cmd", "/C", "\""+path+"\"", "req -x509 -nodes -keyout rsa_private.pem -out rsa_cert.pem -newkey rsa:2048 -subj \"/CN=unused\"")
	//cmd3 := exec.Command("cmd", "/C", string(`"`+path+`"`), "--help")
	cmd3 := exec.Command("cmd", keygenArgs...)
	fmt.Printf(" Executing Command \n%s\n", strings.Join(cmd3.Args, " "))
	out3, _ := cmd3.CombinedOutput()
	fmt.Printf(" Command Output: \n%s\n", string(out3))
}
