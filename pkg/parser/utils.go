// Copyright 2022 Lekko Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/pkg/errors"
)

func getString(s string) string {
	if len(s) > 2 {
		return s[1 : len(s)-1]
	}
	return ""
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

func lexAndParse(rule string) (antlr.ParseTree, error) {
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
	return tree, nil
}
