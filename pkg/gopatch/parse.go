package gopatch

import (
	"bytes"
	"fmt"
	"go/parser"
	"io"
	"reflect"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"honnef.co/go/tools/pattern"
)

type GoPatch struct {
	Patches []*Patch `@@*`
}

type Patch struct {
	MetaVariables []MetaVariable `("@" "@" ("var" @@)* "@" "@")?`
	Diff          Diff           `@@`
}

type MetaVariable struct {
	Name string `@Ident`
	Type string `"expression" | "identifier"`
}

type DiffType int

func (d *DiffType) Capture(values []string) error {
	if len(values) == 0 {
		*d = 0
	} else if values[0] == "-" {
		*d = -1
	} else {
		*d = 1
	}
	return nil
}

type Package struct {
	X    string `@("+" | "-")*`
	Name string `@Ident`
}

type Imports struct {
	Imports []Line `@@+`
}

type Diff struct {
	Package *string  `(('-' | '+')? 'package' @Ident)?`
	Imports []string `(('-' | '+')? 'import' @String )?`
	Expr    []Line   `@@+`
}

func (d *Diff) Before() pattern.Pattern {
	var b strings.Builder
	for _, l := range d.Expr {
		if l.Type == 1 {
			continue
		}
		b.WriteString(l.Data)
		b.WriteRune('\n')
	}
	expr, _ := parser.ParseExpr(b.String())
	root := convert(expr)

	relevant := map[reflect.Type]struct{}{}
	pattern.Roots(root, relevant)
	return pattern.Pattern{
		Root:     root,
		Relevant: relevant,
		Bindings: []string{"e"},
	}
}

func (d *Diff) After() pattern.Pattern {
	var b strings.Builder
	for _, l := range d.Expr {
		if l.Type == -1 {
			continue
		}
		b.WriteString(l.Data)
		b.WriteRune('\n')
	}
	expr, _ := parser.ParseExpr(b.String())
	root := convert(expr)

	relevant := map[reflect.Type]struct{}{}
	pattern.Roots(root, relevant)
	return pattern.Pattern{
		Root:     root,
		Relevant: relevant,
		Bindings: []string{"e"},
	}
}

type Line struct {
	Type DiffType
	Data string
}

func (l *Line) Parse(lex *lexer.PeekingLexer) error {
	peek := lex.Peek()
	lineNo := peek.Pos.Line

	switch peek.Value {
	case "-":
		l.Type = -1
		_ = lex.Next()
	case "+":
		l.Type = 1
		_ = lex.Next()
	default:
	}

	for lex.Peek().Pos.Line == lineNo {
		tok := lex.Next()
		if tok.EOF() {
			return participle.NextMatch
		}
		l.Data += tok.Value
	}
	return nil
}

func Parse(in io.Reader) (*GoPatch, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}

	parser := participle.MustBuild[GoPatch]()

	tokens, err := parser.Lex("", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	fmt.Printf("Tokens: %#v\n", tokens)

	return parser.Parse("", bytes.NewReader(data))
}
