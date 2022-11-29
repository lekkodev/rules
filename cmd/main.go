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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/lekkodev/rules/pkg/parser"
)

func main() {
	jsonFile, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatalf("Error reading obj from file %v", err)
	}
	rulesBytes, err := ioutil.ReadFile("rules.txt")
	if err != nil {
		log.Fatalf("Error reading rule from file %v", err)
	}
	var rulesString = string(rulesBytes)
	var info map[string]interface{}
	if err := json.Unmarshal([]byte(jsonFile), &info); err != nil {
		log.Fatalf("Error unmarshalling from file %v", err)
	}
	ev, err := parser.NewEvaluator(rulesString)
	if err != nil {
		log.Fatalf("Error making evaluator from the rule %v, %v", rulesString, err)
	}
	ans, err := ev.Process(info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		`
		Rule: %v
		Object: %v
		Answer: %v
	\n`, rulesString, info, ans,
	)
}
