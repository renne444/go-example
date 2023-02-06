package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var (
	ip         = flag.String("ip", "127.0.0.1", "server IP")
	connection = flag.Int("conn", 1, "number of socket connections")
)

func main() {
	flag.Usage = func() {
		io.WriteString(os.Stderr, `Web socket client generator
Example usage: ./client -ip=127.0.0.1 -conn=10
`)
		flag.PrintDefaults()
	}
	flag.Parse()

	fmt.Printf("ip = %s\n", *ip)
	fmt.Printf("conn = %d\n", *connection)

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:40404", Path: "/"}
	log.Printf("Connecting to %s", u.String())

	var conns []*websocket.Conn
	for i := 0; i < *connection; i++ {
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			fmt.Println("Fail to connect", i, err)
			c.Close()
		}
		conns = append(conns, c)
		defer func(i int) {
			c.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(time.Second))
			time.Sleep(time.Second)
			c.Close()
			fmt.Println("closed for " + strconv.Itoa(i) + " conn")
		}(i)

	}
	log.Printf("Finish initializing %d connections", *connection)
	tts := time.Second

	if *connection > 100 {
		tts = time.Microsecond * 5
	}
	for {
		for i := 0; i < len(conns); i++ {
			time.Sleep(tts)
			conn := conns[i]
			log.Printf("Conn %d sending msg", i)
			if err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(time.Second*5)); err != nil {
				fmt.Printf("Fail to recieving pong: %v", err)
			}
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("hello from conn %v", i)))
		}
	}
}
