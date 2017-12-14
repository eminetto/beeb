package main

import (
	"flag"

	"github.com/gen2brain/beeep"
)

func main() {
	title := flag.String("t", "Title", "Notification title")
	msg := flag.String("m", "Message", "Message body")
	flag.Parse()
	err := beeep.Notify(*title, *msg)
	if err != nil {
		panic(err)
	}
}
