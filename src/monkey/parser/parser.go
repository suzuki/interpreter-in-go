package parser

import (
	"github.com/suzuki/interpreter-in-go/src/monkey/ast"
	"github.com/suzuki/interpreter-in-go/src/monkey/lexer"
	"github.com/suzuki/interpreter-in-go/src/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 最初に2回読むことで、 curToken, peekToken に値が入る
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
