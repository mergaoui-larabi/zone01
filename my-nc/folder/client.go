package folder

import "net"

type Client struct {
	name string
	s    *Server
	conn net.Conn
}
