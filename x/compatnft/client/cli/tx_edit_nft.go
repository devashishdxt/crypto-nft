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

func CmdEditNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-nft [id] [denom-id] [name] [uri] [data]",
		Short: "Broadcast message edit-NFT",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId := args[0]
			argDenomId := args[1]
			argName := args[2]
			argUri := args[3]
			argData := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditNFT(
				argId,
				argDenomId,
				argName,
				argUri,
				argData,
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
