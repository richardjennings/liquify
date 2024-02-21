package parser

import (
	"fmt"
	"github.com/richardjennings/liquify/expr"
)

// ASTNode is a node of an AST.
type ASTNode interface {
	SourceLocation() SourceLoc
	SourceText() string
	String() string
}

// ASTBlock represents a {% tag %}…{% endtag %}.
type ASTBlock struct {
	Token
	syntax  BlockSyntax
	Body    []ASTNode   // Body is the nodes before the first branch
	Clauses []*ASTBlock // E.g. else and elseif w/in an if
}

func (a ASTBlock) String() string {
	return fmt.Sprintf("%s %s", a.Token, a.Body)
}

// ASTRaw holds the text between the start and end of a raw tag.
type ASTRaw struct {
	Slices []string
	sourcelessNode
}

func (a ASTRaw) String() string {
	return fmt.Sprintf("%s", a.Slices)
}

// ASTTag is a tag {% tag %} that is not a block start or end.
type ASTTag struct {
	Token
	Expr expr.Expr
}

func (a ASTTag) String() string {
	return fmt.Sprintf("%s %s", a.Token.Name, a.Token.Args)
}

// ASTText is a text span, that is rendered verbatim.
type ASTText struct {
	Token
}

func (a ASTText) String() string {
	return a.Source
}

// ASTObject is an {{ object }} object.
type ASTObject struct {
	Token
	Expr expr.Expr
}

func (a ASTObject) String() string {
	return "" //@todo
}

// ASTSeq is a sequence of nodes.
type ASTSeq struct {
	Children []ASTNode
	sourcelessNode
}

func (a ASTSeq) String() string {
	str := ""
	for _, v := range a.Children {
		str += v.String()
	}
	return str
}

// It shouldn't be possible to get an error from one of these node types.
// If it is, this needs to be re-thought to figure out where the source
// location comes from.
type sourcelessNode struct{}

func (n *sourcelessNode) SourceLocation() SourceLoc {
	panic("unexpected call on sourceless node")
}

func (n *sourcelessNode) SourceText() string {
	panic("unexpected call on sourceless node")
}
