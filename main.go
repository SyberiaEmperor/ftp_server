package main

import (
	"log"
	"net"
	"path/filepath"
	"projects/ftp_server/ftp"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:20")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	path, err := filepath.Abs("work_dir")
	if err != nil {
		return
	}

	ftp.Serve(ftp.NewConn(c, path))
}

