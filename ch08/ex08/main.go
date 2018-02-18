package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func main() {
	log.Print("start")
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c := make(chan net.Conn)
		go func() {
			conn, err := l.Accept()
			if err != nil {
				log.Print(err)
			}
			c <- conn
		}()

		select {
		case <-time.After(10 * time.Second):
			log.Println("10秒経ちました")
			l.Close()
		case c := <-c:
			log.Println("connectionを発見")
			go handleConn(c)
		}
	}
}
