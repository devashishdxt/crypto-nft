package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDenoms() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "denom [denom-id]",
		Short: "Query denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenomId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDenomRequest{

				DenomId: reqDenomId,
			}

			res, err := queryClient.Denom(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
