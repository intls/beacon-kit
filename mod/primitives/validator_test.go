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

package primitives_test

import (
	"testing"

	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/constants"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/stretchr/testify/require"
)

func TestNewValidatorFromDeposit(t *testing.T) {
	tests := []struct {
		name                      string
		pubkey                    crypto.BLSPubkey
		withdrawalCredentials     primitives.WithdrawalCredentials
		amount                    math.Gwei
		effectiveBalanceIncrement math.Gwei
		maxEffectiveBalance       math.Gwei
		want                      *primitives.Validator
	}{
		{
			name:   "normal case",
			pubkey: [48]byte{0x01},
			withdrawalCredentials: primitives.
				NewCredentialsFromExecutionAddress(
					common.ExecutionAddress{0x01},
				),
			amount:                    32e9,
			effectiveBalanceIncrement: 1e9,
			maxEffectiveBalance:       32e9,
			want: &primitives.Validator{
				Pubkey: [48]byte{0x01},
				WithdrawalCredentials: primitives.
					NewCredentialsFromExecutionAddress(
						common.ExecutionAddress{0x01},
					),
				EffectiveBalance: 32e9,
				Slashed:          false,
				ActivationEligibilityEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
				ActivationEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
				ExitEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
				WithdrawableEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
			},
		},
		{
			name:   "effective balance capped at max",
			pubkey: [48]byte{0x02},
			withdrawalCredentials: primitives.
				NewCredentialsFromExecutionAddress(
					common.ExecutionAddress{0x02},
				),
			amount:                    40e9,
			effectiveBalanceIncrement: 1e9,
			maxEffectiveBalance:       32e9,
			want: &primitives.Validator{
				Pubkey: [48]byte{0x02},
				WithdrawalCredentials: primitives.
					NewCredentialsFromExecutionAddress(
						common.ExecutionAddress{0x02},
					),
				EffectiveBalance: 32e9,
				Slashed:          false,
				ActivationEligibilityEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
				ActivationEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
				ExitEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
				WithdrawableEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
			},
		},
		{
			name:   "effective balance rounded down",
			pubkey: [48]byte{0x03},
			withdrawalCredentials: primitives.
				NewCredentialsFromExecutionAddress(
					common.ExecutionAddress{0x03},
				),
			amount:                    32.5e9,
			effectiveBalanceIncrement: 1e9,
			maxEffectiveBalance:       32e9,
			want: &primitives.Validator{
				Pubkey: [48]byte{0x03},
				WithdrawalCredentials: primitives.
					NewCredentialsFromExecutionAddress(
						common.ExecutionAddress{0x03},
					),
				EffectiveBalance: 32e9,
				Slashed:          false,
				ActivationEligibilityEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
				ActivationEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
				ExitEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
				WithdrawableEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := primitives.NewValidatorFromDeposit(
				tt.pubkey,
				tt.withdrawalCredentials,
				tt.amount,
				tt.effectiveBalanceIncrement,
				tt.maxEffectiveBalance,
			)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestValidator_IsActive(t *testing.T) {
	tests := []struct {
		name      string
		epoch     math.Epoch
		validator *primitives.Validator
		want      bool
	}{
		{
			name:  "active",
			epoch: 10,
			validator: &primitives.Validator{
				ActivationEpoch: 5,
				ExitEpoch:       15,
			},
			want: true,
		},
		{
			name:  "not active, before activation",
			epoch: 4,
			validator: &primitives.Validator{
				ActivationEpoch: 5,
				ExitEpoch:       15,
			},
			want: false,
		},
		{
			name:  "not active, after exit",
			epoch: 16,
			validator: &primitives.Validator{
				ActivationEpoch: 5,
				ExitEpoch:       15,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.validator.IsActive(tt.epoch))
		})
	}
}

func TestValidator_IsEligibleForActivation(t *testing.T) {
	tests := []struct {
		name           string
		finalizedEpoch math.Epoch
		validator      *primitives.Validator
		want           bool
	}{
		{
			name:           "eligible",
			finalizedEpoch: 10,
			validator: &primitives.Validator{
				ActivationEligibilityEpoch: 5,
				ActivationEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
			},
			want: true,
		},
		{
			name:           "not eligible, activation eligibility in future",
			finalizedEpoch: 4,
			validator: &primitives.Validator{
				ActivationEligibilityEpoch: 5,
				ActivationEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
			},
			want: false,
		},
		{
			name:           "not eligible, already activated",
			finalizedEpoch: 10,
			validator: &primitives.Validator{
				ActivationEligibilityEpoch: 5,
				ActivationEpoch:            8,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(
				t,
				tt.want,
				tt.validator.IsEligibleForActivation(tt.finalizedEpoch),
			)
		})
	}
}

func TestValidator_IsEligibleForActivationQueue(t *testing.T) {
	maxEffectiveBalance := math.Gwei(32e9)
	tests := []struct {
		name      string
		validator *primitives.Validator
		want      bool
	}{
		{
			name: "eligible",
			validator: &primitives.Validator{
				ActivationEligibilityEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
				EffectiveBalance: maxEffectiveBalance,
			},
			want: true,
		},
		{
			name: "not eligible, activation eligibility set",
			validator: &primitives.Validator{
				ActivationEligibilityEpoch: 5,
				EffectiveBalance:           maxEffectiveBalance,
			},
			want: false,
		},
		{
			name: "not eligible, effective balance too low",
			validator: &primitives.Validator{
				ActivationEligibilityEpoch: math.Epoch(
					constants.FarFutureEpoch,
				),
				EffectiveBalance: maxEffectiveBalance - 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(
				t,
				tt.want,
				tt.validator.IsEligibleForActivationQueue(maxEffectiveBalance),
			)
		})
	}
}

func TestValidator_IsSlashable(t *testing.T) {
	tests := []struct {
		name      string
		epoch     math.Epoch
		validator *primitives.Validator
		want      bool
	}{
		{
			name:  "slashable",
			epoch: 10,
			validator: &primitives.Validator{
				Slashed:           false,
				ActivationEpoch:   5,
				WithdrawableEpoch: 15,
			},
			want: true,
		},
		{
			name:  "not slashable, already slashed",
			epoch: 10,
			validator: &primitives.Validator{
				Slashed:           true,
				ActivationEpoch:   5,
				WithdrawableEpoch: 15,
			},
			want: false,
		},
		{
			name:  "not slashable, before activation",
			epoch: 4,
			validator: &primitives.Validator{
				Slashed:           false,
				ActivationEpoch:   5,
				WithdrawableEpoch: 15,
			},
			want: false,
		},
		{
			name:  "not slashable, after withdrawable",
			epoch: 16,
			validator: &primitives.Validator{
				Slashed:           false,
				ActivationEpoch:   5,
				WithdrawableEpoch: 15,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.validator.IsSlashable(tt.epoch))
		})
	}
}

func TestValidator_IsFullyWithdrawable(t *testing.T) {
	tests := []struct {
		name      string
		balance   math.Gwei
		epoch     math.Epoch
		validator *primitives.Validator
		want      bool
	}{
		{
			name:    "fully withdrawable",
			balance: 32e9,
			epoch:   10,
			validator: &primitives.Validator{
				WithdrawalCredentials: primitives.
					NewCredentialsFromExecutionAddress(
						common.ExecutionAddress{0x01},
					),
				WithdrawableEpoch: 5,
			},
			want: true,
		},
		{
			name:    "not fully withdrawable, non-eth1 credentials",
			balance: 32e9,
			epoch:   10,
			validator: &primitives.Validator{
				WithdrawalCredentials: primitives.
					WithdrawalCredentials{0x00},
				WithdrawableEpoch: 5,
			},
			want: false,
		},
		{
			name:    "not fully withdrawable, before withdrawable epoch",
			balance: 32e9,
			epoch:   4,
			validator: &primitives.Validator{
				WithdrawalCredentials: primitives.
					NewCredentialsFromExecutionAddress(
						common.ExecutionAddress{0x01},
					),
				WithdrawableEpoch: 5,
			},
			want: false,
		},
		{
			name:    "not fully withdrawable, zero balance",
			balance: 0,
			epoch:   10,
			validator: &primitives.Validator{
				WithdrawalCredentials: primitives.
					NewCredentialsFromExecutionAddress(
						common.ExecutionAddress{0x01},
					),
				WithdrawableEpoch: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(
				t,
				tt.want,
				tt.validator.IsFullyWithdrawable(tt.balance, tt.epoch),
			)
		})
	}
}

func TestValidator_IsPartiallyWithdrawable(t *testing.T) {
	maxEffectiveBalance := math.Gwei(32e9)
	tests := []struct {
		name      string
		balance   math.Gwei
		validator *primitives.Validator
		want      bool
	}{
		{
			name:    "partially withdrawable",
			balance: 33e9,
			validator: &primitives.Validator{
				WithdrawalCredentials: primitives.
					NewCredentialsFromExecutionAddress(
						common.ExecutionAddress{0x01},
					),
				EffectiveBalance: maxEffectiveBalance,
			},
			want: true,
		},
		{
			name:    "not partially withdrawable, non-eth1 credentials",
			balance: 33e9,
			validator: &primitives.Validator{
				WithdrawalCredentials: primitives.WithdrawalCredentials{
					0x00,
				},
				EffectiveBalance: maxEffectiveBalance,
			},
			want: false,
		},
		{
			name:    "not partially withdrawable, not at max effective balance",
			balance: 33e9,
			validator: &primitives.Validator{
				WithdrawalCredentials: primitives.
					NewCredentialsFromExecutionAddress(
						common.ExecutionAddress{0x01},
					),
				EffectiveBalance: maxEffectiveBalance - 1,
			},
			want: false,
		},
		{
			name:    "not partially withdrawable, no excess balance",
			balance: 32e9,
			validator: &primitives.Validator{
				WithdrawalCredentials: primitives.
					NewCredentialsFromExecutionAddress(
						common.ExecutionAddress{0x01},
					),
				EffectiveBalance: maxEffectiveBalance,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(
				t,
				tt.want,
				tt.validator.IsPartiallyWithdrawable(
					tt.balance,
					maxEffectiveBalance,
				),
			)
		})
	}
}

func TestValidator_HasEth1WithdrawalCredentials(t *testing.T) {
	tests := []struct {
		name      string
		validator *primitives.Validator
		want      bool
	}{
		{
			name: "has eth1 credentials",
			validator: &primitives.Validator{
				WithdrawalCredentials: primitives.
					NewCredentialsFromExecutionAddress(
						common.ExecutionAddress{0x01},
					),
			},
			want: true,
		},
		{
			name: "does not have eth1 credentials",
			validator: &primitives.Validator{
				WithdrawalCredentials: primitives.WithdrawalCredentials{
					0x00,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(
				t,
				tt.want,
				tt.validator.HasEth1WithdrawalCredentials(),
			)
		})
	}
}

func TestValidator_HasMaxEffectiveBalance(t *testing.T) {
	maxEffectiveBalance := math.Gwei(32e9)
	tests := []struct {
		name      string
		validator *primitives.Validator
		want      bool
	}{
		{
			name: "has max effective balance",
			validator: &primitives.Validator{
				EffectiveBalance: maxEffectiveBalance,
			},
			want: true,
		},
		{
			name: "does not have max effective balance",
			validator: &primitives.Validator{
				EffectiveBalance: maxEffectiveBalance - 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(
				t,
				tt.want,
				tt.validator.HasMaxEffectiveBalance(maxEffectiveBalance),
			)
		})
	}
}