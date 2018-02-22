package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/textproto"
	"strings"
)

type clientConn struct {
	conn    net.Conn
	r       *textproto.Reader
	current string
}

func (cc clientConn) close() error {
	return cc.conn.Close()
}

func newClientConn(conn net.Conn) *clientConn {
	var cc clientConn
	cc.conn = conn
	cc.r = textproto.NewReader(bufio.NewReader(conn))
	cc.current = "."
	return &cc
}

func main() {
}

func test() {
	listen, err := net.Listen("tcp", ":21")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	log.Println("connected")
	cc := newClientConn(conn)
	for {
		line, err := cc.r.ReadLine()
		if err != nil {
			if err == io.EOF {
				log.Println("disconnected")
				return
			}
			log.Println(err)
		}
		log.Println(line)

		// 各コマンドをハンドリング
		cmd := strings.Split(line, " ")
		switch cmd[0] {
		case "CWD":
			dir, ok := changeDir(cmd[1])
			if ok {
				cc.current += "/" + dir
			}
		case "NLST":
			showList(cc.current)
		case "RETR":
			// 指定ファイルをクライアントに送信
			bytes, err := ioutil.ReadFile(cmd[1])
			if err != nil {
				log.Println(err)
			}
			// bytesを使ってクライアント側に送る
			log.Println(bytes)
		case "EXIT":
			err := cc.close()
			if err != nil {
				log.Println(err)
			}
			return
		default:
			// not implemented
			log.Println("not implemented")
		}
	}
}

func showList(dir string) {
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		fmt.Printf("%s ", file.Name())
	}
}

func changeDir(to string) (string, bool) {
	_, err := ioutil.ReadDir(to)
	if err != nil {
		log.Println("dont exist such folder")
		return "", false
	}
	return to, true
}
