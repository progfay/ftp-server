package ftp

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/progfay/ftp-server/ftp/server"
)

func Run(w io.Writer, host string, port int) {
	log.SetOutput(w)
	log.SetFlags(0)

	url := fmt.Sprintf("%s:%d", host, port)
	ftpServer, err := server.New(url)
	if err != nil {
		log.Fatal(err)
	}

	ftpServer.Listen()

	pressed := make(chan struct{})
	go func() {
		fmt.Fprintln(w, "to shutdown server, press ENTER key...")
		os.Stdin.Read(make([]byte, 1))
		pressed <- struct{}{}
	}()

	select {
	case <-pressed:
		ftpServer.Close()

	case <-ftpServer.Cancel():
	}
}
