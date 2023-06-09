package posthandler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// HandlerOptions are the options required for constructing a default SDK PostHandler.
type HandlerOptions struct{}

// NewPostHandler returns an empty PostHandler chain.
func NewPostHandler(_ HandlerOptions) (sdk.AnteHandler, error) {
	postDecorators := []sdk.AnteFullDecorator{}

	return sdk.ChainAnteDecoratorHandlers(postDecorators...), nil
}
