package main

import (
	"flag"
	"io"
	"os"

	"github.com/progfay/ftp-server/ftp"
)

func main() {
	host := flag.String("host", "localhost", "string ")
	port := flag.Int("port", 8000, "Sets the port number to port.")
	silent := flag.Bool("silent", false, "Don't show logs.")
	flag.Parse()

	var w io.Writer = os.Stdout
	if *silent == true {
		w = io.Discard
	}

	ftp.Run(w, *host, *port)
}