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

/*
This file contains logic implementing the Configurator interface for multi-geth.

Notes:
When setting the difficulty bomb delay map using a wanted total difficulty
value. The map, following Parity's format, uses aggregating values (summed) to yield a net difficulty delay,
while the specs use a max value (eg 3m, 5m, 9m). The model's difficulty bomb delay map data type has a
method SetValueForHeight which is used for this.
*/

package chipprgeth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/params/types/ctypes"
	"github.com/ethereum/go-ethereum/params/types/internal"
	"github.com/ethereum/go-ethereum/params/vars"
)

func newU64(u uint64) *uint64 {
	return &u
}

func bigNewU64(i *big.Int) *uint64 {
	if i == nil {
		return nil
	}
	return newU64(i.Uint64())
}

func setBig(i *big.Int, u *uint64) *big.Int {
	if u == nil {
		return nil
	}
	i = big.NewInt(int64(*u))
	return i
}

func (c *ChipprGethChainConfig) ensureExistingRewardSchedule() {
	if c.BlockRewardSchedule == nil {
		c.BlockRewardSchedule = ctypes.Uint64BigMapEncodesHex{}
	}
}

func (c *ChipprGethChainConfig) ensureExistingDifficultySchedule() {
	if c.DifficultyBombDelaySchedule == nil {
		c.DifficultyBombDelaySchedule = ctypes.Uint64BigMapEncodesHex{}
	}
}

func (c *ChipprGethChainConfig) GetAccountStartNonce() *uint64 {
	return internal.GlobalConfigurator().GetAccountStartNonce()
}
func (c *ChipprGethChainConfig) SetAccountStartNonce(n *uint64) error {
	return internal.GlobalConfigurator().SetAccountStartNonce(n)
}
func (c *ChipprGethChainConfig) GetMaximumExtraDataSize() *uint64 {
	return internal.GlobalConfigurator().GetMaximumExtraDataSize()
}
func (c *ChipprGethChainConfig) SetMaximumExtraDataSize(n *uint64) error {
	return internal.GlobalConfigurator().SetMaximumExtraDataSize(n)
}
func (c *ChipprGethChainConfig) GetMinGasLimit() *uint64 {
	return internal.GlobalConfigurator().GetMinGasLimit()
}
func (c *ChipprGethChainConfig) SetMinGasLimit(n *uint64) error {
	return internal.GlobalConfigurator().SetMinGasLimit(n)
}
func (c *ChipprGethChainConfig) GetGasLimitBoundDivisor() *uint64 {
	return internal.GlobalConfigurator().GetGasLimitBoundDivisor()
}
func (c *ChipprGethChainConfig) SetGasLimitBoundDivisor(n *uint64) error {
	return internal.GlobalConfigurator().SetGasLimitBoundDivisor(n)
}

func (c *ChipprGethChainConfig) GetNetworkID() *uint64 {
	return newU64(c.NetworkID)
}

func (c *ChipprGethChainConfig) SetNetworkID(n *uint64) error {
	if n == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	c.NetworkID = *n
	return nil
}

func (c *ChipprGethChainConfig) GetChainID() *big.Int {
	return c.ChainID
}

func (c *ChipprGethChainConfig) SetChainID(n *big.Int) error {
	c.ChainID = n
	return nil
}

func (c *ChipprGethChainConfig) GetMaxCodeSize() *uint64 {
	return internal.GlobalConfigurator().GetMaxCodeSize()
}
func (c *ChipprGethChainConfig) SetMaxCodeSize(n *uint64) error {
	return internal.GlobalConfigurator().SetMaxCodeSize(n)
}

func (c *ChipprGethChainConfig) GetEIP7Transition() *uint64 {
	return bigNewU64(c.EIP7FBlock)
}

func (c *ChipprGethChainConfig) SetEIP7Transition(n *uint64) error {
	c.EIP7FBlock = setBig(c.EIP7FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP150Transition() *uint64 {
	return bigNewU64(c.EIP150Block)
}

func (c *ChipprGethChainConfig) SetEIP150Transition(n *uint64) error {
	c.EIP150Block = setBig(c.EIP150Block, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP152Transition() *uint64 {
	return bigNewU64(c.EIP152FBlock)
}

func (c *ChipprGethChainConfig) SetEIP152Transition(n *uint64) error {
	c.EIP152FBlock = setBig(c.EIP152FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP160Transition() *uint64 {
	return bigNewU64(c.EIP160FBlock)
}

func (c *ChipprGethChainConfig) SetEIP160Transition(n *uint64) error {
	c.EIP160FBlock = setBig(c.EIP160FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP161dTransition() *uint64 {
	return bigNewU64(c.EIP161FBlock)
}

func (c *ChipprGethChainConfig) SetEIP161dTransition(n *uint64) error {
	c.EIP161FBlock = setBig(c.EIP161FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP161abcTransition() *uint64 {
	return bigNewU64(c.EIP161FBlock)
}

func (c *ChipprGethChainConfig) SetEIP161abcTransition(n *uint64) error {
	c.EIP161FBlock = setBig(c.EIP161FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP170Transition() *uint64 {
	return bigNewU64(c.EIP170FBlock)
}

func (c *ChipprGethChainConfig) SetEIP170Transition(n *uint64) error {
	c.EIP170FBlock = setBig(c.EIP170FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP155Transition() *uint64 {
	return bigNewU64(c.EIP155Block)
}

func (c *ChipprGethChainConfig) SetEIP155Transition(n *uint64) error {
	c.EIP155Block = setBig(c.EIP155Block, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP140Transition() *uint64 {
	return bigNewU64(c.EIP140FBlock)
}

func (c *ChipprGethChainConfig) SetEIP140Transition(n *uint64) error {
	c.EIP140FBlock = setBig(c.EIP140FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP198Transition() *uint64 {
	return bigNewU64(c.EIP198FBlock)
}

func (c *ChipprGethChainConfig) SetEIP198Transition(n *uint64) error {
	c.EIP198FBlock = setBig(c.EIP198FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP211Transition() *uint64 {
	return bigNewU64(c.EIP211FBlock)
}

func (c *ChipprGethChainConfig) SetEIP211Transition(n *uint64) error {
	c.EIP211FBlock = setBig(c.EIP211FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP212Transition() *uint64 {
	return bigNewU64(c.EIP212FBlock)
}

func (c *ChipprGethChainConfig) SetEIP212Transition(n *uint64) error {
	c.EIP212FBlock = setBig(c.EIP212FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP213Transition() *uint64 {
	return bigNewU64(c.EIP213FBlock)
}

func (c *ChipprGethChainConfig) SetEIP213Transition(n *uint64) error {
	c.EIP213FBlock = setBig(c.EIP213FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP214Transition() *uint64 {
	return bigNewU64(c.EIP214FBlock)
}

func (c *ChipprGethChainConfig) SetEIP214Transition(n *uint64) error {
	c.EIP214FBlock = setBig(c.EIP214FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP658Transition() *uint64 {
	return bigNewU64(c.EIP658FBlock)
}

func (c *ChipprGethChainConfig) SetEIP658Transition(n *uint64) error {
	c.EIP658FBlock = setBig(c.EIP658FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP145Transition() *uint64 {
	return bigNewU64(c.EIP145FBlock)
}

func (c *ChipprGethChainConfig) SetEIP145Transition(n *uint64) error {
	c.EIP145FBlock = setBig(c.EIP145FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP1014Transition() *uint64 {
	return bigNewU64(c.EIP1014FBlock)
}

func (c *ChipprGethChainConfig) SetEIP1014Transition(n *uint64) error {
	c.EIP1014FBlock = setBig(c.EIP1014FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP1052Transition() *uint64 {
	return bigNewU64(c.EIP1052FBlock)
}

func (c *ChipprGethChainConfig) SetEIP1052Transition(n *uint64) error {
	c.EIP1052FBlock = setBig(c.EIP1052FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP1283Transition() *uint64 {
	return bigNewU64(c.EIP1283FBlock)
}

func (c *ChipprGethChainConfig) SetEIP1283Transition(n *uint64) error {
	c.EIP1283FBlock = setBig(c.EIP1283FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP1283DisableTransition() *uint64 {
	return bigNewU64(c.PetersburgBlock)
}

func (c *ChipprGethChainConfig) SetEIP1283DisableTransition(n *uint64) error {
	c.PetersburgBlock = setBig(c.PetersburgBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP1108Transition() *uint64 {
	return bigNewU64(c.EIP1108FBlock)
}

func (c *ChipprGethChainConfig) SetEIP1108Transition(n *uint64) error {
	c.EIP1108FBlock = setBig(c.EIP1108FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP2200Transition() *uint64 {
	return bigNewU64(c.EIP2200FBlock)
}

func (c *ChipprGethChainConfig) SetEIP2200Transition(n *uint64) error {
	c.EIP2200FBlock = setBig(c.EIP2200FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP2200DisableTransition() *uint64 {
	return bigNewU64(c.EIP2200DisableFBlock)
}

func (c *ChipprGethChainConfig) SetEIP2200DisableTransition(n *uint64) error {
	c.EIP2200DisableFBlock = setBig(c.EIP2200DisableFBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP1344Transition() *uint64 {
	return bigNewU64(c.EIP1344FBlock)
}

func (c *ChipprGethChainConfig) SetEIP1344Transition(n *uint64) error {
	c.EIP1344FBlock = setBig(c.EIP1344FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP1884Transition() *uint64 {
	return bigNewU64(c.EIP1884FBlock)
}

func (c *ChipprGethChainConfig) SetEIP1884Transition(n *uint64) error {
	c.EIP1884FBlock = setBig(c.EIP1884FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP2028Transition() *uint64 {
	return bigNewU64(c.EIP2028FBlock)
}

func (c *ChipprGethChainConfig) SetEIP2028Transition(n *uint64) error {
	c.EIP2028FBlock = setBig(c.EIP2028FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetECIP1080Transition() *uint64 {
	return bigNewU64(c.ECIP1080FBlock)
}

func (c *ChipprGethChainConfig) SetECIP1080Transition(n *uint64) error {
	c.ECIP1080FBlock = setBig(c.ECIP1080FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP1706Transition() *uint64 {
	return bigNewU64(c.EIP1706FBlock)
}

func (c *ChipprGethChainConfig) SetEIP1706Transition(n *uint64) error {
	c.EIP1706FBlock = setBig(c.EIP1706FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP2537Transition() *uint64 {
	return bigNewU64(c.EIP2537FBlock)
}

func (c *ChipprGethChainConfig) SetEIP2537Transition(n *uint64) error {
	c.EIP2537FBlock = setBig(c.EIP2537FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) IsEnabled(fn func() *uint64, n *big.Int) bool {
	f := fn()
	if f == nil || n == nil {
		return false
	}
	return big.NewInt(int64(*f)).Cmp(n) <= 0
}

func (c *ChipprGethChainConfig) GetForkCanonHash(n uint64) common.Hash {
	if c.RequireBlockHashes == nil {
		return common.Hash{}
	}
	for k, v := range c.RequireBlockHashes {
		if k == n {
			return v
		}
	}
	return common.Hash{}
}

func (c *ChipprGethChainConfig) SetForkCanonHash(n uint64, h common.Hash) error {
	if c.RequireBlockHashes == nil {
		c.RequireBlockHashes = make(map[uint64]common.Hash)
	}
	c.RequireBlockHashes[n] = h
	return nil
}

func (c *ChipprGethChainConfig) GetForkCanonHashes() map[uint64]common.Hash {
	return c.RequireBlockHashes
}

func (c *ChipprGethChainConfig) GetConsensusEngineType() ctypes.ConsensusEngineT {
	if c.Ethash != nil {
		return ctypes.ConsensusEngineT_Ethash
	}
	if c.Clique != nil {
		return ctypes.ConsensusEngineT_Clique
	}
	return ctypes.ConsensusEngineT_Unknown
}

func (c *ChipprGethChainConfig) MustSetConsensusEngineType(t ctypes.ConsensusEngineT) error {
	switch t {
	case ctypes.ConsensusEngineT_Ethash:
		c.Ethash = new(ctypes.EthashConfig)
		c.Clique = nil
		return nil
	case ctypes.ConsensusEngineT_Clique:
		c.Clique = new(ctypes.CliqueConfig)
		c.Ethash = nil
		return nil
	default:
		return ctypes.ErrUnsupportedConfigFatal
	}
}

func (c *ChipprGethChainConfig) GetEthashMinimumDifficulty() *big.Int {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	return internal.GlobalConfigurator().GetEthashMinimumDifficulty()
}
func (c *ChipprGethChainConfig) SetEthashMinimumDifficulty(i *big.Int) error {
	return internal.GlobalConfigurator().SetEthashMinimumDifficulty(i)
}

func (c *ChipprGethChainConfig) GetEthashDifficultyBoundDivisor() *big.Int {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	return internal.GlobalConfigurator().GetEthashDifficultyBoundDivisor()
}

func (c *ChipprGethChainConfig) SetEthashDifficultyBoundDivisor(i *big.Int) error {
	return internal.GlobalConfigurator().SetEthashDifficultyBoundDivisor(i)
}

func (c *ChipprGethChainConfig) GetEthashDurationLimit() *big.Int {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	return internal.GlobalConfigurator().GetEthashDurationLimit()
}

func (c *ChipprGethChainConfig) SetEthashDurationLimit(i *big.Int) error {
	return internal.GlobalConfigurator().SetEthashDurationLimit(i)
}

func (c *ChipprGethChainConfig) GetEthashHomesteadTransition() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	if c.EIP2FBlock == nil || c.EIP7FBlock == nil {
		return nil
	}
	return bigNewU64(math.BigMax(c.EIP2FBlock, c.EIP7FBlock))
}

func (c *ChipprGethChainConfig) SetEthashHomesteadTransition(n *uint64) error {
	c.EIP2FBlock = setBig(c.EIP2FBlock, n)
	c.EIP7FBlock = setBig(c.EIP7FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEIP2Transition() *uint64 {
	return bigNewU64(c.EIP2FBlock)
}

func (c *ChipprGethChainConfig) SetEIP2Transition(n *uint64) error {
	c.EIP2FBlock = setBig(c.EIP2FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEthashEIP779Transition() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	return bigNewU64(c.DAOForkBlock)
}

func (c *ChipprGethChainConfig) SetEthashEIP779Transition(n *uint64) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	c.DAOForkBlock = setBig(c.DAOForkBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEthashEIP649Transition() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	if c.eip649FInferred {
		return bigNewU64(c.EIP649FBlock)
	}

	var diffN *uint64
	defer func() {
		c.EIP649FBlock = setBig(c.EIP649FBlock, diffN)
		c.eip649FInferred = true
	}()

	// Get block number (key) from maps where EIP649 criteria is met.
	diffN = ctypes.MapMeetsSpecification(
		c.DifficultyBombDelaySchedule,
		c.BlockRewardSchedule,
		vars.EIP649DifficultyBombDelay,
		vars.EIP649FBlockReward,
	)
	return diffN
}

func (c *ChipprGethChainConfig) SetEthashEIP649Transition(n *uint64) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}

	c.EIP649FBlock = setBig(c.EIP649FBlock, n)
	c.eip649FInferred = true

	if n == nil {
		return nil
	}

	c.ensureExistingRewardSchedule()
	c.BlockRewardSchedule[*n] = vars.EIP649FBlockReward

	c.ensureExistingDifficultySchedule()
	c.DifficultyBombDelaySchedule.SetValueTotalForHeight(n, vars.EIP649DifficultyBombDelay)

	return nil
}

func (c *ChipprGethChainConfig) GetEthashEIP1234Transition() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	if c.eip1234FInferred {
		return bigNewU64(c.EIP1234FBlock)
	}

	var diffN *uint64
	defer func() {
		c.EIP1234FBlock = setBig(c.EIP1234FBlock, diffN)
		c.eip1234FInferred = true
	}()

	// Get block number (key) from maps where EIP1234 criteria is met.
	diffN = ctypes.MapMeetsSpecification(
		c.DifficultyBombDelaySchedule,
		c.BlockRewardSchedule,
		vars.EIP1234DifficultyBombDelay,
		vars.EIP1234FBlockReward,
	)
	return diffN
}

func (c *ChipprGethChainConfig) SetEthashEIP1234Transition(n *uint64) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}

	c.EIP1234FBlock = setBig(c.EIP1234FBlock, n)
	c.eip1234FInferred = true

	if n == nil {
		return nil
	}

	// Block reward is a simple lookup; doesn't matter if overwrite or not.
	c.ensureExistingRewardSchedule()
	c.BlockRewardSchedule[*n] = vars.EIP1234FBlockReward

	c.ensureExistingDifficultySchedule()
	c.DifficultyBombDelaySchedule.SetValueTotalForHeight(n, vars.EIP1234DifficultyBombDelay)

	return nil
}

func (c *ChipprGethChainConfig) GetEthashEIP2384Transition() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	if c.eip2384Inferred {
		return bigNewU64(c.EIP2384FBlock)
	}

	var diffN *uint64
	defer func() {
		c.EIP2384FBlock = setBig(c.EIP2384FBlock, diffN)
		c.eip2384Inferred = true
	}()

	// Get block number (key) from map where EIP2384 criteria is met.
	diffN = ctypes.MapMeetsSpecification(c.DifficultyBombDelaySchedule, nil, vars.EIP2384DifficultyBombDelay, nil)
	return diffN
}

func (c *ChipprGethChainConfig) SetEthashEIP2384Transition(n *uint64) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}

	c.EIP2384FBlock = setBig(c.EIP2384FBlock, n)
	c.eip2384Inferred = true

	if n == nil {
		return nil
	}

	c.ensureExistingDifficultySchedule()
	c.DifficultyBombDelaySchedule.SetValueTotalForHeight(n, vars.EIP2384DifficultyBombDelay)

	return nil
}

func (c *ChipprGethChainConfig) GetEthashECIP1010PauseTransition() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	return bigNewU64(c.ECIP1010PauseBlock)
}

func (c *ChipprGethChainConfig) SetEthashECIP1010PauseTransition(n *uint64) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	if c.ECIP1010PauseBlock == nil && c.ECIP1010Length != nil {
		c.ECIP1010PauseBlock = setBig(c.ECIP1010PauseBlock, n)
		c.ECIP1010Length = c.ECIP1010Length.Sub(c.ECIP1010Length, c.ECIP1010PauseBlock)
		return nil
	}
	c.ECIP1010PauseBlock = setBig(c.ECIP1010PauseBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEthashECIP1010ContinueTransition() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	if c.ECIP1010PauseBlock == nil {
		return nil
	}
	if c.ECIP1010Length == nil {
		return nil
	}
	// transition = pause + length
	return bigNewU64(new(big.Int).Add(c.ECIP1010PauseBlock, c.ECIP1010Length))
}

func (c *ChipprGethChainConfig) SetEthashECIP1010ContinueTransition(n *uint64) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	// length = continue - pause
	if n == nil {
		return ctypes.ErrUnsupportedConfigNoop
	}
	if c.ECIP1010PauseBlock == nil {
		c.ECIP1010Length = new(big.Int).SetUint64(*n)
		return nil
	}
	c.ECIP1010Length = new(big.Int).Sub(big.NewInt(int64(*n)), c.ECIP1010PauseBlock)
	return nil
}

func (c *ChipprGethChainConfig) GetEthashECIP1017Transition() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	return bigNewU64(c.ECIP1017FBlock)
}

func (c *ChipprGethChainConfig) SetEthashECIP1017Transition(n *uint64) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	c.ECIP1017FBlock = setBig(c.ECIP1017FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEthashECIP1017EraRounds() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	return bigNewU64(c.ECIP1017EraRounds)
}

func (c *ChipprGethChainConfig) SetEthashECIP1017EraRounds(n *uint64) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	c.ECIP1017EraRounds = setBig(c.ECIP1017EraRounds, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEthashEIP100BTransition() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	return bigNewU64(c.EIP100FBlock)
}

func (c *ChipprGethChainConfig) SetEthashEIP100BTransition(n *uint64) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	c.EIP100FBlock = setBig(c.EIP100FBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEthashECIP1041Transition() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	return bigNewU64(c.DisposalBlock)
}

func (c *ChipprGethChainConfig) SetEthashECIP1041Transition(n *uint64) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	c.DisposalBlock = setBig(c.DisposalBlock, n)
	return nil
}

func (c *ChipprGethChainConfig) GetECIP1092Transition() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}

	return bigNewU64(c.ECIP1092Transition)
}

func (c *ChipprGethChainConfig) SetECIP1092Transition(n *uint64) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	c.ECIP1092Transition = setBig(c.ECIP1092Transition, n)
	return nil
}

func (c *ChipprGethChainConfig) GetPenatlyCheckLength() *uint64 {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}

	return bigNewU64(c.PenatlyCheckLength)
}

func (c *ChipprGethChainConfig) SetPenatlyCheckLength(n *uint64) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	c.PenatlyCheckLength = setBig(c.PenatlyCheckLength, n)
	return nil
}

func (c *ChipprGethChainConfig) GetEthashDifficultyBombDelaySchedule() ctypes.Uint64BigMapEncodesHex {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	return c.DifficultyBombDelaySchedule
}

func (c *ChipprGethChainConfig) SetEthashDifficultyBombDelaySchedule(m ctypes.Uint64BigMapEncodesHex) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	c.DifficultyBombDelaySchedule = m
	return nil
}

func (c *ChipprGethChainConfig) GetEthashBlockRewardSchedule() ctypes.Uint64BigMapEncodesHex {
	if c.GetConsensusEngineType() != ctypes.ConsensusEngineT_Ethash {
		return nil
	}
	return c.BlockRewardSchedule
}

func (c *ChipprGethChainConfig) SetEthashBlockRewardSchedule(m ctypes.Uint64BigMapEncodesHex) error {
	if c.Ethash == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	c.BlockRewardSchedule = m
	return nil
}

func (c *ChipprGethChainConfig) GetCliquePeriod() uint64 {
	if c.Clique == nil {
		return 0
	}
	return c.Clique.Period
}

func (c *ChipprGethChainConfig) SetCliquePeriod(n uint64) error {
	if c.Clique == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	c.Clique.Period = n
	return nil
}

func (c *ChipprGethChainConfig) GetCliqueEpoch() uint64 {
	if c.Clique == nil {
		return 0
	}
	return c.Clique.Epoch
}

func (c *ChipprGethChainConfig) SetCliqueEpoch(n uint64) error {
	if c.Clique == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	c.Clique.Epoch = n
	return nil
}
