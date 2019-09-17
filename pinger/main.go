package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dmichael/go-multicast/multicast"
	"github.com/urfave/cli"
)

const (
	defaultMulticastAddress = "239.0.0.0:9999"
)

type MyStruct struct {
	Index      int       `json:"index"`
	Aggiornato bool      `json:"aggiornato"`
	Timestamp  time.Time `json:"timestamp"`
}

func main() {

	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		address := c.Args().Get(0)
		if address == "" {
			address = defaultMulticastAddress
		}
		fmt.Printf("Broadcasting to %s\n", address)
		ping(address)
		comunicaModifica(address)
		return nil
	}

	app.Run(os.Args)
}

func ping(addr string) {

	testStruct := MyStruct{Index: 1000, Aggiornato: false}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(testStruct)

	reqBodyBytes.Bytes() // this is the []byte

	conn, err := multicast.NewBroadcaster(addr)
	if err != nil {
		log.Fatal(err)
	}

	conn.Write(reqBodyBytes.Bytes())
	time.Sleep(1 * time.Second)
}

func comunicaModifica(addr string) {
	ultimoarchivio := MyStruct{Index: 1000, Aggiornato: false, Timestamp: time.Now()}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(ultimoarchivio)

	reqBodyBytes.Bytes() // this is the []byte

	conn, err := multicast.NewBroadcaster(addr)
	if err != nil {
		log.Fatal(err)
	}

	conn.Write(reqBodyBytes.Bytes())
	time.Sleep(1 * time.Second)
}
