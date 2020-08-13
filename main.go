package main

import (
	"github.com/chenxifun/bsn-sdk-cli/request"
	"github.com/chenxifun/bsn-sdk-cli/secret"
	"github.com/chenxifun/bsn-sdk-cli/version"
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use: "bsnCli"}

func main() {

	mainCmd.AddCommand(version.Cmd())
	mainCmd.AddCommand(secret.Cmd())
	mainCmd.AddCommand(request.Cmd())
	mainCmd.Execute()
}
