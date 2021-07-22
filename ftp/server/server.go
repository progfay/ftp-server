package server

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/progfay/ftp-server/ftp/conn"
	"github.com/progfay/ftp-server/ftp/transfer"
)

type Server struct {
	listener net.Listener
	close    chan struct{}
}

func New(url string) (*Server, error) {
	listener, err := net.Listen("tcp", url)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener: listener,
		close:    make(chan struct{}),
	}, nil
}

func (s *Server) Listen() {
	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				log.Println(err)
				break
			}

			go s.handleConnection(conn)
		}
		close(s.close)
	}()
}

func (s *Server) Close() {
	close(s.close)
}

func (s *Server) Cancel() <-chan struct{} {
	return s.close
}

func (s *Server) handleConnection(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)
	fmt.Fprintln(c, transfer.ReadyForNewUser)
	conn := conn.New(c)

	for input.Scan() {
		req := transfer.ParseRequest(input.Text())

		log.Println()
		res := conn.Handle(req)
		conn.Reply(res)
		if res.Closing {
			break
		}
	}
}
