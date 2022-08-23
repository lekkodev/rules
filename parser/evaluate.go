package parser

import (
	"fmt"
	"runtime/debug"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pkg/errors"
)

type Evaluator struct {
	rule string
	tree antlr.ParseTree

	testHookPanic func()
}

type errorListener struct {
	antlr.DefaultErrorListener
	err error
}

func (e *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, err antlr.RecognitionException) {
	newErr := fmt.Errorf("syntax error when parsing symbol: %v at %d:%d: %s %v", offendingSymbol, line, column, msg, err)
	if e.err != nil {
		e.err = errors.Wrap(newErr, e.err.Error())
	} else {
		e.err = newErr
	}
}

func (e *errorListener) Error() error {
	return e.err
}

func NewEvaluator(rule string) (ret *Evaluator, retErr error) {
	// antlr lib has panics for exceptions so we have to put a recover here
	// in the unlikely case there is an exception
	defer func() {
		info := recover()
		if info != nil {
			retErr = fmt.Errorf("%q\nstack:\n %v", info, string(debug.Stack()))
		}
	}()
	input := antlr.NewInputStream(rule)
	errListener := new(errorListener)
	lex := NewJsonQueryLexer(input)
	// Remove default log listener and add our own.
	lex.RemoveErrorListeners()
	lex.AddErrorListener(errListener)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	if err := errListener.Error(); err != nil {
		return nil, err
	}
	p := NewJsonQueryParser(tokens)
	// Remove default log listener and add our own.
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)
	tree := p.Query()
	if err := errListener.Error(); err != nil {
		return nil, err
	}
	return &Evaluator{
		rule: rule,
		tree: tree,
	}, nil
}

func (e *Evaluator) Process(items map[string]interface{}) (ret bool, retErr error) {
	// antlr lib has panics for exceptions so we have to put a recover here
	// in the unlikely case there is an exception
	defer func() {
		info := recover()
		if info != nil {
			retErr = fmt.Errorf("%q\nstack:\n %v", info, string(debug.Stack()))
			ret = false
		}
	}()

	visitor := NewJsonQueryVisitorImpl(items)
	result := visitor.Visit(e.tree)
	if e.testHookPanic != nil {
		defer e.testHookPanic()
	}
	if result == nil || visitor.err != nil {
		return false, visitor.err
	}

	return result.(bool), visitor.err
}

func Evaluate(rule string, items map[string]interface{}) bool {
	ev, err := NewEvaluator(rule)
	if err != nil {
		return false
	}
	result, _ := ev.Process(items)
	return result
}
