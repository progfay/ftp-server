package ftp

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/progfay/ftp-server/ftp/transfer"
)

type ftpServer struct {
	listener net.Listener
	close    chan struct{}
}

func New(url string) (*ftpServer, error) {
	listener, err := net.Listen("tcp", url)
	if err != nil {
		return nil, err
	}

	return &ftpServer{
		listener: listener,
		close:    make(chan struct{}),
	}, nil
}

func (s *ftpServer) Listen() {
	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				log.Println(err)
				break
			}

			go handleConnection(conn)
		}
		close(s.close)
	}()
}

func (s *ftpServer) Close() {
	close(s.close)
}

func (s *ftpServer) Cancel() <-chan struct{} {
	return s.close
}

func handleConnection(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)
	fmt.Fprintln(c, transfer.ReadyForNewUser)
	conn := newftpConn(c)

	for input.Scan() {
		req := transfer.ParseRequest(input.Text())

		fmt.Println()
		res := conn.handle(req)
		conn.Reply(res)
		if res.Closing {
			break
		}
	}
}
