package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/ibc-transfer/types"
)

const (
	restChannelID = "channel-id"
	restPortID    = "port-id"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	registerTxRoutes(clientCtx, r)
}

// TransferTxReq defines the properties of a transfer tx request's body.
type TransferTxReq struct {
	BaseReq          rest.BaseReq     `json:"base_req" yaml:"base_req"`
	Token            sdk.Coin         `json:"token" yaml:"token"`
	DenomTrace       types.DenomTrace `json:"denom_trace" yaml:"denom_trace"`
	Receiver         string           `json:"receiver" yaml:"receiver"`
	TimeoutHeight    uint64           `json:"timeout_height" yaml:"timeout_height"`
	TimeoutTimestamp uint64           `json:"timeout_timestamp" yaml:"timeout_timestamp"`
}
