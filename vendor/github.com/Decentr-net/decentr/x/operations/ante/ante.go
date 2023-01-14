package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/TessorNetwork/furya/x/operations/keeper"
)

type MinGasPriceDecorator struct {
	keeper keeper.Keeper
}

func NewMinGasPriceDecorator(keeper keeper.Keeper) *MinGasPriceDecorator {
	return &MinGasPriceDecorator{keeper: keeper}
}

func (mgp MinGasPriceDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	price := mgp.keeper.GetParams(ctx.WithGasMeter(sdk.NewInfiniteGasMeter())).MinGasPrice
	ctx = ctx.WithMinGasPrices(sdk.NewFurCoins(price))
	return next(ctx, tx, simulate)
}
