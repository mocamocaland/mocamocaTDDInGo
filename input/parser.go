package input

import (
	"github.com/mocamocaland/tdd-in-go-sample/calculator"
)

type Parser struct {
	engine *calculator.Engine
	//validator *Validator
}

func (p *Parser) ProcessExpression(expr string) (*string, error) {
	// implementation code
	expr = "test"
	return &expr, nil
}
