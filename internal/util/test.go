/*
 * Flow CLI
 *
 * Copyright 2019 Dapper Labs, Inc.
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

package util

import (
	"testing"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-cli/flowkit"
	"github.com/onflow/flow-cli/flowkit/mocks"
	"github.com/onflow/flow-cli/flowkit/output"
	"github.com/onflow/flow-cli/flowkit/tests"
)

var NoLogger = output.NewStdoutLogger(output.NoneLog)

var TestID = flow.HexToID("24993fc99f81641c45c0afa307e683b4f08d407d90041aa9439f487acb33d633")

// TestMocks creates mock flowkit services, an empty state and a mock reader writer
func TestMocks(t *testing.T) (*mocks.MockServices, *flowkit.State, flowkit.ReaderWriter) {
	services := mocks.DefaultMockServices()
	rw, _ := tests.ReaderWriter()
	state, err := flowkit.Init(rw, crypto.ECDSA_P256, crypto.SHA3_256)
	require.NoError(t, err)

	return services, state, rw
}
