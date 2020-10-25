package main

import (
	"github.com/rai-wtnb/accomplist-api/db"
	"github.com/rai-wtnb/accomplist-api/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
