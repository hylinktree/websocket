// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type imeter struct {
	Timestamp                                   int64
	PhaseVoltageR, PhaseVoltageS, PhaseVoltageT float64
	KWatth                                      float64
}

func makeJavaTimestamp(t time.Time) int64 {
	return int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
}

var addr = flag.String("addr", "localhost:8080", "http service address")
var interval = flag.Int64("i", 1000, "interval in ms")

func main() {
	// wiscmd := flag.NewFlagSet("wiscon", flag.ExitOnError)

	// fmt.Println(os.Args)
	// wiscmd.Parse(os.Args[1:])
	// fmt.Println("interval is", *pinterval)

	flag.Parse()
	// log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/report"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	cnt := 0

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			cnt++
			if (cnt % 100) == 0 {
				log.Printf("%d.recv: %s", cnt, message)
			}
		}
	}()

	ticker := time.NewTicker(time.Duration(*interval) * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:

			m := imeter{0, 1.1, 2.2, 3.3, 4.4}

			// m.time = t.String()
			m.Timestamp = makeJavaTimestamp(t) //t.UnixNano()
			// fmt.Printf("Sending [%#v]\n", m)
			err := c.WriteJSON(m)
			// err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}

		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
