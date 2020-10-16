// Code generated by goyacc -o querystring.y.go querystring.y. DO NOT EDIT.

//line querystring.y:2
package querystring

import __yyfmt__ "fmt"

//line querystring.y:2

//line querystring.y:5
type yySymType struct {
	yys int
	s   string
	n   int
	q   Condition
}

const tSTRING = 57346
const tPHRASE = 57347
const tNUMBER = 57348
const tAND = 57349
const tOR = 57350
const tNOT = 57351
const tTO = 57352
const tPLUS = 57353
const tMINUS = 57354
const tCOLON = 57355
const tLEFTBRACKET = 57356
const tRIGHTBRACKET = 57357
const tLEFTRANGE = 57358
const tRIGHTRANGE = 57359
const tGREATER = 57360
const tLESS = 57361
const tEQUAL = 57362

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tSTRING",
	"tPHRASE",
	"tNUMBER",
	"tAND",
	"tOR",
	"tNOT",
	"tTO",
	"tPLUS",
	"tMINUS",
	"tCOLON",
	"tLEFTBRACKET",
	"tRIGHTBRACKET",
	"tLEFTRANGE",
	"tRIGHTRANGE",
	"tGREATER",
	"tLESS",
	"tEQUAL",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 73

var yyAct = [...]int{

	26, 25, 27, 31, 38, 31, 5, 35, 31, 32,
	51, 32, 50, 30, 32, 28, 29, 19, 24, 37,
	45, 31, 34, 21, 22, 23, 47, 32, 46, 33,
	36, 39, 12, 14, 13, 42, 41, 11, 44, 6,
	7, 49, 10, 12, 14, 13, 1, 48, 11, 43,
	31, 40, 31, 10, 31, 8, 32, 4, 32, 3,
	32, 17, 18, 12, 14, 13, 9, 20, 2, 0,
	0, 16, 15,
}
var yyPact = [...]int{

	28, -1000, -1000, 28, 59, -1000, -1000, -1000, 54, -1000,
	39, 39, 10, -1000, -1000, -1000, -1000, 39, 39, 3,
	-1000, -3, -1000, -1000, -1000, -1000, -1000, -1000, 2, -1,
	46, -1000, 30, -1000, 44, -1000, -1000, 15, -1000, 18,
	16, -1000, -1000, -1000, -1000, -1000, 48, 36, -5, -7,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 0, 66, 68, 59, 6, 55, 57, 46,
}
var yyR1 = [...]int{

	0, 8, 3, 3, 4, 4, 5, 5, 5, 6,
	6, 6, 7, 7, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	1, 1,
}
var yyR2 = [...]int{

	0, 1, 2, 1, 2, 1, 1, 3, 3, 1,
	3, 2, 1, 1, 1, 1, 1, 3, 3, 3,
	4, 5, 4, 5, 4, 5, 4, 5, 7, 7,
	1, 2,
}
var yyChk = [...]int{

	-1000, -8, -3, -4, -7, -5, 11, 12, -6, -2,
	14, 9, 4, 6, 5, -3, -2, 7, 8, -5,
	-6, 13, -5, -5, 15, 4, -1, 5, 18, 19,
	16, 6, 12, -1, 20, 5, -1, 20, 5, -1,
	5, 6, -1, 5, -1, 5, 10, 10, -1, 5,
	17, 17,
}
var yyDef = [...]int{

	0, -2, 1, 3, 0, 5, 12, 13, 6, 9,
	0, 0, 14, 15, 16, 2, 4, 0, 0, 0,
	11, 0, 7, 8, 10, 17, 18, 19, 0, 0,
	0, 30, 0, 20, 0, 24, 22, 0, 26, 0,
	0, 31, 21, 25, 23, 27, 0, 0, 0, 0,
	28, 29,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20,
}
var yyTok3 = [...]int{
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
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
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
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
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
	yyn = yyPact[yystate]
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
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
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
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
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
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
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

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line querystring.y:26
		{
			yylex.(*lexerWrapper).query = yyDollar[1].q
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line querystring.y:31
		{
			yyVAL.q = NewAndCondition(yyDollar[1].q, yyDollar[2].q)
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line querystring.y:35
		{
			yyVAL.q = yyDollar[1].q
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line querystring.y:40
		{
			switch yyDollar[1].n {
			case queryMustNot:
				yyVAL.q = NewNotCondition(yyDollar[2].q)
			default:
				yyVAL.q = yyDollar[2].q
			}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line querystring.y:49
		{
			yyVAL.q = yyDollar[1].q
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line querystring.y:54
		{
			yyVAL.q = yyDollar[1].q
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line querystring.y:58
		{
			yyVAL.q = NewAndCondition(yyDollar[1].q, yyDollar[3].q)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line querystring.y:62
		{
			yyVAL.q = NewOrCondition(yyDollar[1].q, yyDollar[3].q)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line querystring.y:67
		{
			yyVAL.q = yyDollar[1].q
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line querystring.y:71
		{
			yyVAL.q = yyDollar[2].q
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line querystring.y:75
		{
			yyVAL.q = NewNotCondition(yyDollar[2].q)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line querystring.y:80
		{
			yyVAL.n = queryMust
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line querystring.y:84
		{
			yyVAL.n = queryMustNot
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line querystring.y:89
		{
			yyVAL.q = newStringCondition(yyDollar[1].s)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line querystring.y:93
		{
			yyVAL.q = NewMatchCondition(yyDollar[1].s)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line querystring.y:97
		{
			phrase := yyDollar[1].s
			q := NewMatchCondition(phrase)
			yyVAL.q = q
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line querystring.y:103
		{
			q := newStringCondition(yyDollar[3].s)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line querystring.y:109
		{
			val := yyDollar[3].s
			q := NewNumberRangeCondition(&val, &val, true, true)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
//line querystring.y:116
		{
			q := NewMatchCondition(yyDollar[3].s)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
//line querystring.y:122
		{
			val := yyDollar[4].s
			q := NewNumberRangeCondition(&val, nil, false, false)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
//line querystring.y:129
		{
			val := yyDollar[5].s
			q := NewNumberRangeCondition(&val, nil, true, false)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line querystring.y:136
		{
			val := yyDollar[4].s
			q := NewNumberRangeCondition(nil, &val, false, false)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
//line querystring.y:143
		{
			val := yyDollar[5].s
			q := NewNumberRangeCondition(nil, &val, false, true)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
//line querystring.y:150
		{
			phrase := yyDollar[4].s
			q := NewTimeRangeCondition(&phrase, nil, false, false)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 25:
		yyDollar = yyS[yypt-5 : yypt+1]
//line querystring.y:157
		{
			phrase := yyDollar[5].s
			q := NewTimeRangeCondition(&phrase, nil, true, false)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line querystring.y:164
		{
			phrase := yyDollar[4].s
			q := NewTimeRangeCondition(nil, &phrase, false, false)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
//line querystring.y:171
		{
			phrase := yyDollar[5].s
			q := NewTimeRangeCondition(nil, &phrase, false, true)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 28:
		yyDollar = yyS[yypt-7 : yypt+1]
//line querystring.y:178
		{
			min := yyDollar[4].s
			max := yyDollar[6].s
			q := NewNumberRangeCondition(&min, &max, true, true)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 29:
		yyDollar = yyS[yypt-7 : yypt+1]
//line querystring.y:186
		{
			min := yyDollar[4].s
			max := yyDollar[6].s
			q := NewTimeRangeCondition(&min, &max, true, true)
			q.SetField(yyDollar[1].s)
			yyVAL.q = q
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line querystring.y:195
		{
			yyVAL.s = yyDollar[1].s
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
//line querystring.y:199
		{
			yyVAL.s = "-" + yyDollar[2].s
		}
	}
	goto yystack /* stack new state and value */
}
