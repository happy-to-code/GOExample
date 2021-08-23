package main

import (
	"fmt"
	"strings"
)

const (
	msgPrefix                  = "badger_msg_"
	heightPrefix               = "badger_height_"
	clientPrefix               = "badger_client_"
	latestHeightWithPeerPrefix = "badger_peer_height_"
	// msgSendHeightPrefix = "badger_msg_send_"
)

func main() {
	var k = fmt.Sprintf("%s%s", latestHeightWithPeerPrefix, "dfdf")
	fmt.Println(k)

	split := strings.Split(k, latestHeightWithPeerPrefix)
	fmt.Println("-->", split)
	fmt.Println("-->", split[0])
	fmt.Println("-->", split[1])

}
