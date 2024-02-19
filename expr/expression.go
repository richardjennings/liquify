package expr

type Expr interface{}

type (
	LiteralExpr struct {
		V interface{}
	}
	IdentExpr struct {
		V interface{}
	}
	PropertyExpr struct {
		V    interface{}
		Name string
	}
	IndexExpr struct {
		V interface{}
	}
	RangeExpr struct {
		V interface{}
	}
	FilterExpr struct {
		V    interface{}
		Name string
		Args []LiteralExpr
	}
	EqExpr struct {
		A Expr
		B Expr
	}
	NeqExpr struct {
		A Expr
		B Expr
	}
	GtExpr struct {
		A Expr
		B Expr
	}
	LtExpr struct {
		A Expr
		B Expr
	}
	GeExpr struct {
		A Expr
		B Expr
	}
	LeExpr struct {
		A Expr
		B Expr
	}
	ContainsExpr struct {
		A Expr
		B Expr
	}
	AndExpr struct {
		A Expr
		B Expr
	}
	OrExpr struct {
		A Expr
		B Expr
	}
)

type ParseValue struct {
	Stmt Statement
}
