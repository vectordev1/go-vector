// Copyright 2014 The go-vector Authors
// This file is part of the go-vector library.
//
// The go-vector library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-vector library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-vector library. If not, see <http://www.gnu.org/licenses/>.

package pow

import (
	"math/big"

	"github.com/vector/go-vector/common"
	"github.com/vector/go-vector/core/types"
)

type Block interface {
	Difficulty() *big.Int
	HashNoNonce() common.Hash
	Nonce() uint64
	MixDigest() common.Hash
	NumberU64() uint64
}

type ChainManager interface {
	GetBlockByNumber(uint64) *types.Block
	CurrentBlock() *types.Block
}