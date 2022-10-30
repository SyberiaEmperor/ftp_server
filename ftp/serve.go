package ftp

import (
	"bufio"
	"log"
	"strings"
)

func Serve(c *Conn) {
	c.respond(status220)

	s := bufio.NewScanner(c.conn)

	for s.Scan() {
		input := strings.Fields(s.Text())

		if len(input) == 0 {
			continue
		}

		command, args := input[0], input[1:]
		log.Printf("<< %s %v", command, args)

		switch command {
		case "USER":
			c.user(args)
		case "PORT":
			c.port(args)
		case "CWD":
			c.cd(args)
		case "NLST":
			c.list(args)
		case "RETR":
		c.get(args)
		case "QUIT":
			c.respond(status221)
			return
		case "TYPE":
			c.setDataType(args)
		default:
			c.respond(status502)
		}
	}

	if s.Err() != nil {
		log.Print(s.Err())
	}
}



