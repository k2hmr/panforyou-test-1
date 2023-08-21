package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/k2hmr/panforyou-test-1/internal/api"
	"github.com/k2hmr/panforyou-test-1/internal/validator"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "panforyou-test-1",
	Short: "データを取得してDBに保存するCLIツールです。",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("引数を１つだけ指定してください。")
		}
		if !validator.IsEntryIdValid(args[0]) {
			return errors.New("引数は無効な形式です。")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		bread, err := api.FetchBread(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ctx := context.Background()
		err = api.SaveBread(*bread, ctx)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
