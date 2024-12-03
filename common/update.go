package common

import (
	"github.com/epicchainlabs/epicchain-go/pkg/interop/runtime"
)

// LegacyOwnerKey is storage key used to store contract owner.
const LegacyOwnerKey = "contractOwner"

// HasUpdateAccess returns true if contract can be updated.
func HasUpdateAccess() bool {
	return runtime.CheckWitness(CommitteeAddress())
}
