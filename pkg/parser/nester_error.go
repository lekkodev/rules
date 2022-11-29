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
	"encoding/json"
	"fmt"
)

type ErrVals map[string]interface{}

func (e ErrVals) Dupe() ErrVals {
	ret := make(ErrVals)
	for k, v := range e {
		ret[k] = v
	}
	return ret

}
func (e ErrVals) Merge(vals ErrVals) {
	for k, v := range vals {
		e[k] = v
	}
}

type NestedError struct {
	Err  error
	Msg  string
	Vals ErrVals
}

func (e *NestedError) Original() error {
	switch val := e.Err.(type) {
	case *NestedError:
		return val.Original()
	default:
		return e.Err
	}
}

func (e *NestedError) Error() string {
	if e.Vals == nil {
		e.Vals = make(map[string]interface{})
	}
	e.Vals["err"] = e.Err.Error()
	e.Vals["msg"] = e.Msg
	data, err := json.Marshal(e.Vals.Dupe())
	if err != nil {
		return fmt.Sprintf("%s: %s", e.Msg, e.Err.Error())
	}
	return string(data)
}

func (e *NestedError) Set(vals ErrVals) *NestedError {
	if e.Vals == nil {
		e.Vals = make(ErrVals)
	}
	e.Vals.Merge(vals.Dupe())
	return e
}

func newNestedError(err error, msg string) *NestedError {
	return &NestedError{
		Err: err,
		Msg: msg,
	}
}
