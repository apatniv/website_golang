package main

import (
	"github.com/apatniv/website_golang/graphs"
	"log/slog"
)

func setupLogging() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func main() {
	setupLogging()
	graphs.Example()
}
