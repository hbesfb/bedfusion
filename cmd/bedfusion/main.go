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

// Validate bed input
func (s *session) Validate() error {
	if err := s.Bedfile.VerifyAndHandle(); err != nil {
		return err
	}
	// Give warnings about wrong unused variables if a
	// config file is used
	if s.ConfigFile != "" {
		s.Bedfile.WarnAboutWrongUnusedVariables()
	}
	return nil
}

func main() {
	var s session
	// Getting variables
	s.ctx = kong.Parse(&s,
		kong.Description("Another tool for sorting and merging bed files.\n\n"+
			"BedFusion follows the bed file standard outlined in: https://github.com/samtools/hts-specs/blob/94500cf76f049e898dec7af23097d877fde5894e/BEDv1.pdf \n\n"+
			"Read priority order: 1. flags 2. configuration file 3. environmental variables"),
		kong.Configuration(kongyaml.Loader),
		kong.UsageOnError(),
	)
	s.ctx.FatalIfErrorf(s.run())
}

func (s *session) run() (error, string) {
	// Read bed file
	if err := s.Bedfile.Read(); err != nil {
		return err, "while reading"
	}
	if !s.Bedfile.NoMerge {
		// Merge and pad lines
		s.Bedfile.MergeAndPadLines()
	} else {
		// Pad lines
		if s.Bedfile.Padding != 0 {
			s.Bedfile.PadLines()
		}
		// Fission
		if s.Bedfile.Fission {
			s.Bedfile.SplitLines()
		}
		// Deduplicate
		if s.Bedfile.Deduplicate {
			s.Bedfile.DeduplicateLines()
		}
	}
	// Sort
	if err := s.Bedfile.Sort(); err != nil {
		return err, "while sorting"
	}
	// Write output
	if err := s.Bedfile.Write(); err != nil {
		return err, "while writing"
	}
	return nil, ""
}
