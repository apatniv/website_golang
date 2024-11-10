package main

import (
	"github.com/alecthomas/kong"
	"github.com/apatniv/website_golang/graphs"
	"github.com/apatniv/website_golang/iter_utils"
	"log"
	"log/slog"
)

func setupLogging() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func main() {
	var cli struct {
		Example string `arg:"" enum:"strongly_connected_component,power_set" required:"true" help:"Name of the example: ${enum}"`
	}
	setupLogging()
	_ = kong.Parse(&cli)
	slog.Info("Running", "example", cli.Example)
	switch cli.Example {
	case "strongly_connected_component":
		graphs.SccExample()
	case "power_set":
		iter_utils.PowerSetExample()
	default:
		log.Panicf("Unexpected command=%v", cli.Example)
	}
}
