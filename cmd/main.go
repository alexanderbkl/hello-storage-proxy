package main

import (
	"fmt"

	"github.com/Hello-Storage/hello-storage-proxy/internal/commands"
)

func main() {
	commands.Start()

	fmt.Println("Server running!!")
}
