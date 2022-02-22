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

func CmdIssueDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue-denom [id] [name] [schema]",
		Short: "Broadcast message issue-denom",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId := args[0]
			argName := args[1]
			argSchema := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssueDenom(
				argId,
				argName,
				argSchema,
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
