package main

import (
	"github.com/chenxifun/bsn-sdk-cli/random"
	"github.com/chenxifun/bsn-sdk-cli/request"
	"github.com/chenxifun/bsn-sdk-cli/secret"
	"github.com/chenxifun/bsn-sdk-cli/version"
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use: "bsn-sdk-cli"}

func main() {

	mainCmd.AddCommand(version.Cmd())
	mainCmd.AddCommand(secret.Cmd())
	mainCmd.AddCommand(random.Cmd())
	mainCmd.AddCommand(request.Cmd())
	mainCmd.Execute()
}
