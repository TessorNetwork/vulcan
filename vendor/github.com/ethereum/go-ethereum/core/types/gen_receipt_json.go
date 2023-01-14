// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*receiptMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (r Receipt) MarshalJSON() ([]byte, error) {
	type Receipt struct {
		Type              hexutil.Uint64 `json:"type,omitempty"`
		PostState         hexutil.Bytes  `json:"root"`
		Status            hexutil.Uint64 `json:"status"`
		CumulativeGasUsed hexutil.Uint64 `json:"cumulativeGasUsed" gencodec:"required"`
		Bloom             Bloom          `json:"logsBloom"         gencodec:"required"`
		Logs              []*Log         `json:"logs"              gencodec:"required"`
		TxHash            common.Hash    `json:"transactionHash" gencodec:"required"`
		ContractAddress   common.Address `json:"contractAddress"`
		GasUsed           hexutil.Uint64 `json:"gasUsed" gencodec:"required"`
		BlockHash         common.Hash    `json:"blockHash,omitempty"`
		BlockNumber       *hexutil.Big   `json:"blockNumber,omitempty"`
		TransactionIndex  hexutil.Uint   `json:"transactionIndex"`
	}
	var enc Receipt
	enc.Type = hexutil.Uint64(r.Type)
	enc.PostState = r.PostState
	enc.Status = hexutil.Uint64(r.Status)
	enc.CumulativeGasUsed = hexutil.Uint64(r.CumulativeGasUsed)
	enc.Bloom = r.Bloom
	enc.Logs = r.Logs
	enc.TxHash = r.TxHash
	enc.ContractAddress = r.ContractAddress
	enc.GasUsed = hexutil.Uint64(r.GasUsed)
	enc.BlockHash = r.BlockHash
	enc.BlockNumber = (*hexutil.Big)(r.BlockNumber)
	enc.TransactionIndex = hexutil.Uint(r.TransactionIndex)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *Receipt) UnmarshalJSON(input []byte) error {
	type Receipt struct {
		Type              *hexutil.Uint64 `json:"type,omitempty"`
		PostState         *hexutil.Bytes  `json:"root"`
		Status            *hexutil.Uint64 `json:"status"`
		CumulativeGasUsed *hexutil.Uint64 `json:"cumulativeGasUsed" gencodec:"required"`
		Bloom             *Bloom          `json:"logsBloom"         gencodec:"required"`
		Logs              []*Log          `json:"logs"              gencodec:"required"`
		TxHash            *common.Hash    `json:"transactionHash" gencodec:"required"`
		ContractAddress   *common.Address `json:"contractAddress"`
		GasUsed           *hexutil.Uint64 `json:"gasUsed" gencodec:"required"`
		BlockHash         *common.Hash    `json:"blockHash,omitempty"`
		BlockNumber       *hexutil.Big    `json:"blockNumber,omitempty"`
		TransactionIndex  *hexutil.Uint   `json:"transactionIndex"`
	}
	var fur Receipt
	if err := json.Unmarshal(input, &fur); err != nil {
		return err
	}
	if fur.Type != nil {
		r.Type = uint8(*fur.Type)
	}
	if fur.PostState != nil {
		r.PostState = *fur.PostState
	}
	if fur.Status != nil {
		r.Status = uint64(*fur.Status)
	}
	if fur.CumulativeGasUsed == nil {
		return errors.New("missing required field 'cumulativeGasUsed' for Receipt")
	}
	r.CumulativeGasUsed = uint64(*fur.CumulativeGasUsed)
	if fur.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Receipt")
	}
	r.Bloom = *fur.Bloom
	if fur.Logs == nil {
		return errors.New("missing required field 'logs' for Receipt")
	}
	r.Logs = fur.Logs
	if fur.TxHash == nil {
		return errors.New("missing required field 'transactionHash' for Receipt")
	}
	r.TxHash = *fur.TxHash
	if fur.ContractAddress != nil {
		r.ContractAddress = *fur.ContractAddress
	}
	if fur.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Receipt")
	}
	r.GasUsed = uint64(*fur.GasUsed)
	if fur.BlockHash != nil {
		r.BlockHash = *fur.BlockHash
	}
	if fur.BlockNumber != nil {
		r.BlockNumber = (*big.Int)(fur.BlockNumber)
	}
	if fur.TransactionIndex != nil {
		r.TransactionIndex = uint(*fur.TransactionIndex)
	}
	return nil
}
