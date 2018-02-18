package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type time struct {
	index int
	time  string
	city  string
}

func main() {
	count := len(os.Args[1:])
	out := make(chan time, count)
	for i, input := range os.Args[1:] {
		str := strings.Split(input, "=")
		city := str[0]
		port := strings.Split(str[1], ":")[1]
		go func(i int) {
			conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			reader := bufio.NewReader(conn)
			for {
				bytes, _, err := reader.ReadLine()
				if err != nil {
					log.Fatal(err)
				}
				out <- time{i, string(bytes), city}
			}
		}(i)
	}

	times := make([]string, count)
	for i := 0; i < count; i++ {
		t := <-out
		times[t.index] = fmt.Sprintf("%s: %s", t.city, t.time)
	}
	fmt.Println(strings.Join(times, ", "))
}
