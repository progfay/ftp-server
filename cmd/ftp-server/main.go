package main

import (
	"flag"

	"github.com/progfay/ftp-server/ftp"
)

func main() {
	flag.Parse()
	args := flag.Args()
	ftp.Run(args)
}