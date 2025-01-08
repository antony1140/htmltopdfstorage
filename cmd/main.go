
package main

import (
	"os"
	"fmt"
	"os/exec"
	"strings"

)

func main() {



	command := os.Args[0:]
	var file string
	var outFile string

	if len(command) > 0 {
		file = command[1]
		vals := strings.Split(file, ".")
		outFile = vals[0] + ".pdf"
	}



	cmd := exec.Command("wkhtmltopdf", file, outFile)

	err := cmd.Run()

	if err != nil {
		fmt.Println("there was an error", err)
	}






}
