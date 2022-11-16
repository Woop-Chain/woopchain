package root

import (
	"fmt"
	"os"

	"github.com/Woop-Chain/woopchain/command/backup"
	"github.com/Woop-Chain/woopchain/command/genesis"
	"github.com/Woop-Chain/woopchain/command/helper"
	"github.com/Woop-Chain/woopchain/command/ibft"
	"github.com/Woop-Chain/woopchain/command/license"
	"github.com/Woop-Chain/woopchain/command/loadbot"
	"github.com/Woop-Chain/woopchain/command/monitor"
	"github.com/Woop-Chain/woopchain/command/peers"
	"github.com/Woop-Chain/woopchain/command/secrets"
	"github.com/Woop-Chain/woopchain/command/server"
	"github.com/Woop-Chain/woopchain/command/status"
	"github.com/Woop-Chain/woopchain/command/txpool"
	"github.com/Woop-Chain/woopchain/command/version"
	"github.com/Woop-Chain/woopchain/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Woop Chain is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		loadbot.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
