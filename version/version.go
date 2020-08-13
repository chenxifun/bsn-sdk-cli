package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	return cobraCommand
}

var cobraCommand = &cobra.Command{
	Use:   "version",
	Short: "输出版本",
	Long:  `输出版本`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("不需要参数")
		}
		// Parsing of the command line is done so silence cmd usage
		cmd.SilenceUsage = true

		fmt.Println("1.0.0")
		return nil
	},
}
