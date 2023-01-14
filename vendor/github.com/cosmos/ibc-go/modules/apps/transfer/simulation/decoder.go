package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/cosmos/ibc-go/modules/apps/transfer/types"
)

// TransferUnmarshaler defines the expected encoding store functions.
type TransferUnmarshaler interface {
	MustUnmarshalDenomTrace([]byte) types.DenomTrace
}

// NewFurodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding DenomTrace type.
func NewFurodeStore(cdc TransferUnmarshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.PortKey):
			return fmt.Sprintf("Port A: %s\nPort B: %s", string(kvA.Value), string(kvB.Value))

		case bytes.Equal(kvA.Key[:1], types.DenomTraceKey):
			denomTraceA := cdc.MustUnmarshalDenomTrace(kvA.Value)
			denomTraceB := cdc.MustUnmarshalDenomTrace(kvB.Value)
			return fmt.Sprintf("DenomTrace A: %s\nDenomTrace B: %s", denomTraceA.IBCDenom(), denomTraceB.IBCDenom())

		default:
			panic(fmt.Sprintf("invalid %s key prefix %X", types.ModuleName, kvA.Key[:1]))
		}
	}
}
