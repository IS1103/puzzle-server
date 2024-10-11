package main

import (
	"puzzle-server/protocol/httpServer"
)

func main() {
	httpServer.StartServer("8080")
}
