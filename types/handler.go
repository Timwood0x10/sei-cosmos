package types

import (
	sdkacltypes "github.com/cosmos/cosmos-sdk/types/accesscontrol"
)

// Handler defines the core of the state transition function of an application.
type Handler func(ctx Context, msg Msg) (*Result, error)

// AnteHandler authenticates transactions, before their internal messages are handled.
// If newCtx.IsZero(), ctx is used instead.
type AnteHandler func(ctx Context, tx Tx, simulate bool) (newCtx Context, err error)
type AnteDepGenerator func(txDeps []sdkacltypes.AccessOperation, tx Tx) (newTxDeps []sdkacltypes.AccessOperation, err error)

// PostHandler like AnteHandler but it executes after RunMsgs. Runs on success
// or failure and enables use cases like gas refunding.
type PostHandler func(ctx Context, tx Tx, simulate, success bool) (newCtx Context, err error)

// AnteDecorator wraps the next AnteHandler to perform custom pre- and post-processing.
type AnteDecorator interface {
	AnteHandle(ctx Context, tx Tx, simulate bool, next AnteHandler) (newCtx Context, err error)
}

type AnteDepDecorator interface {
	AnteDeps(txDeps []sdkacltypes.AccessOperation, tx Tx, next AnteDepGenerator) (newTxDeps []sdkacltypes.AccessOperation, err error)
}

// PostDecorator wraps the next PostHandler to perform custom post-processing.
type PostDecorator interface {
	PostHandle(ctx Context, tx Tx, simulate, success bool, next PostHandler) (newCtx Context, err error)
}

type DefaultDepDecorator struct{}

// Defeault AnteDeps returned
func (d DefaultDepDecorator) AnteDeps(txDeps []sdkacltypes.AccessOperation, tx Tx, next AnteDepGenerator) (newTxDeps []sdkacltypes.AccessOperation, err error) {
	defaultDeps := []sdkacltypes.AccessOperation{
		{
			ResourceType:       sdkacltypes.ResourceType_ANY,
			AccessType:         sdkacltypes.AccessType_UNKNOWN,
			IdentifierTemplate: "*",
		},
	}
	return next(append(txDeps, defaultDeps...), tx)
}

type NoDepDecorator struct{}

func (d NoDepDecorator) AnteDeps(txDeps []sdkacltypes.AccessOperation, tx Tx, next AnteDepGenerator) (newTxDeps []sdkacltypes.AccessOperation, err error) {
	return next(txDeps, tx)
}

type AnteFullDecorator interface {
	AnteDecorator
	AnteDepDecorator
}

func ChainAnteDecorators(chain ...AnteFullDecorator) (AnteHandler, AnteDepGenerator) {
	anteHandlerChainFunc := ChainAnteDecoratorHandlers(chain...)
	anteHandlerDepGenFunc := chainAnteDecoratorDepGenerators(chain...)
	return anteHandlerChainFunc, anteHandlerDepGenFunc

}

// ChainDecorator chains AnteDecorators together with each AnteDecorator
// wrapping over the decorators further along chain and returns a single AnteHandler.
//
// NOTE: The first element is outermost decorator, while the last element is innermost
// decorator. Decorator ordering is critical since some decorators will expect
// certain checks and updates to be performed (e.g. the Context) before the decorator
// is run. These expectations should be documented clearly in a CONTRACT docline
// in the decorator's godoc.
//
// NOTE: Any application that uses GasMeter to limit transaction processing cost
// MUST set GasMeter with the FIRST AnteDecorator. Failing to do so will cause
// transactions to be processed with an infinite gasmeter and open a DOS attack vector.
// Use `ante.SetUpContextDecorator` or a custom Decorator with similar functionality.
// Returns nil when no AnteDecorator are supplied.
func ChainAnteDecoratorHandlers(chain ...AnteFullDecorator) AnteHandler {
	if len(chain) == 0 {
		return nil
	}

	// handle non-terminated decorators chain
	if (chain[len(chain)-1] != Terminator{}) {
		chain = append(chain, Terminator{})
	}

	return func(ctx Context, tx Tx, simulate bool) (Context, error) {
		return chain[0].AnteHandle(ctx, tx, simulate, ChainAnteDecoratorHandlers(chain[1:]...))
	}
}

func chainAnteDecoratorDepGenerators(chain ...AnteFullDecorator) AnteDepGenerator {
	if len(chain) == 0 {
		return nil
	}

	// handle non-terminated decorators chain
	if (chain[len(chain)-1] != Terminator{}) {
		chain = append(chain, Terminator{})
	}

	return func(txDeps []sdkacltypes.AccessOperation, tx Tx) ([]sdkacltypes.AccessOperation, error) {
		return chain[0].AnteDeps(txDeps, tx, chainAnteDecoratorDepGenerators(chain[1:]...))
	}
}

type WrappedAnteDecorator struct {
	Decorator    AnteDecorator
	DepDecorator AnteDepDecorator
}

func CustomDepWrappedAnteDecorator(decorator AnteDecorator, depDecorator AnteDepDecorator) WrappedAnteDecorator {
	return WrappedAnteDecorator{
		Decorator:    decorator,
		DepDecorator: depDecorator,
	}
}

func DefaultWrappedAnteDecorator(decorator AnteDecorator) WrappedAnteDecorator {
	return WrappedAnteDecorator{
		Decorator: decorator,
		// TODO:: Use DefaultDepDecorator when each decorator defines their own
		//		  See NewConsumeGasForTxSizeDecorator for an example of how to implement a decorator
		DepDecorator: NoDepDecorator{},
	}
}

func (wad WrappedAnteDecorator) AnteHandle(ctx Context, tx Tx, simulate bool, next AnteHandler) (newCtx Context, err error) {
	return wad.Decorator.AnteHandle(ctx, tx, simulate, next)
}

func (wad WrappedAnteDecorator) AnteDeps(txDeps []sdkacltypes.AccessOperation, tx Tx, next AnteDepGenerator) (newTxDeps []sdkacltypes.AccessOperation, err error) {
	return wad.DepDecorator.AnteDeps(txDeps, tx, next)
}

// ChainPostDecorators chains PostDecorators together with each PostDecorator
// wrapping over the decorators further along chain and returns a single PostHandler.
//
// NOTE: The first element is outermost decorator, while the last element is innermost
// decorator. Decorator ordering is critical since some decorators will expect
// certain checks and updates to be performed (e.g. the Context) before the decorator
// is run. These expectations should be documented clearly in a CONTRACT docline
// in the decorator's godoc.
func ChainPostDecorators(chain ...PostDecorator) PostHandler {
	if len(chain) == 0 {
		return nil
	}

	handlerChain := make([]PostHandler, len(chain)+1)
	// set the terminal PostHandler decorator
	handlerChain[len(chain)] = func(ctx Context, tx Tx, simulate, success bool) (Context, error) {
		return ctx, nil
	}
	for i := 0; i < len(chain); i++ {
		ii := i
		handlerChain[ii] = func(ctx Context, tx Tx, simulate, success bool) (Context, error) {
			return chain[ii].PostHandle(ctx, tx, simulate, success, handlerChain[ii+1])
		}
	}
	return handlerChain[0]
}

// Terminator AnteDecorator will get added to the chain to simplify decorator code
// Don't need to check if next == nil further up the chain
//
//	                      ______
//	                   <((((((\\\
//	                   /      . }\
//	                   ;--..--._|}
//	(\                 '--/\--'  )
//	 \\                | '-'  :'|
//	  \\               . -==- .-|
//	   \\               \.__.'   \--._
//	   [\\          __.--|       //  _/'--.
//	   \ \\       .'-._ ('-----'/ __/      \
//	    \ \\     /   __>|      | '--.       |
//	     \ \\   |   \   |     /    /       /
//	      \ '\ /     \  |     |  _/       /
//	       \  \       \ |     | /        /
//	 snd    \  \      \        /
type Terminator struct{}

// Simply return provided Context and nil error
func (t Terminator) AnteHandle(ctx Context, _ Tx, _ bool, _ AnteHandler) (Context, error) {
	return ctx, nil
}

// Simply return provided txDeps and nil error
func (t Terminator) AnteDeps(txDeps []sdkacltypes.AccessOperation, _ Tx, _ AnteDepGenerator) ([]sdkacltypes.AccessOperation, error) {
	return txDeps, nil
}

// PostHandle returns the provided Context and nil error
func (t Terminator) PostHandle(ctx Context, _ Tx, _, _ bool, _ PostHandler) (Context, error) {
	return ctx, nil
}
