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

func CmdMintNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-nft [id] [denom-id] [name] [data] [recipient]",
		Short: "Broadcast message mint-nft",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId := args[0]
			argDenomId := args[1]
			argName := args[2]
			argData := args[3]
			argUri := args[4]
			argRecipient := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintNft(
				argId,
				argDenomId,
				argName,
				argUri,
				argData,
				clientCtx.GetFromAddress().String(),
				argRecipient,
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
