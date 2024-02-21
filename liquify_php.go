package liquify

import (
	"bytes"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/richardjennings/liquify/expr"
	"github.com/richardjennings/liquify/parser"
)

type PHP struct {
	TagParsers map[string]func(b *bytes.Buffer, t *parser.ASTTag, p PHP) error
}

func (p PHP) Transpile(l *Liquified) ([]byte, error) {
	if p.TagParsers == nil {
		p.TagParsers = make(map[string]func(b *bytes.Buffer, t *parser.ASTTag, p PHP) error)
	}
	buf := bytes.NewBuffer(nil)
	if len(l.FrontMatter) != 0 {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, err := json.Marshal(l.FrontMatter)
		if err != nil {
			return nil, err
		}
		buf.Write([]byte(fmt.Sprintf(`<?php $page = json_encode('%s');?>%s`, string(b), "\n")))
	}
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
			b.Write([]byte(`<?php } else { ?>`))
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
			b.Write([]byte(`<?php } ?>`))
		case "include":
			// first create an array of values to passs,
			// then call a "render" function with the template and the array
			argsExpr := n.Expr.(expr.IncludeArgExpr)
			if len(argsExpr.Exprs) > 1 {
				// create the array
				b.Write([]byte(`<?php $values = [`))
				for i, v := range argsExpr.Exprs[1:] {
					if i%2 == 0 {
						b.Write([]byte(fmt.Sprintf(`%s=>`, p.Expr(v))))
					} else {
						b.Write([]byte(fmt.Sprintf(`%s,`, p.Expr(v))))
					}
				}
				b.Write([]byte(`]; ?>`))
			}
			b.Write([]byte(fmt.Sprintf(`<?php render(%s, $values);?>`, p.Expr(argsExpr.Exprs[0]))))
		default:
			if pfunc, ok := p.TagParsers[n.Name]; ok {
				if err := pfunc(b, n, p); err != nil {
					return err
				}
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
			return fmt.Sprintf(`%s["%s"]`, p.Expr(e.V), e.Name)
		default:
			panic("variable names must be IdentExprs ?")
		}
	case expr.IdentExpr:
		return fmt.Sprintf("$%s", e.V)
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
