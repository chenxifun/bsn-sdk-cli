package request

import (
	"github.com/chenxifun/bsn-sdk-cli/request/fabric"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	var cobraCommand = &cobra.Command{
		Use:   "request",
		Short: "调用网关发起交易",
		Long:  `调用网关发起交易，可以是fabric或者fisco的一种`,
	}
	cobraCommand.AddCommand(fabric.Cmd())
	return cobraCommand
}
