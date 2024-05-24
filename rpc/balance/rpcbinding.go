// Package balance contains RPC wrappers for NeoFS Balance contract.
//
// Code generated by neo-go contract generate-rpcwrapper --manifest <file.json> --out <file.go> [--hash <hash>] [--config <config>]; DO NOT EDIT.
package balance

import (
	"crypto/elliptic"
	"errors"
	"fmt"
	"github.com/nspcc-dev/neo-go/pkg/core/transaction"
	"github.com/nspcc-dev/neo-go/pkg/crypto/keys"
	"github.com/nspcc-dev/neo-go/pkg/neorpc/result"
	"github.com/nspcc-dev/neo-go/pkg/rpcclient/nep17"
	"github.com/nspcc-dev/neo-go/pkg/rpcclient/unwrap"
	"github.com/nspcc-dev/neo-go/pkg/util"
	"github.com/nspcc-dev/neo-go/pkg/vm/stackitem"
	"math/big"
	"unicode/utf8"
)

// BalanceAccount is a contract-specific balance.Account type used by its methods.
type BalanceAccount struct {
	Balance *big.Int
	Until   *big.Int
	Parent  []byte
}

// BalanceToken is a contract-specific balance.Token type used by its methods.
type BalanceToken struct {
	Symbol         string
	Decimals       *big.Int
	CirculationKey string
}

// CommonBallot is a contract-specific common.Ballot type used by its methods.
type CommonBallot struct {
	ID     []byte
	Voters keys.PublicKeys
	Height *big.Int
}

// LockEvent represents "Lock" event emitted by the contract.
type LockEvent struct {
	TxID   []byte
	From   util.Uint160
	To     util.Uint160
	Amount *big.Int
	Until  *big.Int
}

// TransferXEvent represents "TransferX" event emitted by the contract.
type TransferXEvent struct {
	From    util.Uint160
	To      util.Uint160
	Amount  *big.Int
	Details []byte
}

// Invoker is used by ContractReader to call various safe methods.
type Invoker interface {
	nep17.Invoker
}

// Actor is used by Contract to call state-changing methods.
type Actor interface {
	Invoker

	nep17.Actor

	MakeCall(contract util.Uint160, method string, params ...any) (*transaction.Transaction, error)
	MakeRun(script []byte) (*transaction.Transaction, error)
	MakeUnsignedCall(contract util.Uint160, method string, attrs []transaction.Attribute, params ...any) (*transaction.Transaction, error)
	MakeUnsignedRun(script []byte, attrs []transaction.Attribute) (*transaction.Transaction, error)
	SendCall(contract util.Uint160, method string, params ...any) (util.Uint256, uint32, error)
	SendRun(script []byte) (util.Uint256, uint32, error)
}

// ContractReader implements safe contract methods.
type ContractReader struct {
	nep17.TokenReader
	invoker Invoker
	hash    util.Uint160
}

// Contract implements all contract methods.
type Contract struct {
	ContractReader
	nep17.TokenWriter
	actor Actor
	hash  util.Uint160
}

// NewReader creates an instance of ContractReader using provided contract hash and the given Invoker.
func NewReader(invoker Invoker, hash util.Uint160) *ContractReader {
	return &ContractReader{*nep17.NewReader(invoker, hash), invoker, hash}
}

// New creates an instance of Contract using provided contract hash and the given Actor.
func New(actor Actor, hash util.Uint160) *Contract {
	var nep17t = nep17.New(actor, hash)
	return &Contract{ContractReader{nep17t.TokenReader, actor, hash}, nep17t.TokenWriter, actor, hash}
}

// Version invokes `version` method of contract.
func (c *ContractReader) Version() (*big.Int, error) {
	return unwrap.BigInt(c.invoker.Call(c.hash, "version"))
}

// Burn creates a transaction invoking `burn` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) Burn(from util.Uint160, amount *big.Int, txDetails []byte) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "burn", from, amount, txDetails)
}

// BurnTransaction creates a transaction invoking `burn` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) BurnTransaction(from util.Uint160, amount *big.Int, txDetails []byte) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "burn", from, amount, txDetails)
}

// BurnUnsigned creates a transaction invoking `burn` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) BurnUnsigned(from util.Uint160, amount *big.Int, txDetails []byte) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "burn", nil, from, amount, txDetails)
}

// Lock creates a transaction invoking `lock` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) Lock(txDetails []byte, from util.Uint160, to util.Uint160, amount *big.Int, until *big.Int) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "lock", txDetails, from, to, amount, until)
}

// LockTransaction creates a transaction invoking `lock` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) LockTransaction(txDetails []byte, from util.Uint160, to util.Uint160, amount *big.Int, until *big.Int) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "lock", txDetails, from, to, amount, until)
}

// LockUnsigned creates a transaction invoking `lock` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) LockUnsigned(txDetails []byte, from util.Uint160, to util.Uint160, amount *big.Int, until *big.Int) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "lock", nil, txDetails, from, to, amount, until)
}

// Mint creates a transaction invoking `mint` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) Mint(to util.Uint160, amount *big.Int, txDetails []byte) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "mint", to, amount, txDetails)
}

// MintTransaction creates a transaction invoking `mint` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) MintTransaction(to util.Uint160, amount *big.Int, txDetails []byte) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "mint", to, amount, txDetails)
}

// MintUnsigned creates a transaction invoking `mint` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) MintUnsigned(to util.Uint160, amount *big.Int, txDetails []byte) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "mint", nil, to, amount, txDetails)
}

// NewEpoch creates a transaction invoking `newEpoch` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) NewEpoch(epochNum *big.Int) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "newEpoch", epochNum)
}

// NewEpochTransaction creates a transaction invoking `newEpoch` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) NewEpochTransaction(epochNum *big.Int) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "newEpoch", epochNum)
}

// NewEpochUnsigned creates a transaction invoking `newEpoch` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) NewEpochUnsigned(epochNum *big.Int) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "newEpoch", nil, epochNum)
}

// TransferX creates a transaction invoking `transferX` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) TransferX(from util.Uint160, to util.Uint160, amount *big.Int, details []byte) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "transferX", from, to, amount, details)
}

// TransferXTransaction creates a transaction invoking `transferX` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) TransferXTransaction(from util.Uint160, to util.Uint160, amount *big.Int, details []byte) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "transferX", from, to, amount, details)
}

// TransferXUnsigned creates a transaction invoking `transferX` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) TransferXUnsigned(from util.Uint160, to util.Uint160, amount *big.Int, details []byte) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "transferX", nil, from, to, amount, details)
}

// Update creates a transaction invoking `update` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) Update(script []byte, manifest []byte, data any) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "update", script, manifest, data)
}

// UpdateTransaction creates a transaction invoking `update` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) UpdateTransaction(script []byte, manifest []byte, data any) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "update", script, manifest, data)
}

// UpdateUnsigned creates a transaction invoking `update` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) UpdateUnsigned(script []byte, manifest []byte, data any) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "update", nil, script, manifest, data)
}

// itemToBalanceAccount converts stack item into *BalanceAccount.
func itemToBalanceAccount(item stackitem.Item, err error) (*BalanceAccount, error) {
	if err != nil {
		return nil, err
	}
	var res = new(BalanceAccount)
	err = res.FromStackItem(item)
	return res, err
}

// FromStackItem retrieves fields of BalanceAccount from the given
// [stackitem.Item] or returns an error if it's not possible to do to so.
func (res *BalanceAccount) FromStackItem(item stackitem.Item) error {
	arr, ok := item.Value().([]stackitem.Item)
	if !ok {
		return errors.New("not an array")
	}
	if len(arr) != 3 {
		return errors.New("wrong number of structure elements")
	}

	var (
		index = -1
		err   error
	)
	index++
	res.Balance, err = arr[index].TryInteger()
	if err != nil {
		return fmt.Errorf("field Balance: %w", err)
	}

	index++
	res.Until, err = arr[index].TryInteger()
	if err != nil {
		return fmt.Errorf("field Until: %w", err)
	}

	index++
	res.Parent, err = arr[index].TryBytes()
	if err != nil {
		return fmt.Errorf("field Parent: %w", err)
	}

	return nil
}

// itemToBalanceToken converts stack item into *BalanceToken.
func itemToBalanceToken(item stackitem.Item, err error) (*BalanceToken, error) {
	if err != nil {
		return nil, err
	}
	var res = new(BalanceToken)
	err = res.FromStackItem(item)
	return res, err
}

// FromStackItem retrieves fields of BalanceToken from the given
// [stackitem.Item] or returns an error if it's not possible to do to so.
func (res *BalanceToken) FromStackItem(item stackitem.Item) error {
	arr, ok := item.Value().([]stackitem.Item)
	if !ok {
		return errors.New("not an array")
	}
	if len(arr) != 3 {
		return errors.New("wrong number of structure elements")
	}

	var (
		index = -1
		err   error
	)
	index++
	res.Symbol, err = func(item stackitem.Item) (string, error) {
		b, err := item.TryBytes()
		if err != nil {
			return "", err
		}
		if !utf8.Valid(b) {
			return "", errors.New("not a UTF-8 string")
		}
		return string(b), nil
	}(arr[index])
	if err != nil {
		return fmt.Errorf("field Symbol: %w", err)
	}

	index++
	res.Decimals, err = arr[index].TryInteger()
	if err != nil {
		return fmt.Errorf("field Decimals: %w", err)
	}

	index++
	res.CirculationKey, err = func(item stackitem.Item) (string, error) {
		b, err := item.TryBytes()
		if err != nil {
			return "", err
		}
		if !utf8.Valid(b) {
			return "", errors.New("not a UTF-8 string")
		}
		return string(b), nil
	}(arr[index])
	if err != nil {
		return fmt.Errorf("field CirculationKey: %w", err)
	}

	return nil
}

// itemToCommonBallot converts stack item into *CommonBallot.
func itemToCommonBallot(item stackitem.Item, err error) (*CommonBallot, error) {
	if err != nil {
		return nil, err
	}
	var res = new(CommonBallot)
	err = res.FromStackItem(item)
	return res, err
}

// FromStackItem retrieves fields of CommonBallot from the given
// [stackitem.Item] or returns an error if it's not possible to do to so.
func (res *CommonBallot) FromStackItem(item stackitem.Item) error {
	arr, ok := item.Value().([]stackitem.Item)
	if !ok {
		return errors.New("not an array")
	}
	if len(arr) != 3 {
		return errors.New("wrong number of structure elements")
	}

	var (
		index = -1
		err   error
	)
	index++
	res.ID, err = arr[index].TryBytes()
	if err != nil {
		return fmt.Errorf("field ID: %w", err)
	}

	index++
	res.Voters, err = func(item stackitem.Item) (keys.PublicKeys, error) {
		arr, ok := item.Value().([]stackitem.Item)
		if !ok {
			return nil, errors.New("not an array")
		}
		res := make(keys.PublicKeys, len(arr))
		for i := range res {
			res[i], err = func(item stackitem.Item) (*keys.PublicKey, error) {
				b, err := item.TryBytes()
				if err != nil {
					return nil, err
				}
				k, err := keys.NewPublicKeyFromBytes(b, elliptic.P256())
				if err != nil {
					return nil, err
				}
				return k, nil
			}(arr[i])
			if err != nil {
				return nil, fmt.Errorf("item %d: %w", i, err)
			}
		}
		return res, nil
	}(arr[index])
	if err != nil {
		return fmt.Errorf("field Voters: %w", err)
	}

	index++
	res.Height, err = arr[index].TryInteger()
	if err != nil {
		return fmt.Errorf("field Height: %w", err)
	}

	return nil
}

// LockEventsFromApplicationLog retrieves a set of all emitted events
// with "Lock" name from the provided [result.ApplicationLog].
func LockEventsFromApplicationLog(log *result.ApplicationLog) ([]*LockEvent, error) {
	if log == nil {
		return nil, errors.New("nil application log")
	}

	var res []*LockEvent
	for i, ex := range log.Executions {
		for j, e := range ex.Events {
			if e.Name != "Lock" {
				continue
			}
			event := new(LockEvent)
			err := event.FromStackItem(e.Item)
			if err != nil {
				return nil, fmt.Errorf("failed to deserialize LockEvent from stackitem (execution #%d, event #%d): %w", i, j, err)
			}
			res = append(res, event)
		}
	}

	return res, nil
}

// FromStackItem converts provided [stackitem.Array] to LockEvent or
// returns an error if it's not possible to do to so.
func (e *LockEvent) FromStackItem(item *stackitem.Array) error {
	if item == nil {
		return errors.New("nil item")
	}
	arr, ok := item.Value().([]stackitem.Item)
	if !ok {
		return errors.New("not an array")
	}
	if len(arr) != 5 {
		return errors.New("wrong number of structure elements")
	}

	var (
		index = -1
		err   error
	)
	index++
	e.TxID, err = arr[index].TryBytes()
	if err != nil {
		return fmt.Errorf("field TxID: %w", err)
	}

	index++
	e.From, err = func(item stackitem.Item) (util.Uint160, error) {
		b, err := item.TryBytes()
		if err != nil {
			return util.Uint160{}, err
		}
		u, err := util.Uint160DecodeBytesBE(b)
		if err != nil {
			return util.Uint160{}, err
		}
		return u, nil
	}(arr[index])
	if err != nil {
		return fmt.Errorf("field From: %w", err)
	}

	index++
	e.To, err = func(item stackitem.Item) (util.Uint160, error) {
		b, err := item.TryBytes()
		if err != nil {
			return util.Uint160{}, err
		}
		u, err := util.Uint160DecodeBytesBE(b)
		if err != nil {
			return util.Uint160{}, err
		}
		return u, nil
	}(arr[index])
	if err != nil {
		return fmt.Errorf("field To: %w", err)
	}

	index++
	e.Amount, err = arr[index].TryInteger()
	if err != nil {
		return fmt.Errorf("field Amount: %w", err)
	}

	index++
	e.Until, err = arr[index].TryInteger()
	if err != nil {
		return fmt.Errorf("field Until: %w", err)
	}

	return nil
}

// TransferXEventsFromApplicationLog retrieves a set of all emitted events
// with "TransferX" name from the provided [result.ApplicationLog].
func TransferXEventsFromApplicationLog(log *result.ApplicationLog) ([]*TransferXEvent, error) {
	if log == nil {
		return nil, errors.New("nil application log")
	}

	var res []*TransferXEvent
	for i, ex := range log.Executions {
		for j, e := range ex.Events {
			if e.Name != "TransferX" {
				continue
			}
			event := new(TransferXEvent)
			err := event.FromStackItem(e.Item)
			if err != nil {
				return nil, fmt.Errorf("failed to deserialize TransferXEvent from stackitem (execution #%d, event #%d): %w", i, j, err)
			}
			res = append(res, event)
		}
	}

	return res, nil
}

// FromStackItem converts provided [stackitem.Array] to TransferXEvent or
// returns an error if it's not possible to do to so.
func (e *TransferXEvent) FromStackItem(item *stackitem.Array) error {
	if item == nil {
		return errors.New("nil item")
	}
	arr, ok := item.Value().([]stackitem.Item)
	if !ok {
		return errors.New("not an array")
	}
	if len(arr) != 4 {
		return errors.New("wrong number of structure elements")
	}

	var (
		index = -1
		err   error
	)
	index++
	e.From, err = func(item stackitem.Item) (util.Uint160, error) {
		b, err := item.TryBytes()
		if err != nil {
			return util.Uint160{}, err
		}
		u, err := util.Uint160DecodeBytesBE(b)
		if err != nil {
			return util.Uint160{}, err
		}
		return u, nil
	}(arr[index])
	if err != nil {
		return fmt.Errorf("field From: %w", err)
	}

	index++
	e.To, err = func(item stackitem.Item) (util.Uint160, error) {
		b, err := item.TryBytes()
		if err != nil {
			return util.Uint160{}, err
		}
		u, err := util.Uint160DecodeBytesBE(b)
		if err != nil {
			return util.Uint160{}, err
		}
		return u, nil
	}(arr[index])
	if err != nil {
		return fmt.Errorf("field To: %w", err)
	}

	index++
	e.Amount, err = arr[index].TryInteger()
	if err != nil {
		return fmt.Errorf("field Amount: %w", err)
	}

	index++
	e.Details, err = arr[index].TryBytes()
	if err != nil {
		return fmt.Errorf("field Details: %w", err)
	}

	return nil
}
