// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/BOXFoundation/boxd/core"
	"github.com/BOXFoundation/boxd/crypto"
	"golang.org/x/crypto/ripemd160"
)

var addressTypeP2PKHPrefix = [2]byte{FixPrefix, 0x26}
var addressTypeP2SHPrefix = [2]byte{FixPrefix, 0x2b}

// const
const (
	BoxPrefix           = 'b'
	AddressPrefixLength = 2
	FixPrefix           = 0x13

	AddressLength       = 26
	EncodeAddressLength = 35
)

// AddressHash Alias for address hash
type AddressHash [ripemd160.Size]byte

// Address is an interface type for any type of destination a transaction output may spend to.
type Address interface {
	String() string
	SetString(string) error
	Hash() []byte
	Hash160() *AddressHash
}

// AddressPubKeyHash is an Address for a pay-to-pubkey-hash (P2PKH) transaction.
type AddressPubKeyHash struct {
	hash AddressHash
}

// NewAddressPubKeyHash returns a new AddressPubKeyHash.  pkHash mustbe 20 bytes.
func NewAddressPubKeyHash(pkHash []byte) (*AddressPubKeyHash, error) {
	return newAddressPubKeyHash(pkHash)
}

// NewAddressFromPubKey returns a new AddressPubKeyHash derived from an ecdsa public key
func NewAddressFromPubKey(pubKey *crypto.PublicKey) (*AddressPubKeyHash, error) {
	pkHash := crypto.Hash160(pubKey.Serialize())
	return newAddressPubKeyHash(pkHash)
}

// NewAddress creates an address from string
func NewAddress(address string) (Address, error) {
	addr := &AddressPubKeyHash{}
	err := addr.SetString(address)
	return addr, err
}

func newAddressPubKeyHash(pkHash []byte) (*AddressPubKeyHash, error) {
	// Check for a valid pubkey hash length.
	if len(pkHash) != ripemd160.Size {
		return nil, core.ErrInvalidPKHash
	}

	addr := &AddressPubKeyHash{}
	copy(addr.hash[:], pkHash)
	return addr, nil
}

// Hash returns the bytes to be included in a txout script to pay to a pubkey hash.
func (a *AddressPubKeyHash) Hash() []byte {
	return a.hash[:]
}

// String returns a human-readable string for the pay-to-pubkey-hash address.
func (a *AddressPubKeyHash) String() string {
	return encodeAddress(a.hash[:])
}

// SetString sets the Address's internal byte array using byte array decoded from input
// base58 format string, returns error if input string is invalid
func (a *AddressPubKeyHash) SetString(in string) error {
	if len(in) != EncodeAddressLength || in[0] != BoxPrefix {
		return core.ErrInvalidAddressString
	}
	rawBytes, err := crypto.Base58CheckDecode(in)
	if err != nil {
		return err
	}
	if len(rawBytes) != 22 {
		return core.ErrInvalidAddressString
	}
	var prefix [2]byte
	copy(prefix[:], rawBytes[:2])
	if prefix != addressTypeP2PKHPrefix && prefix != addressTypeP2SHPrefix {
		return core.ErrInvalidAddressString
	}
	copy(a.hash[:], rawBytes[2:])
	return nil
}

// Hash160 returns the underlying array of the pubkey hash.
func (a *AddressPubKeyHash) Hash160() *AddressHash {
	return &a.hash
}

func encodeAddress(hash []byte) string {
	b := make([]byte, 0, len(hash)+2)
	b = append(b, addressTypeP2PKHPrefix[:]...)
	b = append(b, hash[:]...)
	return crypto.Base58CheckEncode(b)
}
