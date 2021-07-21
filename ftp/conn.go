package ftp

import (
	"fmt"
	"net"

	"github.com/progfay/ftp-server/ftp/transfer"
)

type state struct {
	name string
	cwd  Cwd
}

type ftpConn struct {
	ctrlConn net.Conn
	dataConn net.Conn
	state    state
}

func newftpConn(ctrlConn net.Conn) ftpConn {
	return ftpConn{
		ctrlConn: ctrlConn,
		state: state{
			name: "anonymous",
			cwd:  newCwd(),
		},
	}
}

var commandHanderMap = map[string]func(*ftpConn, transfer.Request) transfer.Response{
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

func (conn *ftpConn) handle(req transfer.Request) transfer.Response {
	fmt.Printf("%s >>> %s\n", conn.state.name, req.String())
	handler, ok := commandHanderMap[req.Command]
	if !ok {
		return transfer.NewResponse(transfer.NotImplementedAtThisSite)
	}

	return handler(conn, req)
}

func (conn *ftpConn) Reply(res transfer.Response) {
	fmt.Printf("%s <<< %s\n", conn.state.name, res.String())
	fmt.Fprintf(conn.ctrlConn, "%s\n", res.Message)

	if res.HasData {
		fmt.Fprint(conn.dataConn, res.Data)
		conn.dataConn.Close()
		fmt.Printf("%s <<< %s\n", conn.state.name, transfer.ClosingControlConnection)
		fmt.Fprintf(conn.ctrlConn, "%s\n", transfer.ClosingDataConnection)
	}
}
