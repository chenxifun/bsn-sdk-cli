package random

import (
	"encoding/base64"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/crypto"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
)

var (
	num  int
	path string
)
var flags *pflag.FlagSet

func init() {
	flags = &pflag.FlagSet{}
	flags.IntVarP(&num, "num", "n", 1, "生成的随机数个数")
	flags.StringVarP(&path, "path", "p", "nonce.txt", "生成的文件名以及文件存储位置")
}

func Cmd() *cobra.Command {
	var cobraCommand = &cobra.Command{
		Use:   "random",
		Short: "生成随机字符串",
		Long:  `生成若干个长度为24位的byte数组，并且转换为base64格式的字符串，保存到文件中`,
		Run: func(cmd *cobra.Command, args []string) {

			cmd.SilenceUsage = true

			Random(num, path)
		},
	}
	flagList := []string{
		"num",
		"path",
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

func Random(num int, filePath string) {

	if num <= 0 {
		num = 1
	}

	var n string

	for i := 0; i < num; i++ {
		nonce, _ := crypto.GetRandomNonce()
		n = n + base64.StdEncoding.EncodeToString(nonce) + "\r\n"
		//fmt.Println()
	}
	filename := filePath //"./nonce.txt"
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		// 创建文件失败处理
		fmt.Println("创建文件失败：", err)

	} else {

		_, err = f.Write([]byte(n))
		if err != nil {
			// 写入失败处理
			fmt.Println("写入文件失败：", err)
		}
	}

}
