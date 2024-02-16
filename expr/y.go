// Code generated by goyacc -o expressions/y.go expressions/y.y. DO NOT EDIT.

//line expressions/y.y:2
package expr

import __yyfmt__ "fmt"

//line expressions/y.y:2
import (
	"fmt"
)

func init() {
	// This allows adding and removing references to fmt in the rules below,
	// without having to comment and un-comment the import statement above.
	_ = ""
}

//line expressions/y.y:15
type yySymType struct {
	yys           int
	name          string
	val           interface{}
	f             Expr
	s             string
	ss            []string
	exprs   []Expr
	cycle   CycleStmt
	cyclefn  func(string) CycleStmt
	loop     LoopStmt
	loopmods LoopModifiers
	filter_params []LiteralExpr
}

const LITERAL = 57346
const IDENTIFIER = 57347
const KEYWORD = 57348
const PROPERTY = 57349
const ASSIGN = 57350
const CYCLE = 57351
const LOOP = 57352
const WHEN = 57353
const EQ = 57354
const NEQ = 57355
const GE = 57356
const LE = 57357
const IN = 57358
const AND = 57359
const OR = 57360
const CONTAINS = 57361
const DOTDOT = 57362

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LITERAL",
	"IDENTIFIER",
	"KEYWORD",
	"PROPERTY",
	"ASSIGN",
	"CYCLE",
	"LOOP",
	"WHEN",
	"EQ",
	"NEQ",
	"GE",
	"LE",
	"IN",
	"AND",
	"OR",
	"CONTAINS",
	"DOTDOT",
	"'.'",
	"'|'",
	"'<'",
	"'>'",
	"';'",
	"'='",
	"':'",
	"','",
	"'['",
	"']'",
	"'('",
	"')'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 106

var yyAct = [...]int8{
	9, 47, 42, 18, 8, 14, 15, 23, 10, 11,
	10, 11, 76, 34, 3, 4, 5, 6, 25, 25,
	60, 41, 43, 43, 38, 46, 44, 51, 52, 53,
	54, 55, 56, 57, 58, 12, 25, 12, 39, 62,
	26, 26, 24, 61, 77, 63, 62, 64, 45, 66,
	65, 68, 24, 14, 15, 71, 7, 25, 26, 69,
	70, 13, 27, 28, 31, 32, 72, 73, 75, 33,
	59, 36, 37, 30, 29, 25, 25, 80, 2, 26,
	81, 27, 28, 31, 32, 78, 79, 21, 33, 49,
	50, 35, 30, 29, 16, 19, 48, 26, 26, 1,
	74, 20, 40, 17, 22, 67,
}

var yyPact = [...]int16{
	6, -1000, 36, 89, 91, 82, 4, -1000, 20, 69,
	-1000, -1000, 4, -1000, 4, 4, -2, 13, -6, -1000,
	1, 32, 0, 68, 84, -1000, 4, 4, 4, 4,
	4, 4, 4, 4, 50, -12, -1000, -1000, 4, -1000,
	-1000, 91, -1000, 91, -1000, 4, -1000, -1000, 4, -1000,
	4, 29, 11, 11, 11, 11, 11, 11, 11, 4,
	-1000, 30, 11, -5, -5, 20, 68, -16, 11, -1000,
	12, -1000, -1000, -1000, 80, -1000, 4, -1000, -1000, 4,
	11, 11,
}

var yyPgo = [...]int8{
	0, 0, 56, 4, 78, 105, 104, 1, 103, 102,
	2, 101, 100, 3, 99,
}

var yyR1 = [...]int8{
	0, 14, 14, 14, 14, 14, 8, 9, 9, 10,
	10, 6, 7, 7, 13, 11, 12, 12, 12, 1,
	1, 1, 1, 1, 1, 3, 3, 3, 5, 5,
	2, 2, 2, 2, 2, 2, 2, 2, 4, 4,
	4,
}

var yyR2 = [...]int8{
	0, 2, 5, 3, 3, 3, 2, 3, 1, 0,
	3, 2, 0, 3, 1, 4, 0, 2, 3, 1,
	1, 2, 4, 5, 3, 1, 3, 4, 1, 3,
	1, 3, 3, 3, 3, 3, 3, 3, 1, 3,
	3,
}

var yyChk = [...]int16{
	-1000, -14, -4, 8, 9, 10, 11, -2, -3, -1,
	4, 5, 31, 25, 17, 18, 5, -8, -13, 4,
	-11, 5, -6, -1, 22, 7, 29, 12, 13, 24,
	23, 14, 15, 19, -1, -4, -2, -2, 26, 25,
	-9, 27, -10, 28, 25, 16, 25, -7, 28, 5,
	6, -1, -1, -1, -1, -1, -1, -1, -1, 20,
	32, -3, -1, -13, -13, -3, -1, -5, -1, 30,
	-1, 25, -10, -10, -12, -7, 28, 32, 5, 6,
	-1, -1,
}

var yyDef = [...]int8{
	0, -2, 0, 0, 0, 0, 0, 38, 30, 25,
	19, 20, 0, 1, 0, 0, 0, 0, 9, 14,
	0, 0, 0, 12, 0, 21, 0, 0, 0, 0,
	0, 0, 0, 0, 25, 0, 39, 40, 0, 3,
	6, 0, 8, 0, 4, 0, 5, 11, 0, 26,
	0, 0, 31, 32, 33, 34, 35, 36, 37, 0,
	24, 0, 25, 9, 9, 16, 12, 27, 28, 22,
	0, 2, 7, 10, 15, 13, 0, 23, 17, 0,
	29, 18,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	31, 32, 3, 3, 28, 3, 21, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 27, 25,
	23, 26, 24, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 29, 3, 30, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 22,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions/y.y:45
		{
			yylex.(*lexer).Stmt = ValStmt{ValueFn: yyDollar[1].f}
		}
	case 2:
		yyDollar = yyS[yypt-5 : yypt+1]
//line expressions/y.y:46
		{
			yylex.(*lexer).Stmt = AssignmentStmt{yyDollar[2].name, yyDollar[4].f}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:49
		{
			yylex.(*lexer).Stmt = yyDollar[2].cycle
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:50
		{
			yylex.(*lexer).Stmt = yyDollar[2].loop
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:51
		{
			yylex.(*lexer).Stmt = WhenStmt{yyDollar[2].exprs}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions/y.y:54
		{
			yyVAL.cycle = yyDollar[2].cyclefn(yyDollar[1].s)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:57
		{
			h, t := yyDollar[2].s, yyDollar[3].ss
			yyVAL.cyclefn = func(g string) CycleStmt { return CycleStmt{g, append([]string{h}, t...)} }
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions/y.y:61
		{
			vals := yyDollar[1].ss
			yyVAL.cyclefn = func(h string) CycleStmt { return CycleStmt{Values: append([]string{h}, vals...)} }
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
//line expressions/y.y:68
		{
			yyVAL.ss = []string{}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:69
		{
			yyVAL.ss = append([]string{yyDollar[2].s}, yyDollar[3].ss...)
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions/y.y:72
		{
			yyVAL.exprs = append([]Expr{yyDollar[1].f}, yyDollar[2].exprs...)
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
//line expressions/y.y:74
		{
			yyVAL.exprs = []Expr{}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:75
		{
			yyVAL.exprs = append([]Expr{yyDollar[2].f}, yyDollar[3].exprs...)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions/y.y:78
		{
			s, ok := yyDollar[1].val.(string)
			if !ok {
				panic(SyntaxError(fmt.Sprintf("expected a string for %q", yyDollar[1].val)))
			}
			yyVAL.s = s
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
//line expressions/y.y:86
		{
			name, expr, mods := yyDollar[1].name, yyDollar[3].f, yyDollar[4].loopmods
			yyVAL.loop = LoopStmt{name, expr, mods}
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
//line expressions/y.y:92
		{
			yyVAL.loopmods = LoopModifiers{}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions/y.y:93
		{
			switch yyDollar[2].name {
			case "reversed":
				yyDollar[1].loopmods.Reversed = true
			default:
				panic(SyntaxError(fmt.Sprintf("undefined loop modifier %q", yyDollar[2].name)))
			}
			yyVAL.loopmods = yyDollar[1].loopmods
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:102
		{
			switch yyDollar[2].name {
			case "cols":
				yyDollar[1].loopmods.Cols = yyDollar[3].f
			case "limit":
				yyDollar[1].loopmods.Limit = yyDollar[3].f
			case "offset":
				yyDollar[1].loopmods.Offset = yyDollar[3].f
			default:
				panic(SyntaxError(fmt.Sprintf("undefined loop modifier %q", yyDollar[2].name)))
			}
			yyVAL.loopmods = yyDollar[1].loopmods
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions/y.y:118
		{
			val := yyDollar[1].val
			yyVAL.f = MakeLiteralExpr(val)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions/y.y:119
		{
			name := yyDollar[1].name
			yyVAL.f = MakeIdentExpr(name)
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions/y.y:120
		{
			yyVAL.f = MakePropertyExpr(yyDollar[1].f, yyDollar[2].name)
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line expressions/y.y:121
		{
			yyVAL.f = MakeIndexExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
//line expressions/y.y:122
		{
			yyVAL.f = MakeRangeExpr(yyDollar[2].f, yyDollar[4].f)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:123
		{
			yyVAL.f = yyDollar[2].f
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:128
		{
			yyVAL.f = MakeFilterExpr(yyDollar[1].f, yyDollar[3].name, nil)
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
//line expressions/y.y:129
		{
			yyVAL.f = MakeFilterExpr(yyDollar[1].f, yyDollar[3].name, yyDollar[4].filter_params)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions/y.y:133
		{
			yyVAL.filter_params = []LiteralExpr{MakeLiteralExpr(yyDollar[1].f)}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:135
		{
			yyVAL.filter_params = append(yyDollar[1].filter_params, MakeLiteralExpr(yyDollar[3].f))
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:139
		{
			//fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = MakeEQExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:146
		{
			//fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = MakeNEQExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:153
		{
			//fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = MakeGtExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:160
		{
			//fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = MakeLtExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:167
		{
			//fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = MakeGeExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:174
		{
			//fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = MakeLeExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:181
		{
			yyVAL.f = MakeContainsExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:186
		{
			yyVAL.f = MakeAndExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions/y.y:192
		{
			yyVAL.f = MakeOrExpr(yyDollar[1].f, yyDollar[3].f)
		}
	}
	goto yystack /* stack new state and value */
}
