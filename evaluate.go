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

package rules

import "github.com/lekkodev/rules/pkg/parser"

func Evaluate(rule string, items map[string]interface{}) (bool, error) {
	ev, err := parser.NewEvaluator(rule)
	if err != nil {
		return false, err
	}
	return ev.Process(items)
}
