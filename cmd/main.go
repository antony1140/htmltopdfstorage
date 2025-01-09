
package main

import (
	"os"
	"fmt"
	"os/exec"
	"strings"
	"github.com/antony1140/htmltopdfstorage/service"

)


func main() {

	log.Println("server app running")
	serverErr := http.ListenAndServe(":3000", nil)

	if serverErr != nil {
		log.Println("server failed", serverErr)
	}

http.Handle("/invoice",func invoice(w ResponseWriter, r *Request){

	log.Println(r.Body)

})
	command := os.Args[0:]
	var file string
	var outFile string

	if len(command) > 0 {
		file = command[1]
		vals := strings.Split(file, ".")
		outFile = vals[0] + ".pdf"
	}

	downloadErr := service.DownloadInvoice(file)
	if downloadErr != nil {
		panic("no such file stored in aws")
		
	}
	


	cmd := exec.Command("wkhtmltopdf", file, outFile)

	err := cmd.Run()

	if err != nil {
		fmt.Println("there was an error", err)
	}

	
	uploadErr := service.UploadInvoice(outFile)

	if uploadErr != nil {
		fmt.Println(uploadErr)
	}





}
