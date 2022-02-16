package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBurnNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-nft [class-id] [nft-id]",
		Short: "Broadcast message burn-nft",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argClassId := args[0]
			argNftId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnNFT(
				clientCtx.GetFromAddress().String(),
				argClassId,
				argNftId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
