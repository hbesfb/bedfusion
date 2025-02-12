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
	return nil
}

func main() {
	var s session
	// Getting variables
	s.ctx = kong.Parse(&s,
		kong.Description("Another tool for sorting and merging bed files.\n\n"+
			"BedFusion follows the bed file standard outlined in: https://github.com/samtools/hts-specs/blob/94500cf76f049e898dec7af23097d877fde5894e/BEDv1.pdf \n\n"+
			"Read priority order: 1. flags 2. configuration file 3. environmental variables \n\n"+
			"Order of actions: 1. reading files 2. padding(*) 3. merging(*)/deduplication(*) 4. sorting 5. writing output (* = can be turned on/off using flags)"),
		kong.Vars{
			// Sorting types
			"lexST":  bed.LexST,
			"natST":  bed.NatST,
			"ccsST":  bed.CcsST,
			"fidxST": bed.FidxST,
			// Padding types
			"failPT":  bed.SafePT,
			"warnPT":  bed.LaxPT,
			"forcePT": bed.ForcePT,
		},
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
		if err := s.Bedfile.MergeAndPadLines(); err != nil {
			return err, "while padding"
		}
	} else {
		// Pad lines
		if s.Bedfile.Padding != 0 {
			if err := s.Bedfile.PadLines(); err != nil {
				return err, "while padding"
			}
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
