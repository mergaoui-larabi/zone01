package folder

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

type Server struct {
	Addr      string
	ln        net.Listener
	Client    map[net.Conn]*Client
	Clientmux sync.Mutex
}

func Newserver(addr string) *Server {
	return &Server{
		Addr:   addr,
		Client: make(map[net.Conn]*Client, 10),
	}
}

func (s *Server) ListenToConn() error {
	ln, err := net.Listen("tcp", ":"+s.Addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	s.AcceptConn()
	return nil
}

func (s *Server) AcceptConn() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(s.Client) <= 10 {
			s.Clientmux.Lock()
			s.Client[conn] = &Client{s: s, conn: conn}
			s.Clientmux.Unlock()
			go s.HandleConn(conn)
		} else {
			conn.Write([]byte("server is currently full!!!"))
			conn.Close()
		}
	}

}

func (s *Server) HandleConn(conn net.Conn) {
	defer func() {
		s.Clientmux.Lock()
		delete(s.Client, conn)
		s.Clientmux.Unlock()
		conn.Close()
	}()

	buffer := make([]byte, 2048)
	if s.Client[conn].name == "" {
		conn.Write(LINUX_LOGO)
		conn.Write([]byte("[ENTER YOUR NAME]: "))
	}
	for {
		lenght, err := conn.Read(buffer)
		if err != nil {
			return
		}
		username := string(buffer[:lenght])
		if strings.Contains(username, "\n") {
			s.Clientmux.Lock()
			s.Client[conn].name = username[:len(username)-1]
			s.Clientmux.Unlock()
			break
		}
	}

	for {
		length, err := conn.Read(buffer)
		fmt.Println(string(buffer))
		if err != nil {
			return
		}
		message := string(buffer[:length])
		if strings.Contains(message, "\n") {
			s.Broadcast(s.Client[conn].name, message)
		}
	}
}

func (s *Server) Broadcast(sender, message string) {
	s.Clientmux.Lock()
	defer s.Clientmux.Unlock()

	for conn := range s.Client {
		if s.Client[conn].name != sender {
			conn.Write([]byte("[" + sender + "] :" + message + "\n"))
		}
	}

}
