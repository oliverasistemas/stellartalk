package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"stellartalk/x/stellartalk/types"
)

var _ = strconv.Itoa(0)

func CmdMyReceivedChats() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "my-received-chats",
		Short: "Query my received chats",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryMyReceivedChatsRequest{
				Pagination: pageReq,
				Receiver:   clientCtx.GetFromAddress().String(),
			}

			res, err := queryClient.MyReceivedChats(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
