package ftp

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/progfay/ftp-server/ftp/server"
)

func Run(w io.Writer, host string, port int) {
	log.SetOutput(w)
	log.SetFlags(0)

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	url := fmt.Sprintf("%s:%d", host, port)
	ftpServer, err := server.New(ctx, url)
	if err != nil {
		log.Fatal(err)
	}

	ftpServer.Listen()
	<-ctx.Done()
}
