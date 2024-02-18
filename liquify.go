package liquify

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/adrg/frontmatter"
	"github.com/richardjennings/liquify/expr"
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

type PHP struct {
}

func (p PHP) Transpile(l *Liquified) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	switch v := l.Ast.(type) {
	case *parser.ASTSeq:
		for _, n := range v.Children {
			if err := p.transpile(buf, n); err != nil {
				return nil, err
			}
		}
	default:
		return nil, errors.New("unhandled")
	}

	return buf.Bytes(), nil
}

func (p PHP) transpile(b *bytes.Buffer, n parser.ASTNode) error {
	switch n := n.(type) {
	case *parser.ASTTag:
		switch n.Name {
		case "assign":
			b.Write([]byte(fmt.Sprintf(`<?php $%s = %s;?>`, n.Expr.(expr.AssignmentStmt).Variable, p.trans(n.Expr.(expr.AssignmentStmt).ValueFn))))
		}
	}
	return nil
}

func (p PHP) trans(e expr.Expr) string {
	switch e := e.(type) {
	case expr.LiteralExpr:
		switch v := e.V.(type) {
		case string:
			return fmt.Sprintf(`"%v"`, v)
		}
	default:
		fmt.Println(e)

	}
	return ""
}
