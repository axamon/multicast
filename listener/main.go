package main

import (
	multic "github.com/axamon/multicast"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/dmichael/go-multicast/multicast"
	"github.com/urfave/cli"
	
)

const (
	defaultMulticastAddress = "239.0.0.0:9999"
)

var lastIndex = 2000

// type MyStruct struct {
// 	Index      int       `json:"index"`
// 	Aggiornato bool      `json:"aggiornato"`
// 	Timestamp  time.Time `json:"timestamp"`
// }

var archivio = multic.Archivio{Index: 1000, Aggiornato: false, Timestamp: time.Now()}

func main() {

	app := cli.NewApp()

	fmt.Println(archivio)

	app.Action = func(c *cli.Context) error {
		address := c.Args().Get(0)
		if address == "" {
			address = defaultMulticastAddress
		}
		fmt.Printf("Listening on %s\n", address)
		multicast.Listen(address, msgHandler)
		return nil
	}

	app.Run(os.Args)
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
	log.Println(string(b[:n]))
	e := new(multic.Archivio)
	json.Unmarshal(b[:n], &e)
	log.Println(e.Index)

	if e.Timestamp.After(archivio.Timestamp) {
		fmt.Println(e.Timestamp)
		brodcasthigherindex(defaultMulticastAddress)
	}
}

func brodcasthigherindex(addr string) {
	conn, err := multicast.NewBroadcaster(addr)
	if err != nil {

		log.Fatal(err)
	}

	conn.Write([]byte("Bisogna aggiornare"))

	return
}

// func ping(addr string) {
// 	conn, err := multicast.NewBroadcaster(addr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for {
// 		conn.Write([]byte("hello, world\n"))
// 		time.Sleep(1 * time.Second)
// 	}
// }
