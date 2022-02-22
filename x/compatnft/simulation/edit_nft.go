package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/devashishdxt/crypto-nft/x/compatnft/keeper"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
)

func SimulateMsgEditNFT(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgEditNFT{
			Sender: simAccount.Address.String(),
		}

		// TODO: Handling the EditNFT simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "EditNFT simulation not implemented"), nil, nil
	}
}
