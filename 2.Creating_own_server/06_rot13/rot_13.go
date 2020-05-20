package main

import (
	"bufio"
	"fmt"
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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		ls := strings.ToLower(ln)
		bs := []byte(ls)
		bsr := rot13(bs)
		rs := string(bsr)
		fmt.Fprintln(conn, ln, "-", rs)
	}
}

func rot13(bs []byte) []byte {
	r13 := make([]byte, len(bs))
	for i, v := range bs {
		//ascii 97-122
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}

	}
	return r13
}
