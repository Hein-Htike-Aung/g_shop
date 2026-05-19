package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"yixiang.co/go-mall/pkg/util"
)

// Env holds global --env flag
var Env string

// RegisterGlobalFlags registers global flags
func RegisterGlobalFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "load .env file, example: --env=testing will use .env.testing file")
}

// RegisterDefaultCmd sets default subcommand
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firtArg := util.FirstElement(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && firtArg != "-h" && firtArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}
