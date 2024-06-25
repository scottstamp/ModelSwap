package main

import (
	"fmt"
	"strings"

	g "xabbo.b7c.io/goearth"
)

var ext = g.NewExt(g.ExtInfo{
	Title:       "Talk To Shout",
	Description: "Shout To Talk, reverses the functionality for ease of use :)",
	Version:     "1.0",
	Author:      "Eduard, b7",
})
var (
	outShout = g.Out.Id("SHOUT")
	outChat  = g.Out.Id("CHAT")
	outDance = g.Out.Id("DANCE")
	outWave  = g.Out.Id("WAVE")
)

var emoteBlock bool

func main() {
	ext.Intercept(outChat).With(handleTalk)
	ext.Intercept(outShout).With(handleShout)
	ext.Run()
}

func handleShout(e *g.InterceptArgs) {
	msg := e.Packet.ReadString()
	handleEmotes(msg)
	e.Block()
	if !emoteBlock {
		ext.Send(outChat, msg)
	}
}
func handleTalk(e *g.InterceptArgs) {
	msg := e.Packet.ReadString()
	handleEmotes(msg)
	e.Block()
	if !emoteBlock {
		ext.Send(outShout, msg)
	}
}

func handleEmotes(msg string) {
	fmt.Println(msg)
	if msg == "/dance" {
		emoteBlock = true
		ext.Send(outDance)
	} else if strings.Contains(msg, "o/") {
		if msg == "o/" {
			emoteBlock = true
		} else {
			emoteBlock = false
		}
		ext.Send(outWave)
	} else {
		emoteBlock = false
	}
}
