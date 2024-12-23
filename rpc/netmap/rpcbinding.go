// Package netmap contains RPC wrappers for NeoFS Netmap contract.
//
// Code generated by neo-go contract generate-rpcwrapper --manifest <file.json> --out <file.go> [--hash <hash>] [--config <config>]; DO NOT EDIT.
package netmap

import (
	"crypto/elliptic"
	"errors"
	"fmt"
	"github.com/epicchainlabs/epicchain-go/pkg/core/transaction"
	"github.com/epicchainlabs/epicchain-go/pkg/crypto/keys"
	"github.com/epicchainlabs/epicchain-go/pkg/neorpc/result"
	"github.com/epicchainlabs/epicchain-go/pkg/rpcclient/unwrap"
	"github.com/epicchainlabs/epicchain-go/pkg/util"
	"github.com/epicchainlabs/epicchain-go/pkg/vm/stackitem"
	"math/big"
)

// CommonBallot is a contract-specific common.Ballot type used by its methods.
type CommonBallot struct {
	ID     []byte
	Voters keys.PublicKeys
	Height *big.Int
}

// CommonIRNode is a contract-specific common.IRNode type used by its methods.
type CommonIRNode struct {
	PublicKey *keys.PublicKey
}

// NetmapNode is a contract-specific netmap.Node type used by its methods.
type NetmapNode struct {
	BLOB  []byte
	State *big.Int
}

// Netmaprecord is a contract-specific netmap.record type used by its methods.
type Netmaprecord struct {
	Key []byte
	Val []byte
}

// AddPeerSuccessEvent represents "AddPeerSuccess" event emitted by the contract.
type AddPeerSuccessEvent struct {
	PublicKey *keys.PublicKey
}

// UpdateStateSuccessEvent represents "UpdateStateSuccess" event emitted by the contract.
type UpdateStateSuccessEvent struct {
	PublicKey *keys.PublicKey
	State     *big.Int
}

// NewEpochEvent represents "NewEpoch" event emitted by the contract.
type NewEpochEvent struct {
	Epoch *big.Int
}

// NewEpochSubscriptionEvent represents "NewEpochSubscription" event emitted by the contract.
type NewEpochSubscriptionEvent struct {
	Contract util.Uint160
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

// Config invokes `config` method of contract.
func (c *ContractReader) Config(key []byte) (any, error) {
	return func(item stackitem.Item, err error) (any, error) {
		if err != nil {
			return nil, err
		}
		return item.Value(), error(nil)
	}(unwrap.Item(c.invoker.Call(c.hash, "config", key)))
}

// Epoch invokes `epoch` method of contract.
func (c *ContractReader) Epoch() (*big.Int, error) {
	return unwrap.BigInt(c.invoker.Call(c.hash, "epoch"))
}

// InnerRingList invokes `innerRingList` method of contract.
func (c *ContractReader) InnerRingList() ([]*CommonIRNode, error) {
	return func(item stackitem.Item, err error) ([]*CommonIRNode, error) {
		if err != nil {
			return nil, err
		}
		return func(item stackitem.Item) ([]*CommonIRNode, error) {
			arr, ok := item.Value().([]stackitem.Item)
			if !ok {
				return nil, errors.New("not an array")
			}
			res := make([]*CommonIRNode, len(arr))
			for i := range res {
				res[i], err = itemToCommonIRNode(arr[i], nil)
				if err != nil {
					return nil, fmt.Errorf("item %d: %w", i, err)
				}
			}
			return res, nil
		}(item)
	}(unwrap.Item(c.invoker.Call(c.hash, "innerRingList")))
}

// ListConfig invokes `listConfig` method of contract.
func (c *ContractReader) ListConfig() ([]*Netmaprecord, error) {
	return func(item stackitem.Item, err error) ([]*Netmaprecord, error) {
		if err != nil {
			return nil, err
		}
		return func(item stackitem.Item) ([]*Netmaprecord, error) {
			arr, ok := item.Value().([]stackitem.Item)
			if !ok {
				return nil, errors.New("not an array")
			}
			res := make([]*Netmaprecord, len(arr))
			for i := range res {
				res[i], err = itemToNetmaprecord(arr[i], nil)
				if err != nil {
					return nil, fmt.Errorf("item %d: %w", i, err)
				}
			}
			return res, nil
		}(item)
	}(unwrap.Item(c.invoker.Call(c.hash, "listConfig")))
}

// Netmap invokes `netmap` method of contract.
func (c *ContractReader) Netmap() ([]*NetmapNode, error) {
	return func(item stackitem.Item, err error) ([]*NetmapNode, error) {
		if err != nil {
			return nil, err
		}
		return func(item stackitem.Item) ([]*NetmapNode, error) {
			arr, ok := item.Value().([]stackitem.Item)
			if !ok {
				return nil, errors.New("not an array")
			}
			res := make([]*NetmapNode, len(arr))
			for i := range res {
				res[i], err = itemToNetmapNode(arr[i], nil)
				if err != nil {
					return nil, fmt.Errorf("item %d: %w", i, err)
				}
			}
			return res, nil
		}(item)
	}(unwrap.Item(c.invoker.Call(c.hash, "netmap")))
}

// NetmapCandidates invokes `netmapCandidates` method of contract.
func (c *ContractReader) NetmapCandidates() ([]*NetmapNode, error) {
	return func(item stackitem.Item, err error) ([]*NetmapNode, error) {
		if err != nil {
			return nil, err
		}
		return func(item stackitem.Item) ([]*NetmapNode, error) {
			arr, ok := item.Value().([]stackitem.Item)
			if !ok {
				return nil, errors.New("not an array")
			}
			res := make([]*NetmapNode, len(arr))
			for i := range res {
				res[i], err = itemToNetmapNode(arr[i], nil)
				if err != nil {
					return nil, fmt.Errorf("item %d: %w", i, err)
				}
			}
			return res, nil
		}(item)
	}(unwrap.Item(c.invoker.Call(c.hash, "netmapCandidates")))
}

// Snapshot invokes `snapshot` method of contract.
func (c *ContractReader) Snapshot(diff *big.Int) ([]*NetmapNode, error) {
	return func(item stackitem.Item, err error) ([]*NetmapNode, error) {
		if err != nil {
			return nil, err
		}
		return func(item stackitem.Item) ([]*NetmapNode, error) {
			arr, ok := item.Value().([]stackitem.Item)
			if !ok {
				return nil, errors.New("not an array")
			}
			res := make([]*NetmapNode, len(arr))
			for i := range res {
				res[i], err = itemToNetmapNode(arr[i], nil)
				if err != nil {
					return nil, fmt.Errorf("item %d: %w", i, err)
				}
			}
			return res, nil
		}(item)
	}(unwrap.Item(c.invoker.Call(c.hash, "snapshot", diff)))
}

// SnapshotByEpoch invokes `snapshotByEpoch` method of contract.
func (c *ContractReader) SnapshotByEpoch(epoch *big.Int) ([]*NetmapNode, error) {
	return func(item stackitem.Item, err error) ([]*NetmapNode, error) {
		if err != nil {
			return nil, err
		}
		return func(item stackitem.Item) ([]*NetmapNode, error) {
			arr, ok := item.Value().([]stackitem.Item)
			if !ok {
				return nil, errors.New("not an array")
			}
			res := make([]*NetmapNode, len(arr))
			for i := range res {
				res[i], err = itemToNetmapNode(arr[i], nil)
				if err != nil {
					return nil, fmt.Errorf("item %d: %w", i, err)
				}
			}
			return res, nil
		}(item)
	}(unwrap.Item(c.invoker.Call(c.hash, "snapshotByEpoch", epoch)))
}

// Version invokes `version` method of contract.
func (c *ContractReader) Version() (*big.Int, error) {
	return unwrap.BigInt(c.invoker.Call(c.hash, "version"))
}

// AddPeer creates a transaction invoking `addPeer` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) AddPeer(nodeInfo []byte) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "addPeer", nodeInfo)
}

// AddPeerTransaction creates a transaction invoking `addPeer` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) AddPeerTransaction(nodeInfo []byte) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "addPeer", nodeInfo)
}

// AddPeerUnsigned creates a transaction invoking `addPeer` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) AddPeerUnsigned(nodeInfo []byte) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "addPeer", nil, nodeInfo)
}

// AddPeerIR creates a transaction invoking `addPeerIR` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) AddPeerIR(nodeInfo []byte) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "addPeerIR", nodeInfo)
}

// AddPeerIRTransaction creates a transaction invoking `addPeerIR` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) AddPeerIRTransaction(nodeInfo []byte) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "addPeerIR", nodeInfo)
}

// AddPeerIRUnsigned creates a transaction invoking `addPeerIR` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) AddPeerIRUnsigned(nodeInfo []byte) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "addPeerIR", nil, nodeInfo)
}

// LastEpochBlock creates a transaction invoking `lastEpochBlock` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) LastEpochBlock() (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "lastEpochBlock")
}

// LastEpochBlockTransaction creates a transaction invoking `lastEpochBlock` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) LastEpochBlockTransaction() (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "lastEpochBlock")
}

// LastEpochBlockUnsigned creates a transaction invoking `lastEpochBlock` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) LastEpochBlockUnsigned() (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "lastEpochBlock", nil)
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

// SetConfig creates a transaction invoking `setConfig` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) SetConfig(id []byte, key []byte, val []byte) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "setConfig", id, key, val)
}

// SetConfigTransaction creates a transaction invoking `setConfig` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) SetConfigTransaction(id []byte, key []byte, val []byte) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "setConfig", id, key, val)
}

// SetConfigUnsigned creates a transaction invoking `setConfig` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) SetConfigUnsigned(id []byte, key []byte, val []byte) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "setConfig", nil, id, key, val)
}

// SubscribeForNewEpoch creates a transaction invoking `subscribeForNewEpoch` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) SubscribeForNewEpoch(contract util.Uint160) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "subscribeForNewEpoch", contract)
}

// SubscribeForNewEpochTransaction creates a transaction invoking `subscribeForNewEpoch` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) SubscribeForNewEpochTransaction(contract util.Uint160) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "subscribeForNewEpoch", contract)
}

// SubscribeForNewEpochUnsigned creates a transaction invoking `subscribeForNewEpoch` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) SubscribeForNewEpochUnsigned(contract util.Uint160) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "subscribeForNewEpoch", nil, contract)
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

// UpdateSnapshotCount creates a transaction invoking `updateSnapshotCount` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) UpdateSnapshotCount(count *big.Int) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "updateSnapshotCount", count)
}

// UpdateSnapshotCountTransaction creates a transaction invoking `updateSnapshotCount` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) UpdateSnapshotCountTransaction(count *big.Int) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "updateSnapshotCount", count)
}

// UpdateSnapshotCountUnsigned creates a transaction invoking `updateSnapshotCount` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) UpdateSnapshotCountUnsigned(count *big.Int) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "updateSnapshotCount", nil, count)
}

// UpdateState creates a transaction invoking `updateState` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) UpdateState(state *big.Int, publicKey *keys.PublicKey) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "updateState", state, publicKey)
}

// UpdateStateTransaction creates a transaction invoking `updateState` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) UpdateStateTransaction(state *big.Int, publicKey *keys.PublicKey) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "updateState", state, publicKey)
}

// UpdateStateUnsigned creates a transaction invoking `updateState` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) UpdateStateUnsigned(state *big.Int, publicKey *keys.PublicKey) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "updateState", nil, state, publicKey)
}

// UpdateStateIR creates a transaction invoking `updateStateIR` method of the contract.
// This transaction is signed and immediately sent to the network.
// The values returned are its hash, ValidUntilBlock value and error if any.
func (c *Contract) UpdateStateIR(state *big.Int, publicKey *keys.PublicKey) (util.Uint256, uint32, error) {
	return c.actor.SendCall(c.hash, "updateStateIR", state, publicKey)
}

// UpdateStateIRTransaction creates a transaction invoking `updateStateIR` method of the contract.
// This transaction is signed, but not sent to the network, instead it's
// returned to the caller.
func (c *Contract) UpdateStateIRTransaction(state *big.Int, publicKey *keys.PublicKey) (*transaction.Transaction, error) {
	return c.actor.MakeCall(c.hash, "updateStateIR", state, publicKey)
}

// UpdateStateIRUnsigned creates a transaction invoking `updateStateIR` method of the contract.
// This transaction is not signed, it's simply returned to the caller.
// Any fields of it that do not affect fees can be changed (ValidUntilBlock,
// Nonce), fee values (NetworkFee, SystemFee) can be increased as well.
func (c *Contract) UpdateStateIRUnsigned(state *big.Int, publicKey *keys.PublicKey) (*transaction.Transaction, error) {
	return c.actor.MakeUnsignedCall(c.hash, "updateStateIR", nil, state, publicKey)
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

// itemToCommonIRNode converts stack item into *CommonIRNode.
func itemToCommonIRNode(item stackitem.Item, err error) (*CommonIRNode, error) {
	if err != nil {
		return nil, err
	}
	var res = new(CommonIRNode)
	err = res.FromStackItem(item)
	return res, err
}

// FromStackItem retrieves fields of CommonIRNode from the given
// [stackitem.Item] or returns an error if it's not possible to do to so.
func (res *CommonIRNode) FromStackItem(item stackitem.Item) error {
	arr, ok := item.Value().([]stackitem.Item)
	if !ok {
		return errors.New("not an array")
	}
	if len(arr) != 1 {
		return errors.New("wrong number of structure elements")
	}

	var (
		index = -1
		err   error
	)
	index++
	res.PublicKey, err = func(item stackitem.Item) (*keys.PublicKey, error) {
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
		return fmt.Errorf("field PublicKey: %w", err)
	}

	return nil
}

// itemToNetmapNode converts stack item into *NetmapNode.
func itemToNetmapNode(item stackitem.Item, err error) (*NetmapNode, error) {
	if err != nil {
		return nil, err
	}
	var res = new(NetmapNode)
	err = res.FromStackItem(item)
	return res, err
}

// FromStackItem retrieves fields of NetmapNode from the given
// [stackitem.Item] or returns an error if it's not possible to do to so.
func (res *NetmapNode) FromStackItem(item stackitem.Item) error {
	arr, ok := item.Value().([]stackitem.Item)
	if !ok {
		return errors.New("not an array")
	}
	if len(arr) != 2 {
		return errors.New("wrong number of structure elements")
	}

	var (
		index = -1
		err   error
	)
	index++
	res.BLOB, err = arr[index].TryBytes()
	if err != nil {
		return fmt.Errorf("field BLOB: %w", err)
	}

	index++
	res.State, err = arr[index].TryInteger()
	if err != nil {
		return fmt.Errorf("field State: %w", err)
	}

	return nil
}

// itemToNetmaprecord converts stack item into *Netmaprecord.
func itemToNetmaprecord(item stackitem.Item, err error) (*Netmaprecord, error) {
	if err != nil {
		return nil, err
	}
	var res = new(Netmaprecord)
	err = res.FromStackItem(item)
	return res, err
}

// FromStackItem retrieves fields of Netmaprecord from the given
// [stackitem.Item] or returns an error if it's not possible to do to so.
func (res *Netmaprecord) FromStackItem(item stackitem.Item) error {
	arr, ok := item.Value().([]stackitem.Item)
	if !ok {
		return errors.New("not an array")
	}
	if len(arr) != 2 {
		return errors.New("wrong number of structure elements")
	}

	var (
		index = -1
		err   error
	)
	index++
	res.Key, err = arr[index].TryBytes()
	if err != nil {
		return fmt.Errorf("field Key: %w", err)
	}

	index++
	res.Val, err = arr[index].TryBytes()
	if err != nil {
		return fmt.Errorf("field Val: %w", err)
	}

	return nil
}

// AddPeerSuccessEventsFromApplicationLog retrieves a set of all emitted events
// with "AddPeerSuccess" name from the provided [result.ApplicationLog].
func AddPeerSuccessEventsFromApplicationLog(log *result.ApplicationLog) ([]*AddPeerSuccessEvent, error) {
	if log == nil {
		return nil, errors.New("nil application log")
	}

	var res []*AddPeerSuccessEvent
	for i, ex := range log.Executions {
		for j, e := range ex.Events {
			if e.Name != "AddPeerSuccess" {
				continue
			}
			event := new(AddPeerSuccessEvent)
			err := event.FromStackItem(e.Item)
			if err != nil {
				return nil, fmt.Errorf("failed to deserialize AddPeerSuccessEvent from stackitem (execution #%d, event #%d): %w", i, j, err)
			}
			res = append(res, event)
		}
	}

	return res, nil
}

// FromStackItem converts provided [stackitem.Array] to AddPeerSuccessEvent or
// returns an error if it's not possible to do to so.
func (e *AddPeerSuccessEvent) FromStackItem(item *stackitem.Array) error {
	if item == nil {
		return errors.New("nil item")
	}
	arr, ok := item.Value().([]stackitem.Item)
	if !ok {
		return errors.New("not an array")
	}
	if len(arr) != 1 {
		return errors.New("wrong number of structure elements")
	}

	var (
		index = -1
		err   error
	)
	index++
	e.PublicKey, err = func(item stackitem.Item) (*keys.PublicKey, error) {
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
		return fmt.Errorf("field PublicKey: %w", err)
	}

	return nil
}

// UpdateStateSuccessEventsFromApplicationLog retrieves a set of all emitted events
// with "UpdateStateSuccess" name from the provided [result.ApplicationLog].
func UpdateStateSuccessEventsFromApplicationLog(log *result.ApplicationLog) ([]*UpdateStateSuccessEvent, error) {
	if log == nil {
		return nil, errors.New("nil application log")
	}

	var res []*UpdateStateSuccessEvent
	for i, ex := range log.Executions {
		for j, e := range ex.Events {
			if e.Name != "UpdateStateSuccess" {
				continue
			}
			event := new(UpdateStateSuccessEvent)
			err := event.FromStackItem(e.Item)
			if err != nil {
				return nil, fmt.Errorf("failed to deserialize UpdateStateSuccessEvent from stackitem (execution #%d, event #%d): %w", i, j, err)
			}
			res = append(res, event)
		}
	}

	return res, nil
}

// FromStackItem converts provided [stackitem.Array] to UpdateStateSuccessEvent or
// returns an error if it's not possible to do to so.
func (e *UpdateStateSuccessEvent) FromStackItem(item *stackitem.Array) error {
	if item == nil {
		return errors.New("nil item")
	}
	arr, ok := item.Value().([]stackitem.Item)
	if !ok {
		return errors.New("not an array")
	}
	if len(arr) != 2 {
		return errors.New("wrong number of structure elements")
	}

	var (
		index = -1
		err   error
	)
	index++
	e.PublicKey, err = func(item stackitem.Item) (*keys.PublicKey, error) {
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
		return fmt.Errorf("field PublicKey: %w", err)
	}

	index++
	e.State, err = arr[index].TryInteger()
	if err != nil {
		return fmt.Errorf("field State: %w", err)
	}

	return nil
}

// NewEpochEventsFromApplicationLog retrieves a set of all emitted events
// with "NewEpoch" name from the provided [result.ApplicationLog].
func NewEpochEventsFromApplicationLog(log *result.ApplicationLog) ([]*NewEpochEvent, error) {
	if log == nil {
		return nil, errors.New("nil application log")
	}

	var res []*NewEpochEvent
	for i, ex := range log.Executions {
		for j, e := range ex.Events {
			if e.Name != "NewEpoch" {
				continue
			}
			event := new(NewEpochEvent)
			err := event.FromStackItem(e.Item)
			if err != nil {
				return nil, fmt.Errorf("failed to deserialize NewEpochEvent from stackitem (execution #%d, event #%d): %w", i, j, err)
			}
			res = append(res, event)
		}
	}

	return res, nil
}

// FromStackItem converts provided [stackitem.Array] to NewEpochEvent or
// returns an error if it's not possible to do to so.
func (e *NewEpochEvent) FromStackItem(item *stackitem.Array) error {
	if item == nil {
		return errors.New("nil item")
	}
	arr, ok := item.Value().([]stackitem.Item)
	if !ok {
		return errors.New("not an array")
	}
	if len(arr) != 1 {
		return errors.New("wrong number of structure elements")
	}

	var (
		index = -1
		err   error
	)
	index++
	e.Epoch, err = arr[index].TryInteger()
	if err != nil {
		return fmt.Errorf("field Epoch: %w", err)
	}

	return nil
}

// NewEpochSubscriptionEventsFromApplicationLog retrieves a set of all emitted events
// with "NewEpochSubscription" name from the provided [result.ApplicationLog].
func NewEpochSubscriptionEventsFromApplicationLog(log *result.ApplicationLog) ([]*NewEpochSubscriptionEvent, error) {
	if log == nil {
		return nil, errors.New("nil application log")
	}

	var res []*NewEpochSubscriptionEvent
	for i, ex := range log.Executions {
		for j, e := range ex.Events {
			if e.Name != "NewEpochSubscription" {
				continue
			}
			event := new(NewEpochSubscriptionEvent)
			err := event.FromStackItem(e.Item)
			if err != nil {
				return nil, fmt.Errorf("failed to deserialize NewEpochSubscriptionEvent from stackitem (execution #%d, event #%d): %w", i, j, err)
			}
			res = append(res, event)
		}
	}

	return res, nil
}

// FromStackItem converts provided [stackitem.Array] to NewEpochSubscriptionEvent or
// returns an error if it's not possible to do to so.
func (e *NewEpochSubscriptionEvent) FromStackItem(item *stackitem.Array) error {
	if item == nil {
		return errors.New("nil item")
	}
	arr, ok := item.Value().([]stackitem.Item)
	if !ok {
		return errors.New("not an array")
	}
	if len(arr) != 1 {
		return errors.New("wrong number of structure elements")
	}

	var (
		index = -1
		err   error
	)
	index++
	e.Contract, err = func(item stackitem.Item) (util.Uint160, error) {
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
		return fmt.Errorf("field Contract: %w", err)
	}

	return nil
}
