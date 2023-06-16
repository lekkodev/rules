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
	"strconv"
	"strings"

	rulesv1beta3 "buf.build/gen/go/lekkodev/cli/protocolbuffers/go/lekko/rules/v1beta3"
	"google.golang.org/protobuf/types/known/structpb"
)

const incorrectNumArgsErrorMsg = "incorrect number of args for function %s: got %v, expected %v"
const invalidTypeErrorMsg = "invalid type for arg %s for function %s: %v"
const outOfRangeErrorMsg = "value for arg %s for function %s out of range: %v"

var functionVisitors = map[string]func(*ASTBuilderV3, *CallExpContext) interface{}{
	// Add new function name -> visitor functions here, implementing them below
	"bucket": (*ASTBuilderV3).visitBucket,
}

func (a *ASTBuilderV3) visitBucket(ctx *CallExpContext) interface{} {
	funcArgs := ctx.AllFunctionArg()

	if len(funcArgs) != 2 {
		return fmt.Errorf(incorrectNumArgsErrorMsg, "bucket", len(funcArgs), 2)
	}

	contextKey := funcArgs[0].Accept(a)
	if err, ok := contextKey.(error); ok {
		return err
	}
	contextKeyStr, ok := contextKey.(string)
	if !ok {
		return fmt.Errorf(invalidTypeErrorMsg, "context_key", "bucket", contextKey)
	}

	threshold := funcArgs[1].Accept(a)
	if err, ok := threshold.(error); ok {
		return err
	}
	thresholdValue, ok := threshold.(*structpb.Value)
	if !ok {
		return fmt.Errorf(invalidTypeErrorMsg, "threshold", "bucket", threshold)
	}
	if _, ok := thresholdValue.GetKind().(*structpb.Value_NumberValue); !ok {
		return fmt.Errorf(invalidTypeErrorMsg, "threshold", "bucket", threshold)
	}
	if thresholdValue.GetNumberValue() < 0 || thresholdValue.GetNumberValue() > 100 {
		return fmt.Errorf(outOfRangeErrorMsg, "threshold", "bucket", thresholdValue)
	}
	// The client-side interface is floats in [0, 100] but we store ints in [0, 100000]
	// Need to store converted value while avoiding precision issues
	// Stop at 3 decimal places and convert to int, "multiplying" by 1000
	// FormatFloat adds/rounds decimal places in the string representation
	thresholdStr := strconv.FormatFloat(thresholdValue.GetNumberValue(), 'f', 3, 64)
	thresholdNum, err := strconv.Atoi(strings.Replace(thresholdStr, ".", "", 1))
	if err != nil {
		return err
	}

	return &rulesv1beta3.Rule{
		Rule: &rulesv1beta3.Rule_CallExpression{
			CallExpression: &rulesv1beta3.CallExpression{
				Function: &rulesv1beta3.CallExpression_Bucket_{
					Bucket: &rulesv1beta3.CallExpression_Bucket{
						ContextKey: contextKeyStr,
						Threshold:  uint32(thresholdNum),
					},
				},
			},
		},
	}
}
