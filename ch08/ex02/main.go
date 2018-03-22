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

const (
	statusTransferStarting      = 125
	statusCommandOk             = 200
	statusCommandNotImplemented = 202
	statusName                  = 215
	statusReady                 = 220
	statusLoggedOut             = 221
	statusClosingDataConnection = 226
	statusLoggedIn              = 230
	statusFileActionCompleted   = 250
	statusPathCreated           = 257

	statusUserOK                      = 331
	statusRequestedFileActionNotTaken = 450
	statusCommandNotImplemented502    = 502
	statusActionNotTaken              = 550
)

type clientConn struct {
	conn    net.Conn
	r       *textproto.Reader
	current string
}

func (cc clientConn) close() error {
	return cc.conn.Close()
}

func (cc clientConn) write(code int, input string) error {
	var err error
	if input == "" {
		_, err = io.WriteString(cc.conn, fmt.Sprintf("%d\n", code))
	} else {
		_, err = io.WriteString(cc.conn, fmt.Sprintf("%d %s\n", code, input))
	}
	return err
}

func newClientConn(conn net.Conn) *clientConn {
	var cc clientConn
	cc.conn = conn
	cc.r = textproto.NewReader(bufio.NewReader(conn))
	cc.current = "."
	return &cc
}

func main() {
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
	cc.write(statusCommandOk, "")
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
			} else {
				cc.write(statusCommandOk, dir)
			}
		case "LIST":
			list := getList(cc.current)
			err := cc.write(statusCommandOk, list)
			if err != nil {
				log.Println(err)
			}
		case "RETR":
			// 指定ファイルをクライアントに送信
			bytes, err := ioutil.ReadFile(cmd[1])
			if err != nil {
				log.Println(err)
			}
			// bytesを使ってクライアント側に送る
			if err != nil {
				log.Println(err)
			}
			cc.conn.Write(bytes)
		case "EXIT":
			err := cc.close()
			if err != nil {
				log.Println(err)
			}
			return
		case "USER":
			cc.write(statusUserOK, "")
		case "PASS":
			cc.write(statusLoggedIn, fmt.Sprintf("Hello %s!", cmd[1]))
		case "SYST":
			err := cc.write(statusName, "UNIX")
			if err != nil {
				log.Printf("%v", err)
			}
		case "FEAT", "LPRT", "EPRT", "LPSV", "EPSV":
			if err := cc.write(statusCommandNotImplemented502, ""); err != nil {
				log.Printf("%v", err)
			}
		case "PWD":
			cc.write(statusPathCreated, cc.current)

		default:
			// not implemented
			log.Println("not implemented")
			cc.write(statusCommandNotImplemented, "not implemented")
		}
	}
}

func getList(dir string) string {
	var nameList string
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		fmt.Printf("%s ", file.Name())
		nameList += file.Name() + " "
	}
	return nameList
}

func changeDir(to string) (string, bool) {
	_, err := ioutil.ReadDir(to)
	if err != nil {
		message := fmt.Sprintf("dont exist %s\n", to)
		log.Print(message)
		return message, false
	}
	return to, true
}
