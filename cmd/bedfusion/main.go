package main

import (
	"github.com/alecthomas/kong"

	kongyaml "github.com/alecthomas/kong-yaml"

	"github.com/hbesfb/bedfusion/internal/bed"
)

type session struct {
	ConfigFile kong.ConfigFlag `env:"CONFIG_FILE" short:"c" help:"The path to configuration file (must be in key-value yaml format)"`
	Bedfile    bed.Bedfile     `embed:""`
	ctx        *kong.Context
}

func main() {
	var s session
	// Getting variables
	s.ctx = kong.Parse(&s,
		kong.Description("Another tool for sorting and merging bed files.\n\n"+
			"Read priority order: 1. flags 2. configuration file 3. environmental variables"),
		kong.Configuration(kongyaml.Loader),
		kong.UsageOnError(),
	)
}
