package server

import (
	"flag"
	"os"
	"time"
)

var (
	addr = flag.String("addr", ":", os.Getenv("PORT"))
	cert = flag.String("cert", "", "")
	key = flag.String.("key", "", "")
)

func Run() error{
	flag.Parse()

	if *addr == ":" {
		*addr = ":8080"
	}

	app.Get("/", handler.Welcome)
	app.Get("/room/create", handler.RoomCreate)
	app.Get("/room/:uuid", handler.Room)
	app.Get("/room/:uuid/websocket", websocket.New(handler.RoomChatWebsocket))
	app.Get("/room/:uuid/chat/websocket", websocket.New(handler.RoomViewerWebsocket))
	app.Get("/stream/:ssuid", handler.Stream)
	app.Get("/stream/:ssuid/websocket")
	app.Get("/stream/ssuid/chat/websocket")
	app.Get("stream/:ssuid/viewer/websocket")
}