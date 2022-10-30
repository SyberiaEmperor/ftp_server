package ftp

import "net"

type Conn struct {
	conn		net.Conn

	dataPort 	*dataPort
	dataType	dataType
	rootDir		string
	workDir		string
}

func NewConn(conn net.Conn, rootDir string) *Conn {
	return &Conn{
		conn: conn,
		rootDir: rootDir,
		workDir: "/",
	}
}