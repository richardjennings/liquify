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
		case "if":
			b.Write([]byte(fmt.Sprintf(`<?php if (%s) { ?>`, p.trans(n.Expr))))
		case "else":
			b.Write([]byte(fmt.Sprintf(`<?php } else { ?>`)))
		case "endif":
			b.Write([]byte("<?php } ?>"))
		case "assign":
			b.Write([]byte(fmt.Sprintf(`<?php $%s = %s;?>`, n.Expr.(expr.AssignmentStmt).Variable, p.trans(n.Expr.(expr.AssignmentStmt).ValueFn))))
		default:
			fmt.Println(n)
		}
	case *parser.ASTText:
		b.Write([]byte(n.Source))
	default:
		fmt.Println(n)
	}
	return nil
}

func (p PHP) trans(e expr.Expr) string {
	switch e := e.(type) {
	case expr.LiteralExpr:
		switch v := e.V.(type) {
		case string:
			return fmt.Sprintf(`"%s"`, v)
		}
	case expr.ValStmt:
		return p.trans(e.ValueFn)
	case expr.EqExpr:
		return fmt.Sprintf(`%s == %s`, p.trans(e.A), p.trans(e.B))
	case expr.PropertyExpr:
		switch e.V.(type) {
		case expr.IdentExpr:
			return fmt.Sprintf(`$%s["%s"]`, p.trans(e.V), e.Name)
		default:
			panic("variable names must be IdentExprs ?")
		}
	case expr.IdentExpr:
		return fmt.Sprintf("%s", e.V)
	default:
		fmt.Println(e)

	}
	return ""
}
