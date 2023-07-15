package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"go-cmb-service/cmd/api"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "go-cmb-service",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `go-cmb-service`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("需要参数")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `go-cmb-service 1.0.0 欢迎使用，可以是用 -h 查看命令`
		log.Printf("%s\n", usageStr)
	},
}

func init() {
	// 添加每日同步

	rootCmd.AddCommand(api.StartCmd)
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {

		os.Exit(-1)
	}
}
