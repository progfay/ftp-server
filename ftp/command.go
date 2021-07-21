package ftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/progfay/ftp-server/ftp/transfer"
)

func handleUSER(conn *ftpConn, req transfer.Request) transfer.Response {
	conn.state.name = req.Message
	return transfer.NewResponse(transfer.NeedPassword)
}

func handlePASS(conn *ftpConn, req transfer.Request) transfer.Response {
	return transfer.NewResponse(transfer.UserLoggedIn)
}

func handlePORT(conn *ftpConn, req transfer.Request) transfer.Response {
	hostPort := strings.Split(req.Message, ",")
	if len(hostPort) != 6 {
		return transfer.NewResponse(transfer.WrongArguments)
	}
	large, err := strconv.Atoi(hostPort[4])
	if err != nil {
		log.Println(err)
		return transfer.NewResponse(transfer.WrongArguments)
	}
	small, err := strconv.Atoi(hostPort[5])
	if err != nil {
		log.Println(err)
		return transfer.NewResponse(transfer.WrongArguments)
	}
	host := strings.Join(hostPort[:4], ".")
	port := int64(large*256 + small)
	address := fmt.Sprintf("%s:%d", host, port)
	dataConn, err := net.Dial("tcp", address)
	if err != nil {
		return transfer.NewResponse(transfer.CantOpenConnection)
	}
	conn.dataConn = dataConn
	return transfer.NewResponse(transfer.Ok)
}

func handleLIST(conn *ftpConn, req transfer.Request) transfer.Response {
	files, err := conn.state.cwd.Ls(req.Message)
	if err != nil {
		log.Println(err)
		return transfer.NewResponse(transfer.WrongArguments)
	}

	lines := []string{}
	for _, file := range files {
		fileType := "file"
		if file.IsDir() {
			fileType = "dir "
		}
		lines = append(lines, fmt.Sprintf("%s\t%d\t%s\t%s", file.Mode(), file.Size(), fileType, file.Name()))
	}
	res := transfer.NewResponse(transfer.FileStatusOk)
	res.SetData(strings.Join(lines, "\r\n"))
	return res
}

func handleNLST(conn *ftpConn, req transfer.Request) transfer.Response {
	files, err := conn.state.cwd.Ls(req.Message)
	if err != nil {
		log.Println(err)
		return transfer.NewResponse(transfer.WrongArguments)
	}

	lines := []string{}
	for _, file := range files {
		lines = append(lines, fmt.Sprintf("%s", file.Name()))
	}
	res := transfer.NewResponse(transfer.FileStatusOk)
	res.SetData(strings.Join(lines, "\r\n"))
	return res
}

func handleCWD(conn *ftpConn, req transfer.Request) transfer.Response {
	err := conn.state.cwd.Cd(req.Message)
	if err != nil {
		return transfer.NewResponse(transfer.WrongArguments)
	}
	return transfer.NewResponse(transfer.FileActionOk)
}

func handlePWD(conn *ftpConn, req transfer.Request) transfer.Response {
	return transfer.NewResponse(fmt.Sprintf(transfer.Created, conn.state.cwd.Pwd()))
}

func handleSIZE(conn *ftpConn, req transfer.Request) transfer.Response {
	return transfer.NewResponse(transfer.NotImplementedAtThisSite)
}

func handleSYST(conn *ftpConn, req transfer.Request) transfer.Response {
	return transfer.NewResponse(fmt.Sprintf(transfer.NameSystemType, "UNIX"))
}

func handleRETR(conn *ftpConn, req transfer.Request) transfer.Response {
	data, err := conn.state.cwd.Get(req.Message)
	if err != nil {
		log.Println(err)
		return transfer.NewResponse(transfer.WrongArguments)
	}
	res := transfer.NewResponse(transfer.FileStatusOk)
	res.SetData(string(data))
	return res
}

func handleSTOR(conn *ftpConn, req transfer.Request) transfer.Response {
	conn.Reply(transfer.NewResponse(transfer.FileStatusOk))
	data, err := ioutil.ReadAll(conn.dataConn)
	if err != nil {
		log.Println(err)
		return transfer.NewResponse(transfer.WrongArguments)
	}
	err = conn.state.cwd.Put(req.Message, data)
	if err != nil {
		log.Println(err)
		return transfer.NewResponse(transfer.WrongArguments)
	}
	return transfer.NewResponse(transfer.ClosingDataConnection)
}

func handleNOOP(conn *ftpConn, req transfer.Request) transfer.Response {
	return transfer.NewResponse(transfer.Ok)
}

func handleQUIT(conn *ftpConn, req transfer.Request) transfer.Response {
	res := transfer.NewResponse(transfer.ClosingControlConnection)
	res.Close()
	return res
}

func handleFEAT(conn *ftpConn, req transfer.Request) transfer.Response {
	return transfer.NewResponse(transfer.NotImplemented)
}

func handleEPSV(conn *ftpConn, req transfer.Request) transfer.Response {
	return transfer.NewResponse(transfer.NotImplemented)
}

func handlePASV(conn *ftpConn, req transfer.Request) transfer.Response {
	return transfer.NewResponse(transfer.NotImplemented)
}
