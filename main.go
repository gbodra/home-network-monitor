package main

import (
	"log"

	"github.com/ghedo/go.pkt/capture/pcap"
	"github.com/ghedo/go.pkt/layers"
	"github.com/ghedo/go.pkt/packet"
)

func main() {
	src, err := pcap.Open("en0")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	// you may configure the source further, e.g. by activating
	// promiscuous mode.

	err = src.Activate()
	if err != nil {
		log.Fatal(err)
	}

	i := 0

	for {
		buf, err := src.Capture()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("PACKET!!!")

		pkt, err := layers.UnpackAll(buf, packet.Eth)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(pkt)

		if i == 5 {
			break
		}
		i++

		// do something with the packet
	}
}
