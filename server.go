package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
	"io"
	"bytes"
	"strings"
)

// conns is a map holding all the currently active connections to subscribers
var conns = make(map[string]net.Conn)

func addConn(remoteAddr string, conn net.Conn) {
	fmt.Printf("Received new connection on %s\n", remoteAddr)
	conns[remoteAddr] = conn
}

func main() {
	// Accept connections to clients
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	uri := requestURI(conn)
	switch uri {
	case "/register":
		addConn(conn.RemoteAddr().String(), conn)
		conn.Write([]byte("Registering!\n"))
	case "/message":
		defer conn.Close()
		for k, v := range conns {
			fmt.Printf("Writing to conn: %s", k)
			v.Write([]byte("Message!\n"))
		}
		conn.Write([]byte("Thanks for stopping by!\n"))
		// TODO below would pluck the correct subscriber
		//conns[conn.RemoteAddr().String()].Write([]byte("Message!"))
	default:
		conn.Write([]byte("Invalid Request URI"))
	}
}

func requestURI(conn net.Conn) string {
	var buffer bytes.Buffer
	// TODO error handling
	bytes, _, _ := bufio.NewReader(io.MultiReader(&buffer, conn)).ReadLine()
	return strings.Split(strings.TrimPrefix(string(bytes), "GET "), " ")[0]
}
