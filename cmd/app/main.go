package main

import (
	"log/slog"

	"github.com/Abiskar-Timsina/prometheus-showcase/core/initializers"
)

func main() {
	slog.Info("Server Started")
	initializers.InitializeServer(8000)
}
