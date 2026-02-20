package main

import (
	"fmt"

	"github.com/sabirkekw/YANDEX_gRPCserver/internal/cfg"
)

func main() {
	config := cfg.MustLoad()
	fmt.Print(config)
	// todo: init logger
	// todo: init database
	// todo: start server
	// todo: handle graceful shutdown
}
