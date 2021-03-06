package cryptonft

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/devashishdxt/crypto-nft/testutil/sample"
	cryptonftsimulation "github.com/devashishdxt/crypto-nft/x/cryptonft/simulation"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cryptonftsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgNewClass = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgNewClass int = 100

	opWeightMsgMintNft = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintNft int = 100

	opWeightMsgBurnNft = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBurnNft int = 100

	opWeightMsgUpdateNft = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateNft int = 100

	opWeightMsgUpdateClass = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateClass int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cryptonftGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cryptonftGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgNewClass int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgNewClass, &weightMsgNewClass, nil,
		func(_ *rand.Rand) {
			weightMsgNewClass = defaultWeightMsgNewClass
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgNewClass,
		cryptonftsimulation.SimulateMsgNewClass(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMintNft int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintNft, &weightMsgMintNft, nil,
		func(_ *rand.Rand) {
			weightMsgMintNft = defaultWeightMsgMintNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintNft,
		cryptonftsimulation.SimulateMsgMintNFT(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBurnNft int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBurnNft, &weightMsgBurnNft, nil,
		func(_ *rand.Rand) {
			weightMsgBurnNft = defaultWeightMsgBurnNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurnNft,
		cryptonftsimulation.SimulateMsgBurnNft(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateNft int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateNft, &weightMsgUpdateNft, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateNft = defaultWeightMsgUpdateNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateNft,
		cryptonftsimulation.SimulateMsgUpdateNFT(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateClass int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateClass, &weightMsgUpdateClass, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClass = defaultWeightMsgUpdateClass
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateClass,
		cryptonftsimulation.SimulateMsgUpdateClass(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
