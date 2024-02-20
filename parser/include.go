package parser

import (
	"errors"
	"github.com/richardjennings/liquify/expr"
	"unicode"
)

// IncludeParser constructs an AST representation of an Include Tag and its arguments.
type (
	IncludeParser struct {
	}
	ParseState int
)

const (
	ExpectFilePath ParseState = iota
	ExpectArgName
	ExpectEquals
	ExpectArgVal
)

// Parse returns a slice of ASTNode representing the include tag arguments
func (ip IncludeParser) Parse(argLiteral string) (expr.IncludeArgExpr, error) {
	var state ParseState
	var val []rune
	var exprs []expr.Expr
	var sq bool
	var dq bool
	var ws bool
	ret := expr.IncludeArgExpr{}

	for _, v := range argLiteral {

		if ws {
			// allow consuming whitespace when indicated by ws
			if unicode.IsSpace(v) {
				continue
			}
		}

		switch state {

		// we expect the first argument to be the filepath to be included
		// this is expressed without quotes and terminated by a space
		case ExpectFilePath:
			if unicode.IsSpace(v) {
				state = ExpectArgName
				exprs = append(exprs, expr.LiteralExpr{V: string(val)})
				val = []rune{}
				ws = true
				continue
			}
			val = append(val, v)
			continue

		// plain string literal
		case ExpectArgName:
			ws = false // disable consuming whitespace

			// a space or equals ends the arg name
			if unicode.IsSpace(v) || v == '=' {
				if unicode.IsSpace(v) {
					ws = true
					state = ExpectEquals
				} else {
					state = ExpectArgVal
				}
				exprs = append(exprs, expr.LiteralExpr{V: string(val)})
				val = []rune{}
				continue
			}
			// @todo check alpha
			val = append(val, v)
			continue

		case ExpectEquals:
			if v != '=' {
				return ret, errors.New("expected =")
			}
			state = ExpectArgVal

		case ExpectArgVal:
			// can be single quoted value, double quoted value or an expression
			switch true {
			// expecting a double quote to close
			case !sq && dq && v == '\'':
				return ret, errors.New("unexpected '")
			// expecting a single quote to close
			case !dq && sq && v == '"':
				return ret, errors.New("unexpected \"")
			// start single quote string literal
			case !sq && !dq && v == '\'':
				sq = true
				continue
			// start double quote string literal
			case !dq && !sq && v == '"':
				dq = true
				continue
			// finished reading single or double quote string literal
			case sq && v == '\'' || dq && v == '"':
				exprs = append(exprs, expr.LiteralExpr{V: string(val)})
				val = []rune{}
				state = ExpectArgName
				sq = false
				dq = false
				ws = true // allow consuming whitespace
				continue

			default:
				// if it is not single or double quoted literal and there is a space,
				// this means we have an expression
				if unicode.IsSpace(v) && !sq && !dq {
					// we have an expression
					e, err := expr.Parse(string(val))
					if err != nil {
						return ret, err
					}
					exprs = append(exprs, e)
					val = []rune{}
					state = ExpectArgName
					sq = false
					dq = false
					ws = true // allow consuming whitespace

					continue
				}
				// if neither quotes, we should be able to read up to the next space,
				// and then get the yacc parser to evaluate as an expression
				val = append(val, v)
			}
		}

	}

	if len(val) > 0 {
		// we have an expression
		e, err := expr.Parse(string(val))
		if err != nil {
			return ret, err
		}
		exprs = append(exprs, e)
	}

	if len(exprs) == 0 {
		return ret, errors.New("missing filepath")
	}

	ret.Exprs = exprs
	return ret, nil
}
