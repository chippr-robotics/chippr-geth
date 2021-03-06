// Copyright 2019 The multi-geth Authors
// This file is part of the multi-geth library.
//
// The multi-geth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The multi-geth library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the multi-geth library. If not, see <http://www.gnu.org/licenses/>.

package tests

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/internal/build"
)

// This file holds variable and type relating specifically
// to the task of generating tests.

var (
	CG_GENERATE_STATE_TESTS_KEY      = "CHIPPRGETH_TESTS_GENERATE_STATE_TESTS"
	CG_GENERATE_DIFFICULTY_TESTS_KEY = "CHIPPRGETH_TESTS_GENERATE_DIFFICULTY_TESTS"

	// Feature Equivalence tests use convert.Convert to
	// run tests using alternating ChainConfig data type implementations.
	CG_CHAINCONFIG_FEATURE_EQ_CHIPPRGETH_KEY   = "CHIPPRGETH_TESTS_CHAINCONFIG_FEATURE_EQUIVALENCE_CHIPPRGETH"
	CG_CHAINCONFIG_FEATURE_EQ_MULTIGETHV0_KEY  = "CHIPPRGETH_TESTS_CHAINCONFIG_FEATURE_EQUIVALENCE_MULTIGETHV0"
	CG_CHAINCONFIG_FEATURE_EQ_OPENETHEREUM_KEY = "CHIPPRGETH_TESTS_CHAINCONFIG_FEATURE_EQUIVALENCE_OPENETHEREUM"

	CG_CHAINCONFIG_CONSENSUS_EQ_CLIQUE = "CHIPPRGETH_TESTS_CHAINCONFIG_CONSENSUS_EQUIVALENCE_CLIQUE"

	// Parity specs tests use Parity JSON config data (in params/parity.json.d/)
	// when applicable as equivalent config implementations for the default Go data type
	// configs.
	CG_CHAINCONFIG_CHAINSPECS_OPENETHEREUM_KEY = "CHIPPRGETH_TESTS_CHAINCONFIG_OPENETHEREUM_SPECS"
)

type chainspecRefsT map[string]chainspecRef

var chainspecRefsState = chainspecRefsT{}
var chainspecRefsDifficulty = chainspecRefsT{}

type chainspecRef struct {
	Filename string `json:"filename"`
	Sha1Sum  []byte `json:"sha1sum"`
}

func (c chainspecRef) String() string {
	return fmt.Sprintf("file: %s, file.sha1sum: %x", c.Filename, c.Sha1Sum)
}

func (c *chainspecRef) UnmarshalJSON(input []byte) error {
	type xT struct {
		F string `json:"filename"`
		S string `json:"sha1sum"`
	}
	var x = xT{}
	err := json.Unmarshal(input, &x)
	if err != nil {
		return err
	}
	c.Filename = x.F
	c.Sha1Sum = common.Hex2Bytes(x.S)
	return nil
}

func (c chainspecRef) MarshalJSON() ([]byte, error) {
	var x = struct {
		F string `json:"filename"`
		S string `json:"sha1sum"`
	}{
		F: c.Filename,
		S: common.Bytes2Hex(c.Sha1Sum[:]),
	}

	return json.MarshalIndent(x, "", "    ")
}

// submoduleParentRef captures the current git status of the tests submodule.
// This is used for reference when writing tests.
var submoduleParentRef = func() string {
	subModOut := build.RunGit("submodule", "status")
	subModOut = strings.ReplaceAll(strings.TrimSpace(subModOut), " ", "_")
	return subModOut
}()
