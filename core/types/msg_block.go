// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package types

import (
	"errors"

	corepb "github.com/BOXFoundation/Quicksilver/core/pb"
	"github.com/BOXFoundation/Quicksilver/crypto"
	conv "github.com/BOXFoundation/Quicksilver/p2p/convert"
	proto "github.com/gogo/protobuf/proto"
)

// Define error
var (
	ErrSerializeHeader                = errors.New("Serialize block header error")
	ErrEmptyProtoMessage              = errors.New("Empty proto message")
	ErrInvalidBlockHeaderProtoMessage = errors.New("Invalid block header proto message")
	ErrInvalidBlockProtoMessage       = errors.New("Invalid block proto message")
)

// MsgBlock defines a block containing header and transactions within it.
type MsgBlock struct {
	Header *BlockHeader
	Txs    []*MsgTx
	Height int32
}

var _ conv.Convertible = (*MsgBlock)(nil)
var _ conv.Serializable = (*MsgBlock)(nil)

// BlockHeader defines information about a block and is used in the
// block (MsgBlock) and headers (MsgHeaders) messages.
type BlockHeader struct {
	// Version of the block.  This is not the same as the protocol version.
	Version int32

	// Hash of the previous block header in the block chain.
	PrevBlockHash crypto.HashType

	// Merkle tree reference to hash of all transactions for the block.
	TxsRoot crypto.HashType

	// Time the block was created.  This is, unfortunately, encoded as a
	// uint32 on the wire and therefore is limited to 2106.
	TimeStamp int64

	// Distinguish between mainnet and testnet.
	Magic uint32
}

var _ conv.Convertible = (*BlockHeader)(nil)
var _ conv.Serializable = (*BlockHeader)(nil)

////////////////////////////////////////////////////////////////////////////////

// ToProtoMessage converts block to proto message.
func (msgBlock *MsgBlock) ToProtoMessage() (proto.Message, error) {
	header, _ := msgBlock.Header.ToProtoMessage()
	if header, ok := header.(*corepb.BlockHeader); ok {
		var txs []*corepb.MsgTx
		for _, v := range msgBlock.Txs {
			tx, err := v.ToProtoMessage()
			if err != nil {
				return nil, err
			}
			if tx, ok := tx.(*corepb.MsgTx); ok {
				txs = append(txs, tx)
			}
		}
		return &corepb.MsgBlock{
			Header: header,
			Txs:    txs,
			Height: msgBlock.Height,
		}, nil
	}

	return nil, ErrSerializeHeader
}

// FromProtoMessage converts proto message to block.
func (msgBlock *MsgBlock) FromProtoMessage(message proto.Message) error {
	if message, ok := message.(*corepb.MsgBlock); ok {
		if message != nil {
			header := new(BlockHeader)
			if err := header.FromProtoMessage(message.Header); err != nil {
				return err
			}
			var txs []*MsgTx
			for _, v := range message.Txs {
				tx := new(MsgTx)
				if err := tx.FromProtoMessage(v); err != nil {
					return err
				}
				txs = append(txs, tx)
			}
			msgBlock.Header = header
			msgBlock.Txs = txs
			msgBlock.Height = message.Height
			return nil
		}
		return ErrEmptyProtoMessage
	}

	return ErrInvalidBlockProtoMessage
}

// Marshal method marshal MsgBlock object to binary
func (msgBlock *MsgBlock) Marshal() (data []byte, err error) {
	return conv.MarshalConvertible(msgBlock)
}

// Unmarshal method unmarshal binary data to MsgBlock object
func (msgBlock *MsgBlock) Unmarshal(data []byte) error {
	msg := &corepb.MsgBlock{}
	if err := proto.Unmarshal(data, msg); err != nil {
		return err
	}
	return msgBlock.FromProtoMessage(msg)
}

// BlockHash returns the block identifier hash for the Block.
func (msgBlock *MsgBlock) BlockHash() (*crypto.HashType, error) {
	pbHeader, err := msgBlock.Header.ToProtoMessage()
	if err != nil {
		return nil, err
	}
	header, err := proto.Marshal(pbHeader)
	if err != nil {
		return nil, err
	}
	hash := crypto.DoubleHashH(header)
	return &hash, nil
}

////////////////////////////////////////////////////////////////////////////////

// ToProtoMessage converts block header to proto message.
func (header *BlockHeader) ToProtoMessage() (proto.Message, error) {

	return &corepb.BlockHeader{
		Version:       header.Version,
		PrevBlockHash: header.PrevBlockHash[:],
		TxsRoot:       header.TxsRoot[:],
		TimeStamp:     header.TimeStamp,
		Magic:         header.Magic,
	}, nil
}

// FromProtoMessage converts proto message to block header.
func (header *BlockHeader) FromProtoMessage(message proto.Message) error {

	if message, ok := message.(*corepb.BlockHeader); ok {
		if message != nil {
			header.Version = message.Version
			copy(header.PrevBlockHash[:], message.PrevBlockHash)
			copy(header.TxsRoot[:], message.TxsRoot)
			header.TimeStamp = message.TimeStamp
			header.Magic = message.Magic
			return nil
		}
		return ErrEmptyProtoMessage
	}

	return ErrInvalidBlockHeaderProtoMessage
}

// Marshal method marshal BlockHeader object to binary
func (header *BlockHeader) Marshal() (data []byte, err error) {
	return conv.MarshalConvertible(header)
}

// Unmarshal method unmarshal binary data to BlockHeader object
func (header *BlockHeader) Unmarshal(data []byte) error {
	msg := &corepb.BlockHeader{}
	if err := proto.Unmarshal(data, msg); err != nil {
		return err
	}
	return header.FromProtoMessage(msg)
}
