package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBurnNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-nft [id] [denom-id]",
		Short: "Broadcast message burn-NFT",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId := args[0]
			argDenomId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnNFT(
				argId,
				argDenomId,
				clientCtx.GetFromAddress().String(),
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
