package main

import (
	"wgsync"

	"./channelsync"
)

func main() {
	channelsync.ChannelSyncMain()
	wgsync.WgSyncMain()
}
