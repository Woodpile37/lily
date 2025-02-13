// Code generated by: `make actors-gen`. DO NOT EDIT.

package market

import (
	"unicode/utf8"

	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
	cbg "github.com/whyrusleeping/cbor-gen"

	market8 "github.com/filecoin-project/go-state-types/builtin/v8/market"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"

	builtin5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"

	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"

	builtin7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"

	builtintypes8 "github.com/filecoin-project/go-state-types/builtin"

	lotusactors "github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/lily/chain/actors"
)

var (
	Address = builtintypes8.StorageMarketActorAddr
	Methods = builtintypes8.MethodsMarket
)

func Load(store adt.Store, act *types.Actor) (State, error) {
	if name, av, ok := lotusactors.GetActorMetaByCode(act.Code); ok {
		if name != actors.MarketKey {
			return nil, xerrors.Errorf("actor code is not market: %s", name)
		}

		switch actors.Version(av) {

		case actors.Version8:
			return load8(store, act.Head)

		}
	}

	switch act.Code {

	case builtin0.StorageMarketActorCodeID:
		return load0(store, act.Head)

	case builtin2.StorageMarketActorCodeID:
		return load2(store, act.Head)

	case builtin3.StorageMarketActorCodeID:
		return load3(store, act.Head)

	case builtin4.StorageMarketActorCodeID:
		return load4(store, act.Head)

	case builtin5.StorageMarketActorCodeID:
		return load5(store, act.Head)

	case builtin6.StorageMarketActorCodeID:
		return load6(store, act.Head)

	case builtin7.StorageMarketActorCodeID:
		return load7(store, act.Head)

	}

	return nil, xerrors.Errorf("unknown actor code %s", act.Code)
}

type State interface {
	cbor.Marshaler

	Code() cid.Cid
	ActorKey() string
	ActorVersion() actors.Version

	StatesChanged(State) (bool, error)
	States() (DealStates, error)
	ProposalsChanged(State) (bool, error)
	Proposals() (DealProposals, error)

	DealProposalsAmtBitwidth() int
	DealStatesAmtBitwidth() int
}

type DealStates interface {
	ForEach(cb func(id abi.DealID, ds DealState) error) error
	Get(id abi.DealID) (*DealState, bool, error)

	array() adt.Array
	decode(*cbg.Deferred) (*DealState, error)
}

type DealProposals interface {
	ForEach(cb func(id abi.DealID, dp market8.DealProposal) error) error
	Get(id abi.DealID) (*market8.DealProposal, bool, error)

	array() adt.Array
	decode(*cbg.Deferred) (*market8.DealProposal, error)
}

type DealProposal = market8.DealProposal

type DealState = market8.DealState

type DealStateChanges struct {
	Added    []DealIDState
	Modified []DealStateChange
	Removed  []DealIDState
}

type DealIDState struct {
	ID   abi.DealID
	Deal DealState
}

// DealStateChange is a change in deal state from -> to
type DealStateChange struct {
	ID   abi.DealID
	From *DealState
	To   *DealState
}

type DealProposalChanges struct {
	Added   []ProposalIDState
	Removed []ProposalIDState
}

type ProposalIDState struct {
	ID       abi.DealID
	Proposal market8.DealProposal
}

func labelFromGoString(s string) (market8.DealLabel, error) {
	if utf8.ValidString(s) {
		return market8.NewLabelFromString(s)
	} else {
		return market8.NewLabelFromBytes([]byte(s))
	}
}

func AllCodes() []cid.Cid {
	return []cid.Cid{
		(&state0{}).Code(),
		(&state2{}).Code(),
		(&state3{}).Code(),
		(&state4{}).Code(),
		(&state5{}).Code(),
		(&state6{}).Code(),
		(&state7{}).Code(),
		(&state8{}).Code(),
	}
}

func VersionCodes() map[actors.Version]cid.Cid {
	return map[actors.Version]cid.Cid{
		actors.Version0: (&state0{}).Code(),
		actors.Version2: (&state2{}).Code(),
		actors.Version3: (&state3{}).Code(),
		actors.Version4: (&state4{}).Code(),
		actors.Version5: (&state5{}).Code(),
		actors.Version6: (&state6{}).Code(),
		actors.Version7: (&state7{}).Code(),
		actors.Version8: (&state8{}).Code(),
	}
}
