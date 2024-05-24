// Package audit contains RPC wrappers for NeoFS Audit contract.
//
// Code generated by neo-go contract generate-rpcwrapper --manifest <file.json> --out <file.go> [--hash <hash>] [--config <config>]; DO NOT EDIT.
package audit

import (
	"crypto/elliptic"
	"errors"
	"fmt"
	"github.com/nspcc-dev/neo-go/pkg/core/transaction"
	"github.com/nspcc-dev/neo-go/pkg/crypto/keys"
	"github.com/nspcc-dev/neo-go/pkg/neorpc/result"
	"github.com/nspcc-dev/neo-go/pkg/rpcclient/unwrap"
	"github.com/nspcc-dev/neo-go/pkg/util"
	"github.com/nspcc-dev/neo-go/pkg/vm/stackitem"
	"math/big"
)

// AuditAuditHeader is a contract-specific audit.AuditHeader type used by its methods.
type AuditAuditHeader struct {
	Epoch *big.Int
	CID   []byte
	From  *keys.PublicKey
}

// Invoker is used by ContractReader to call various safe methods.
type Invoker interface {
	Call(contract util.Uint160, operation string, params ...any) (*result.Invoke, error)
}

// Actor is used by Contract to call state-changing methods.
type Actor interface {
	Invoker

	MakeCall(contract util.Uint160, method string, params ...any) (*transaction.Transaction, error)
	MakeRun(script []byte) (*transaction.Transaction, error)
	MakeUnsignedCall(contract util.Uint160, method string, attrs []transaction.Attribute, params ...any) (*transaction.Transaction, error)
	MakeUnsignedRun(script []byte, attrs []transaction.Attribute) (*transaction.Transaction, error)
	SendCall(contract util.Uint160, method string, params ...any) (util.Uint256, uint32, error)
	SendRun(script []byte) (util.Uint256, uint32, error)
}

// ContractReader implements safe contract methods.
type ContractReader struct {
	invoker Invoker
	hash    util.Uint160
}

// Contract implements all contract methods.
type Contract struct {
	ContractReader
	actor Actor
	hash  util.Uint160
}

// NewReader creates an instance of ContractReader using provided contract hash and the given Invoker.
func NewReader(invoker Invoker, hash util.Uint160) *ContractReader {
	return &ContractReader{invoker, hash}
}

// New creates an instance of Contract using provided contract hash and the given Actor.
func New(actor Actor, hash util.Uint160) *Contract {
	return &Contract{ContractReader{actor, hash}, actor, hash}
}

// Get invokes `get` method of contract.
func (c *ContractReader) Get(id []byte) ([]byte, error) {
	return unwrap.Bytes(c.invoker.Call(c.hash, "get", id))
}

// List invokes `list` method of contract.
func (c *ContractReader) List() ([][]byte, error) {
	return unwrap.ArrayOfBytes(c.invoker.Call(c.hash, "list"))
}

// ListByCID invokes `listByCID` method of contract.
func (c *ContractReader) ListByCID(epoch *big.Int, cid []byte) ([][]byte, error) {
	return unwrap.ArrayOfBytes(c.invoker.Call(c.hash, "listByCID", epoch, cid))
}

// ListByEpoch invokes `listByEpoch` method of contract.
func (c *ContractReader) ListByEpoch(epoch *big.Int) ([][]byte, error) {
	return unwrap.ArrayOfBytes(c.invoker.Call(c.hash, "listByEpoch", epoch))
}

// ListByNode invokes `listByNode` method of contract.
func (c *ContractReader) ListByNode(epoch *big.Int, cid []byte, key *keys.PublicKey) ([][]byte, error) {
	return unwrap.ArrayOfBytes(c.invoker.Call(c.hash, "listByNode", epoch, cid, key))
}

// Version invokes `version` method of contract.
func (c *ContractReader) Version() (*big.Int, error) {
	return unwrap.BigInt(c.invoker.Call(c.hash, "version"))
}

// Put creates a transaction invoking `put` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) Put(rawAuditResult []byte) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "put", rawAuditResult)
}

// PutTransaction creates a transaction invoking `put` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) PutTransaction(rawAuditResult []byte) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "put", rawAuditResult)
}

// PutUnsigned creates a transaction invoking `put` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) PutUnsigned(rawAuditResult []byte) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "put", nil, rawAuditResult)
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

// itemToAuditAuditHeader converts stack item into *AuditAuditHeader.
func itemToAuditAuditHeader(item stackitem.Item, err error) (*AuditAuditHeader, error) {
	if err != nil {
		return nil, err
	}
	var res = new(AuditAuditHeader)
	err = res.FromStackItem(item)
	return res, err
}

// FromStackItem retrieves fields of AuditAuditHeader from the given
// [stackitem.Item] or returns an error if it's not possible to do to so.
func (res *AuditAuditHeader) FromStackItem(item stackitem.Item) error {
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
	res.Epoch, err = arr[index].TryInteger()
	if err != nil {
		return fmt.Errorf("field Epoch: %w", err)
	}

	index++
	res.CID, err = arr[index].TryBytes()
	if err != nil {
		return fmt.Errorf("field CID: %w", err)
	}

	index++
	res.From, err = func(item stackitem.Item) (*keys.PublicKey, error) {
		b, err := item.TryBytes()
		if err != nil {
			return nil, err
		}
		k, err := keys.NewPublicKeyFromBytes(b, elliptic.P256())
		if err != nil {
			return nil, err
		}
		return k, nil
	}(arr[index])
	if err != nil {
		return fmt.Errorf("field From: %w", err)
	}

	return nil
}
