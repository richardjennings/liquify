package liquify

import (
	"github.com/adrg/frontmatter"
	"github.com/richardjennings/liquify/parser"
	"log"
	"os"
)

type Liquified struct {
	Path        string
	FrontMatter map[string]interface{}
	Ast         parser.ASTNode
}

func Liquify(path string, config parser.Config) (*Liquified, error) {
	l := &Liquified{Path: path}
	fh, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = fh.Close() }()
	fm := make(map[string]interface{})
	template, err := frontmatter.Parse(fh, &fm)
	fh.Close()
	l.FrontMatter = fm
	loc := parser.SourceLoc{}
	n, err := config.Parse(string(template), loc)
	if err != nil {
		return l, err
	}
	l.Ast = n
	return l, nil
}
