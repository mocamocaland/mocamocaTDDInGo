package input

import (
	"github.com/mocamocaland/mocamocaTDDInGo/calculator"
)

type Parser struct {
	engine    *calculator.Engine
	validator *Validator
}

func (p *Parser) ProcessExpression(expr string) (*string, error) {
	// implementation code
}
