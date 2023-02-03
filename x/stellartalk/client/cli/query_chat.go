package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"stellartalk/x/stellartalk/types"
)

const FlagSenderFilter = "sender"
const FlagRecipientFilter = "recipient"

func CmdListChat() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-chat",
		Short: "list all chat",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			sender, err := cmd.Flags().GetString(FlagSenderFilter)
			if err != nil {
				return err
			}

			recipient, err := cmd.Flags().GetString(FlagRecipientFilter)
			if err != nil {
				return err
			}

			params := &types.QueryAllChatRequest{
				Pagination: pageReq,
				Sender:     sender,
				Recipient:  recipient,
			}

			res, err := queryClient.ChatAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)
	cmd.Flags().StringP(FlagSenderFilter, "s", "", "Filter by sender")
	cmd.Flags().StringP(FlagRecipientFilter, "r", "", "Filter by recipient")

	return cmd
}

func CmdShowChat() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-chat [id]",
		Short: "shows a chat",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetChatRequest{
				Id: id,
			}

			res, err := queryClient.Chat(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
