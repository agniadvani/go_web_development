package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	io.WriteString(conn, "Use\nSET to set a value to a key\nGET to get a value from a key\nDEL to delete the key and the value\n\n\n")
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		ln := scanner.Text()
		fs := strings.Fields(ln) //breaking the text into words

		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n", v)

		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "unexpected value")
			}
			k := fs[1]
			v := fs[2]
			data[k] = v

		case "DEL":
			k := fs[1]
			delete(data, k)

		default:
			fmt.Fprintln(conn, "invalid command")
		}

	}
}
