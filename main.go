package main

import (
	"strings"

	g "xabbo.b7c.io/goearth"
	"xabbo.b7c.io/goearth/shockwave/out"
)

var ext = g.NewExt(g.ExtInfo{
	Title:       "Talk To Shout",
	Description: "Shout To Talk, reverses the functionality for ease of use :)",
	Version:     "1.0",
	Author:      "Eduard, b7",
})

var emoteBlock bool

func main() {
	ext.Intercept(out.CHAT).With(handleTalk)
	ext.Intercept(out.SHOUT).With(handleShout)
	ext.Run()
}

func handleShout(e *g.InterceptArgs) {
	msg := e.Packet.ReadString()
	handleEmotes(msg)
	e.Block()
	if !emoteBlock {
		ext.Send(out.CHAT, msg)
	}
}
func handleTalk(e *g.InterceptArgs) {
	msg := e.Packet.ReadString()
	handleEmotes(msg)
	e.Block()
	if !emoteBlock {
		ext.Send(out.SHOUT, msg)
	}
}

func handleEmotes(msg string) {
	if msg == "/dance" {
		emoteBlock = true
		ext.Send(out.DANCE)
	} else if strings.Contains(msg, "o/") {
		if msg == "o/" {
			emoteBlock = true
		} else {
			emoteBlock = false
		}
		ext.Send(out.WAVE)
	} else {
		emoteBlock = false
	}
}
