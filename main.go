package main

import (
	"strings"

	g "xabbo.b7c.io/goearth"
	"xabbo.b7c.io/goearth/shockwave/out"
	"xabbo.b7c.io/goearth/shockwave/in"
)

var ext = g.NewExt(g.ExtInfo{
	Title:       "ModelSwap",
	Description: "",
	Version:     "1.0",
	Author:      "Scott",
})

var emoteBlock bool

func main() {
	// ext.Intercept(out.CHAT).With(handleTalk)
	ext.Intercept(in.ROOM_READY).With(handleRoomReady)
	ext.Run()
}

func handleRoomReady(e *g.InterceptArgs) {
	roomInfo := e.Packet.ReadString()
	modifiedRoomInfo := strings.replace(roomInfo, "model_z", "model_f", 1)
	e.Block()
	ext.Send(in.ROOM_READY, modifiedRoomInfo)
}

// func handleShout(e *g.InterceptArgs) {
// 	msg := e.Packet.ReadString()
// 	handleEmotes(msg)
// 	e.Block()
// 	if !emoteBlock {
// 		ext.Send(out.CHAT, msg)
// 	}
// }
// func handleTalk(e *g.InterceptArgs) {
// 	msg := e.Packet.ReadString()
// 	handleEmotes(msg)
// 	e.Block()
// 	if !emoteBlock {
// 		ext.Send(out.SHOUT, msg)
// 	}
// }

// func handleEmotes(msg string) {
// 	if msg == "/dance" {
// 		emoteBlock = true
// 		ext.Send(out.DANCE)
// 	} else if strings.Contains(msg, "o/") {
// 		if msg == "o/" {
// 			emoteBlock = true
// 		} else {
// 			emoteBlock = false
// 		}
// 		ext.Send(out.WAVE)
// 	} else {
// 		emoteBlock = false
// 	}
// }
