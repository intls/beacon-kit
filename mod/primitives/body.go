// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package primitives

import (
	"errors"

	// engineprimitives "github.com/berachain/beacon-kit/mod/primitives-engine".
	"github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	eip4844 "github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/ssz"
)

const (
	// BodyLengthDeneb is the number of fields in the BeaconBlockBodyDeneb
	// struct.
	BodyLengthDeneb uint64 = 6

	// KZGPosition is the position of BlobKzgCommitments in the block body.
	KZGPositionDeneb = BodyLengthDeneb - 1
)

// BeaconBlockBodyBase represents the base body of a beacon block that is
// shared between all forks.
type BeaconBlockBodyBase struct {
	// RandaoReveal is the reveal of the RANDAO.
	RandaoReveal crypto.BLSSignature `ssz-size:"96"`

	// Eth1Data is the data from the Eth1 chain.
	Eth1Data *Eth1Data

	// Graffiti is for a fun message or meme.
	Graffiti [32]byte `ssz-size:"32"`

	// Deposits is the list of deposits included in the body.
	Deposits []*Deposit `ssz-max:"16"`
}

// GetRandaoReveal returns the RandaoReveal of the Body.
func (b *BeaconBlockBodyBase) GetRandaoReveal() crypto.BLSSignature {
	return b.RandaoReveal
}

// SetRandaoReveal sets the RandaoReveal of the Body.
func (b *BeaconBlockBodyBase) SetRandaoReveal(reveal crypto.BLSSignature) {
	b.RandaoReveal = reveal
}

// GetEth1Data returns the Eth1Data of the Body.
func (b *BeaconBlockBodyBase) GetEth1Data() *Eth1Data {
	return b.Eth1Data
}

// SetBlobKzgCommitments sets the BlobKzgCommitments of the
// BeaconBlockBodyDeneb.
func (b *BeaconBlockBodyDeneb) SetEth1Data(eth1Data *Eth1Data) {
	b.Eth1Data = eth1Data
}

// GetGraffiti returns the Graffiti of the Body.
func (b *BeaconBlockBodyBase) GetGraffiti() bytes.B32 {
	return b.Graffiti
}

// GetDeposits returns the Deposits of the BeaconBlockBodyBase.
func (b *BeaconBlockBodyBase) GetDeposits() []*Deposit {
	return b.Deposits
}

// SetDeposits sets the Deposits of the BeaconBlockBodyBase.
func (b *BeaconBlockBodyBase) SetDeposits(deposits []*Deposit) {
	b.Deposits = deposits
}

// BeaconBlockBodyDeneb represents the body of a beacon block in the Deneb
// chain.
//
//go:generate go run github.com/ferranbt/fastssz/sszgen --path ./body.go -objs BeaconBlockBodyDeneb -include ./pkg/crypto,./primitives.go,./payload.go,./pkg/eip4844,./pkg/bytes,./eth1data.go,./pkg/math,./pkg/common,./deposit.go,./withdrawal_credentials.go,./withdrawal.go,$GETH_PKG_INCLUDE/common,$GETH_PKG_INCLUDE/common/hexutil -output body.ssz.go
type BeaconBlockBodyDeneb struct {
	BeaconBlockBodyBase

	// ExecutionPayload is the execution payload of the body.
	ExecutionPayload *ExecutableDataDeneb

	// BlobKzgCommitments is the list of KZG commitments for the EIP-4844 blobs.
	BlobKzgCommitments []eip4844.KZGCommitment `ssz-size:"?,48" ssz-max:"16"`
}

// IsNil checks if the BeaconBlockBodyDeneb is nil.
func (b *BeaconBlockBodyDeneb) IsNil() bool {
	return b == nil
}

// GetExecutionPayload returns the ExecutionPayload of the Body.
func (b *BeaconBlockBodyDeneb) GetExecutionPayload() ExecutionPayload {
	return b.ExecutionPayload
}

// SetExecutionData sets the ExecutionData of the BeaconBlockBodyDeneb.
func (b *BeaconBlockBodyDeneb) SetExecutionData(
	executionData ExecutionPayload,
) error {
	var ok bool
	b.ExecutionPayload, ok = executionData.(*ExecutableDataDeneb)
	if !ok {
		return errors.New("invalid execution data type")
	}
	return nil
}

// GetBlobKzgCommitments returns the BlobKzgCommitments of the Body.
//
//nolint:lll // annoying to fix.
func (b *BeaconBlockBodyDeneb) GetBlobKzgCommitments() eip4844.KZGCommitments[common.ExecutionHash] {
	return b.BlobKzgCommitments
}

// SetBlobKzgCommitments sets the BlobKzgCommitments of the
// BeaconBlockBodyDeneb.
func (b *BeaconBlockBodyDeneb) SetBlobKzgCommitments(
	commitments eip4844.KZGCommitments[common.ExecutionHash],
) {
	b.BlobKzgCommitments = commitments
}

// GetTopLevelRoots returns the top-level roots of the BeaconBlockBodyDeneb.
func (b *BeaconBlockBodyDeneb) GetTopLevelRoots() ([][32]byte, error) {
	layer := make([][32]byte, BodyLengthDeneb)
	var err error
	randao := b.GetRandaoReveal()
	layer[0], err = ssz.MerkleizeByteSlice[math.U64, [32]byte](randao[:])
	if err != nil {
		return nil, err
	}

	layer[1], err = b.Eth1Data.HashTreeRoot()
	if err != nil {
		return nil, err
	}

	// graffiti
	layer[2] = b.GetGraffiti()

	layer[3], err = Deposits(b.GetDeposits()).HashTreeRoot()
	if err != nil {
		return nil, err
	}

	// Execution Payload
	layer[4], err = b.GetExecutionPayload().HashTreeRoot()
	if err != nil {
		return nil, err
	}

	// KZG commitments is not needed
	return layer, nil
}

func (b *BeaconBlockBodyDeneb) AttachExecution(
	executionData ExecutionPayload,
) error {
	var ok bool
	b.ExecutionPayload, ok = executionData.(*ExecutableDataDeneb)
	if !ok {
		return errors.New("invalid execution data type")
	}
	return nil
}