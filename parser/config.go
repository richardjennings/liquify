package parser

import (
	"github.com/richardjennings/liquify/expr"
)

// A Config holds configuration information for parsing and rendering.
type Config struct {
	expr.Config
	Grammar Grammar
	Delims  []string
}

// NewConfig creates a parser Config.
func NewConfig(g Grammar) Config {
	return Config{Grammar: g}
}
