package conn

import (
	"fmt"
	"log"
	"net"

	"github.com/progfay/ftp-server/ftp/os"
	"github.com/progfay/ftp-server/ftp/transfer"
)

type state struct {
	name string
	cwd  os.Cwd
}

type Conn struct {
	ctrlConn net.Conn
	dataConn net.Conn
	state    state
}

func New(ctrlConn net.Conn) Conn {
	return Conn{
		ctrlConn: ctrlConn,
		state: state{
			name: "anonymous",
			cwd:  os.NewCwd(),
		},
	}
}

var commandHandlerMap = map[string]func(*Conn, transfer.Request) transfer.Response{
	"USER": handleUSER,
	"PASS": handlePASS,
	"PORT": handlePORT,
	"LIST": handleLIST,
	"NLST": handleNLST,
	"CWD":  handleCWD,
	"PWD":  handlePWD,
	"SIZE": handleSIZE,
	"SYST": handleSYST,
	"RETR": handleRETR,
	"STOR": handleSTOR,
	"NOOP": handleNOOP,
	"QUIT": handleQUIT,
	"FEAT": handleFEAT,
	"EPSV": handleEPSV,
	"PASV": handlePASV,
}

func (conn *Conn) Handle(req transfer.Request) transfer.Response {
	log.Printf("%s >>> %s\n", conn.state.name, req.String())
	handler, ok := commandHandlerMap[req.Command]
	if !ok {
		return transfer.NewResponse(transfer.NotImplementedAtThisSite)
	}

	return handler(conn, req)
}

func (conn *Conn) Reply(res transfer.Response) {
	log.Printf("%s <<< %s\n", conn.state.name, res.String())
	fmt.Fprintf(conn.ctrlConn, "%s\n", res.Message)

	if res.HasData {
		fmt.Fprint(conn.dataConn, res.Data)
		conn.dataConn.Close()
		log.Printf("%s <<< %s\n", conn.state.name, transfer.ClosingControlConnection)
		fmt.Fprintf(conn.ctrlConn, "%s\n", transfer.ClosingDataConnection)
	}
}
