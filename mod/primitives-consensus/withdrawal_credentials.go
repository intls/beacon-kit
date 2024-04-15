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
// conditions
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

package consensusprimitives

import (
	"github.com/berachain/beacon-kit/mod/primitives"
)

// EthSecp256k1CredentialPrefix is the prefix for an Ethereum secp256k1.
const EthSecp256k1CredentialPrefix = byte(iota + 1)

// WithdrawalCredentials is a staking credential that is used to identify a
// validator.
type WithdrawalCredentials primitives.Bytes32

// NewCredentialsFromExecutionAddress creates a new WithdrawalCredentials from
// an.
func NewCredentialsFromExecutionAddress(
	address primitives.ExecutionAddress,
) WithdrawalCredentials {
	credentials := WithdrawalCredentials{}
	credentials[0] = 0x01
	copy(credentials[12:], address[:])
	return credentials
}

// ToExecutionAddress converts the WithdrawalCredentials to an ExecutionAddress.
func (wc WithdrawalCredentials) ToExecutionAddress() (
	primitives.ExecutionAddress,
	error,
) {
	if wc[0] != EthSecp256k1CredentialPrefix {
		return primitives.ExecutionAddress{}, ErrInvalidWithdrawalCredentials
	}
	return primitives.ExecutionAddress(wc[12:]), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Bytes32.
// TODO: Figure out how to not have to do this.
func (h *WithdrawalCredentials) UnmarshalJSON(input []byte) error {
	return (*primitives.Bytes32)(h).UnmarshalJSON(input)
}

// String returns the hex string representation of Bytes32.
// TODO: Figure out how to not have to do this.
func (h WithdrawalCredentials) String() string {
	return primitives.Bytes32(h).String()
}

// MarshalText implements the encoding.TextMarshaler interface for Bytes32.
// TODO: Figure out how to not have to do this.
func (h WithdrawalCredentials) MarshalText() ([]byte, error) {
	return primitives.Bytes32(h).MarshalText()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Bytes32.
// TODO: Figure out how to not have to do this.
func (h *WithdrawalCredentials) UnmarshalText(text []byte) error {
	return (*primitives.Bytes32)(h).UnmarshalText(text)
}