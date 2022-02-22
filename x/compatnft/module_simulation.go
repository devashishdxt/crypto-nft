package compatnft

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/devashishdxt/crypto-nft/testutil/sample"
	compatnftsimulation "github.com/devashishdxt/crypto-nft/x/compatnft/simulation"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = compatnftsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgIssueDenom = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgIssueDenom int = 100

	opWeightMsgMintNFT = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintNFT int = 100

	opWeightMsgEditNFT = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEditNFT int = 100

	opWeightMsgTransferNFT = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferNFT int = 100

	opWeightMsgBurnNFT = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBurnNFT int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	compatnftGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&compatnftGenesis)
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

	var weightMsgIssueDenom int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgIssueDenom, &weightMsgIssueDenom, nil,
		func(_ *rand.Rand) {
			weightMsgIssueDenom = defaultWeightMsgIssueDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgIssueDenom,
		compatnftsimulation.SimulateMsgIssueDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMintNFT int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintNFT, &weightMsgMintNFT, nil,
		func(_ *rand.Rand) {
			weightMsgMintNFT = defaultWeightMsgMintNFT
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintNFT,
		compatnftsimulation.SimulateMsgMintNFT(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEditNFT int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEditNFT, &weightMsgEditNFT, nil,
		func(_ *rand.Rand) {
			weightMsgEditNFT = defaultWeightMsgEditNFT
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEditNFT,
		compatnftsimulation.SimulateMsgEditNFT(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferNFT int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransferNFT, &weightMsgTransferNFT, nil,
		func(_ *rand.Rand) {
			weightMsgTransferNFT = defaultWeightMsgTransferNFT
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferNFT,
		compatnftsimulation.SimulateMsgTransferNFT(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBurnNFT int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBurnNFT, &weightMsgBurnNFT, nil,
		func(_ *rand.Rand) {
			weightMsgBurnNFT = defaultWeightMsgBurnNFT
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurnNFT,
		compatnftsimulation.SimulateMsgBurnNFT(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
