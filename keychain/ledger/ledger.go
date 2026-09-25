// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ledger

import (
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
)

// Semantic is the ledger app version. avalanchego v1.15.0 removed version.Semantic,
// so the SDK carries its own copy of the shape the ledger device reports.
type Semantic struct {
	Major int
	Minor int
	Patch int
}

// Ledger interface for the ledger wrapper
type Ledger interface {
	Version() (v *Semantic, err error)
	PubKey(addressIndex uint32) (*secp256k1.PublicKey, error)
	PubKeys(addressIndices []uint32) ([]*secp256k1.PublicKey, error)
	SignHash(hash []byte, addressIndices []uint32) ([][]byte, error)
	Sign(unsignedTxBytes []byte, addressIndices []uint32) ([][]byte, error)
	Disconnect() error
}

// Verify that the Device implementation satisfies the Ledger interface
var _ Ledger = (*Device)(nil)
