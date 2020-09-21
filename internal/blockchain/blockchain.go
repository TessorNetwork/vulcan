// Package blockchain contains code for interacting with the decentr blockchain.
package blockchain

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/Decentr-net/decentr/app"
)

//go:generate mockgen -destination=./blockchain_mock.go -package=blockchain -source=blockchain.go

// ErrInvalidAddress is returned when address is invalid. It is unexpected situation.
var ErrInvalidAddress = errors.New("invalid address")

// Blockchain is interface for interacting with the blockchain.
type Blockchain interface {
	SendStakes(address string, amount int64) error
}

type blockchain struct {
	ctx       context.CLIContext
	txBuilder auth.TxBuilder
}

func NewBlockchain(ctx context.CLIContext, b auth.TxBuilder) Blockchain { // nolint
	return &blockchain{
		ctx:       ctx,
		txBuilder: b,
	}
}

func (b *blockchain) SendStakes(address string, amount int64) error {
	to, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidAddress, address)
	}

	msg := bank.NewMsgSend(b.ctx.GetFromAddress(), to, sdk.Coins{sdk.Coin{
		Denom:  app.DefaultBondDenom,
		Amount: sdk.NewInt(amount),
	}})
	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	return b.BroadcastMsg(msg)
}

func (b *blockchain) BroadcastMsg(msg sdk.Msg) error {
	txBldr, err := utils.PrepareTxBuilder(b.txBuilder, b.ctx)
	if err != nil {
		return fmt.Errorf("failed to prepare builder: %w", err)
	}

	msgs := []sdk.Msg{msg}

	if txBldr, err = utils.EnrichWithGas(txBldr, b.ctx, msgs); err != nil {
		return errors.New("failed to calculate gas") // nolint: goerr113
	}

	txBytes, err := txBldr.BuildAndSign(b.ctx.GetFromName(), keys.DefaultKeyPass, msgs)
	if err != nil {
		return fmt.Errorf("failed to build and sign tx: %w", err)
	}

	resp, err := b.ctx.BroadcastTx(txBytes)
	if err != nil {
		return fmt.Errorf("failed to broadcast tx: %w", err)
	}

	if resp.Height == 0 {
		return fmt.Errorf("failed to broadcast tx: %s", resp.String()) // nolint: goerr113
	}

	return nil
}
