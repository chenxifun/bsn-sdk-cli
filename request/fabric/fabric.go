package fabric

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	api      string
	userCode string
	appCode  string

	key string

	isTest  bool
	isTrust bool

	mspDir string

	subUser      string
	registerUser bool

	chainCode string
	funcName  string
	args      []string
)
var flags *pflag.FlagSet

func init() {
	flags = &pflag.FlagSet{}
	flags.StringVar(&api, "api", "", "请求的网关地址")
	flags.StringVar(&userCode, "userCode", "", "用户编号")
	flags.StringVar(&appCode, "appCode", "", "应用编号")
	flags.StringVar(&key, "key", "key.pem", "私钥文件的路径")

	flags.BoolVar(&isTest, "test", false, "是否调用测试环境")
	flags.BoolVar(&isTrust, "trust", false, "是否托管的应用")

	flags.StringVar(&mspDir, "msp", "msp/", "子用户证书文件的存储路径")

	flags.StringVar(&subUser, "subUser", "", "调用交易的子用户")
	flags.BoolVar(&registerUser, "regSubUser", false, "是否需要注册子用户")

	flags.StringVar(&chainCode, "chainCode", "", "调用的链码")
	flags.StringVar(&funcName, "fcn", "set", "调用的方法")
	flags.StringArrayVar(&args, "args", nil, "调用的参数")

}

func Cmd() *cobra.Command {
	var cobraCommand = &cobra.Command{
		Use:   "fabric",
		Short: "调用fabric交易",
		Long:  `调用网关发起一个fabric交易`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.SilenceUsage = true
			goRequest()

		},
	}
	flagList := []string{
		"api",
		"userCode",
		"appCode",
		"key",
		"test",
		"trust",
		"msp",
		"subUser",
		"regSubUser",
		"chainCode",
		"fcn",
		"args",
	}
	attachFlags(cobraCommand, flagList)
	return cobraCommand
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
