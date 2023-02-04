package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"stellartalk/x/stellartalk/types"
)

var _ = strconv.Itoa(0)

func CmdMySentChats() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "my-sent-chats",
		Short: "List all chats sent from my wallet",
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

			params := &types.QueryMySentChatsRequest{
				Pagination: pageReq,
				Sender:     clientCtx.GetFromAddress().String(),
			}

			res, err := queryClient.MySentChats(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
