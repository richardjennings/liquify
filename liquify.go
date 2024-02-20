package liquify

import (
	"bytes"
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

type Exporter interface {
	Transpile(l *Liquified) ([]byte, error)
}

func LiquifyFromFile(path string, config parser.Config) (*Liquified, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	l, err := Liquify(b, config)
	if err != nil {
		return nil, err
	}
	l.Path = path
	return l, nil
}

func Liquify(content []byte, config parser.Config) (*Liquified, error) {
	l := &Liquified{}
	buf := bytes.NewBuffer(content)
	fm := make(map[string]interface{})
	template, err := frontmatter.Parse(buf, &fm)
	if err != nil {
		return nil, err
	}
	l.FrontMatter = fm
	loc := parser.SourceLoc{}
	n, err := config.Parse(string(template), loc)
	if err != nil {
		return l, err
	}
	l.Ast = n
	return l, nil
}
