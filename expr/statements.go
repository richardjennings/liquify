package expr

import (
	"errors"
	"fmt"
)

// These strings match lexer tokens.
const (
	AssignStatementSelector = "%assign "
	CycleStatementSelector  = "{%cycle "
	LoopStatementSelector   = "%loop "
	WhenStatementSelector   = "{%when "
)

// A Statement is the result of parsing a string.
type Statement interface{}

//struct{ ParseValue }

// Expr returns a statement's Expr function.
// func (s *Statement) Expr() Expr { return &Expr{s.val} }

// An AssignmentStmt is a parse of an {% assign %} statement
type AssignmentStmt struct {
	Variable string
	ValueFn  Expr
}

type ValStmt struct {
	ValueFn Expr
}

// A CycleStmt is a parse of an {% assign %} statement
type CycleStmt struct {
	Group  string
	Values []string
}

// A LoopStmt is a parse of a {% loop %} statement
type LoopStmt struct {
	Variable string
	Expr     Expr
	LoopModifiers
}

type LoopModifiers struct {
	Limit    Expr
	Offset   Expr
	Cols     Expr
	Reversed bool
}

// A WhenStmt is a parse of a {% when %} clause
type WhenStmt struct {
	Exprs []Expr
}

// SyntaxError represents a syntax error. The yacc-generated compiler
// doesn't use error returns; this lets us recognize them.
type SyntaxError string

func (e SyntaxError) Error() string { return string(e) }

func Parse(source string) (Statement, error) {
	p, err := parse(source)
	if err != nil {
		return nil, err
	}
	if p.Stmt != nil {
		return p.Stmt, nil
	}

	return nil, errors.New("no statement")
}

// ParseStatement parses an statement into an Expr that can evaluated to return a
// structure specific to the statement.
func ParseStatement(sel, source string) (Statement, error) {
	p, err := parse(sel + source)
	if err != nil {
		return nil, err
	}
	if p.Stmt != nil {
		return p.Stmt, nil
	}

	return nil, errors.New("no statement")
}

func parse(source string) (p *ParseValue, err error) {
	// FIXME hack to recognize EOF
	lex := newLexer([]byte(source + ";"))
	n := yyParse(lex)
	if n != 0 {
		return nil, SyntaxError(fmt.Errorf("syntax error in %q", source).Error())
	}
	return &lex.ParseValue, nil
}

func MakeLiteralExpr(expr Expr) LiteralExpr {
	return LiteralExpr{V: expr}
}

func MakeIdentExpr(expr Expr) IdentExpr {
	return IdentExpr{V: expr}
}

func MakePropertyExpr(expr Expr, name string) PropertyExpr {
	return PropertyExpr{V: expr, Name: name}
}

func MakeIndexExpr(a Expr, b Expr) IndexExpr {
	return IndexExpr{V: a}
}

func MakeRangeExpr(expr Expr, b Expr) RangeExpr {
	return RangeExpr{V: expr}
}

func MakeFilterExpr(expr Expr, name string, args []LiteralExpr) FilterExpr {
	return FilterExpr{V: expr}
}

func MakeEQExpr(a Expr, b Expr) EqExpr {
	return EqExpr{A: a, B: b}
}

func MakeNEQExpr(a Expr, b Expr) NeqExpr {
	return NeqExpr{A: a, B: b}
}

func MakeGtExpr(a Expr, b Expr) GtExpr {
	return GtExpr{A: a, B: b}
}

func MakeLtExpr(a Expr, b Expr) LtExpr {
	return LtExpr{A: a, B: b}
}

func MakeGeExpr(a Expr, b Expr) GeExpr {
	return GeExpr{A: a, B: b}
}

func MakeLeExpr(a Expr, b Expr) LeExpr {
	return LeExpr{A: a, B: b}
}

func MakeContainsExpr(a Expr, b Expr) ContainsExpr {
	return ContainsExpr{A: a, B: b}
}

func MakeAndExpr(a Expr, b Expr) AndExpr {
	return AndExpr{A: a, B: b}
}

func MakeOrExpr(a Expr, b Expr) OrExpr {
	return OrExpr{A: a, B: b}
}
