package main

import (
	"context"
	_ "shared/logging"

	"shared/server/httpserver"
)

func main() {
	ctx := context.Background()
	httpserver.StartServer(ctx)
}
