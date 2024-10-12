package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"reds-internals/config"
)

func readCommand(c net.Conn) (string, error) {
	// takes the read connection and fires the system call
	var buf []byte = make([]byte, 512)
	// blocking call :: waiting for the data to be read from the connection
	n, err := c.Read(buf[:])
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func respond(c net.Conn, cmd string) error {
	// takes the connection and the command and writes the command back to the client
	_, err := c.Write([]byte(fmt.Sprintf("You said %s", cmd)))
	return err
}

func RunSyncTcpServer() {
	// This is the main entry point for the sync_tcp server
	// It will listen on the configured host and port
	// and handle incoming connections

	log.Println("Starting a TCP Server on", config.Host, config.Port)

	var con_clients int = 0 // Number of connected clients

	//server listening
	lsnr, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))

	if err != nil {
		panic(err)
	}

	for {
		// blocking call :: waiting for a connection, until my client makes connection my code will wait at here
		conn, err := lsnr.Accept()
		if err != nil {
			panic(err)
		}

		//increment the concurrent clients
		con_clients++
		log.Println("client connected with address", conn.RemoteAddr(), "Number of connected clients", con_clients)

		for {
			// over the socket, continuous read the data and print it out
			cmd, err := readCommand(conn)
			if err != nil {
				con_clients--
				log.Println("client disconnected", conn.RemoteAddr())
				if err == io.EOF {
					break
				}
				log.Println("error reading command", err)
			}
			log.Println("received command", cmd)
			if err = respond(conn, cmd); err != nil {
				log.Println("error responding", err)
			}
		}
	}

}
