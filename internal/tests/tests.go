/*
 * Flow CLI
 *
 * Copyright 2022 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tests

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/onflow/flow-cli/internal/command"
	"github.com/onflow/flow-cli/pkg/flowkit/util"

	testFramework "github.com/onflow/cadence/test-framework"
)

var Cmd = &cobra.Command{
	Use:              "test",
	Short:            "Utilities to run tests",
	TraverseChildren: true,
}

func init() {
	ExecuteCommand.AddToParent(Cmd)
}

var _ command.Result = &TestResult{}

type TestResult struct {
	testFramework.Results
}

func (r *TestResult) JSON() any {
	results := make([]map[string]string, 0, len(r.Results))

	for _, result := range r.Results {
		results = append(results, map[string]string{
			"testName": result.TestName,
			"error":    result.Error.Error(),
		})
	}

	return results
}

func (r *TestResult) String() string {
	var b bytes.Buffer
	writer := util.CreateTabWriter(&b)

	_, _ = fmt.Fprintf(writer, testFramework.PrettyPrintResults(r.Results))

	_ = writer.Flush()

	return b.String()
}

func (r *TestResult) Oneliner() string {
	return testFramework.PrettyPrintResults(r.Results)
}
