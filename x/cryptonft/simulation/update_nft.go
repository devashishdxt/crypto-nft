package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/keeper"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
)

func SimulateMsgUpdateNFT(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateNFT{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateNFT simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UpdateNFT simulation not implemented"), nil, nil
	}
}
