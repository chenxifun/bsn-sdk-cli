package secret

import (
	"fmt"
	"github.com/chenxifun/bsn-sdk-cli/secret/k1"
	"github.com/chenxifun/bsn-sdk-cli/secret/r1"
	"github.com/chenxifun/bsn-sdk-cli/secret/sm"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	algo string
	data string
)
var flags *pflag.FlagSet

func init() {
	flags = &pflag.FlagSet{}
	flags.StringVarP(&algo, "algo", "a", "sm2", "秘钥算法类型,可选为【sm2、secp256r1、secp256k1】")
	flags.StringVarP(&data, "data", "d", "123456", "签名的数据")

}

func Cmd() *cobra.Command {
	var cobraCommand = &cobra.Command{
		Use:   "secret",
		Short: "生成秘钥",
		Long:  `生成对应算法的公私钥，并且得到一个对数据签名的base64结果`,
		Run: func(cmd *cobra.Command, args []string) {

			cmd.SilenceUsage = true

			secret()
		},
	}
	flagList := []string{
		"algo",
		"data",
	}
	attachFlags(cobraCommand, flagList)
	return cobraCommand
}

func secret() {

	if algo != "sm2" && algo != "secp256r1" && algo != "secp256k1" {
		fmt.Println("秘钥算法不正确，应当为【sm2、secp256r1、secp256k1】中的一种，实际为：", algo)
		return
	}

	fmt.Println("秘钥类型：", algo)
	fmt.Println("签名数据：", data)

	switch algo {
	case "sm2":
		sm.NewKey(data)
		break
	case "secp256r1":
		r1.NewKey(data)
		break
	case "secp256k1":
		k1.NewKey(data)
		break
	default:
		fmt.Println("秘钥算法不正确，应当为【sm、r1、k1】中的一种，实际为：", algo)
		return
	}
}

func attachFlags(cmd *cobra.Command, names []string) {
	cmdFlags := cmd.Flags()
	for _, name := range names {
		if flag := flags.Lookup(name); flag != nil {
			cmdFlags.AddFlag(flag)
		} else {
			//logger.Fatalf("Could not find flag '%s' to attach to command '%s'", name, cmd.Name())
		}
	}
}
