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
	TagParsers map[string]func(b *bytes.Buffer, t *parser.ASTTag, p PHP) error
}

func (p PHP) Transpile(l *Liquified) ([]byte, error) {
	if p.TagParsers == nil {
		p.TagParsers = make(map[string]func(b *bytes.Buffer, t *parser.ASTTag, p PHP) error)
	}
	buf := bytes.NewBuffer(nil)
	switch v := l.Ast.(type) {
	case *parser.ASTSeq:
		for _, n := range v.Children {
			if err := p.AstNode(buf, n); err != nil {
				return nil, err
			}
		}
	default:
		return nil, errors.New("unhandled")
	}

	return buf.Bytes(), nil
}

func (p PHP) AstNode(b *bytes.Buffer, n parser.ASTNode) error {
	switch n := n.(type) {
	case *parser.ASTTag:
		switch n.Name {
		case "if":
			b.Write([]byte(fmt.Sprintf(`<?php if (%s) { ?>`, p.Expr(n.Expr))))
		case "else":
			b.Write([]byte(fmt.Sprintf(`<?php } else { ?>`)))
		case "endif":
			b.Write([]byte("<?php } ?>"))
		case "assign":
			b.Write([]byte(fmt.Sprintf(`<?php $%s = %s;?>`, n.Expr.(expr.AssignmentStmt).Variable, p.Expr(n.Expr.(expr.AssignmentStmt).ValueFn))))
		case "capture":
			b.Write([]byte(fmt.Sprintf(`<?php $%s = "`, n.Token.Args)))
		case "endcapture":
			b.Write([]byte(`";?>`))
		case "comment":
			b.Write([]byte(`/* `))
		case "endcomment":
			b.Write([]byte(` */`))
		case "for":
			b.Write([]byte(fmt.Sprintf(`<?php for ($i, $%s in %s){ ?>`, n.Expr.(expr.LoopStmt).Variable, p.Expr(n.Expr.(expr.LoopStmt).Expr))))
		case "endfor":
			b.Write([]byte(fmt.Sprintf(`<?php } ?>`)))
		default:
			if pfunc, ok := p.TagParsers[n.Name]; ok {
				pfunc(b, n, p)
			} else {
				panic("not handled")
			}
		}
	case *parser.ASTText:
		b.Write([]byte(n.Source))
	default:
		fmt.Println(n)
	}
	return nil
}

func (p PHP) Stmt(s expr.Statement) string {
	switch s := s.(type) {
	case expr.ValStmt:
		return p.Expr(s.ValueFn)
	default:
		fmt.Println(s)
	}
	return ""
}

func (p PHP) Expr(e expr.Expr) string {
	switch e := e.(type) {
	case expr.LiteralExpr:
		switch v := e.V.(type) {
		case string:
			return fmt.Sprintf(`"%s"`, v)
		case expr.LiteralExpr:
			return p.Expr(v)
		default:
			fmt.Println(v)
		}
	case expr.ValStmt:
		return p.Expr(e.ValueFn)
	case expr.EqExpr:
		return fmt.Sprintf(`%s == %s`, p.Expr(e.A), p.Expr(e.B))
	case expr.PropertyExpr:
		switch e.V.(type) {
		case expr.IdentExpr:
			return fmt.Sprintf(`$%s["%s"]`, p.Expr(e.V), e.Name)
		default:
			panic("variable names must be IdentExprs ?")
		}
	case expr.IdentExpr:
		return fmt.Sprintf("%s", e.V)
	case expr.FilterExpr:
		// not sure about this as likely very specific implementation details
		s := ""
		for _, v := range e.Args {
			s += p.Expr(v)
		}
		return fmt.Sprintf("%s /* filter %s %s */", p.Expr(e.V), e.Name, s)
	default:
		fmt.Println(e)

	}
	return ""
}
