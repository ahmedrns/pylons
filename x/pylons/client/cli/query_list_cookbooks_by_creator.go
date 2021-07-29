package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/Pylons-tech/pylons/x/pylons/types"
)

var _ = strconv.Itoa(0)

func CmdListCookbooksByCreator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-cookbooks [creator]",
		Short: "List cookbooks by creator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			reqCreator := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListCookbooksByCreatorRequest{
				Creator: reqCreator,
			}

			res, err := queryClient.ListCookbooksByCreator(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
