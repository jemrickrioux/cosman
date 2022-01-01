/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		daemon := viper.Get("settings").(map[string]interface{})["daemon"].(string)
		moniker := viper.Get("settings").(map[string]interface{})["moniker"].(string)
		chain_id := viper.Get("settings").(map[string]interface{})["chain_id"].(string)
		commission_rate := viper.Get("settings").(map[string]interface{})["commission_rate"].(string)
		fmt.Println("create called")
		shell(daemon, []string{"tx", "staking", "create-validator", "--amount=100000uakt", "--pubkey=$(" + daemon + " tendermint show-validator)", "--moniker=" + moniker, "--chain-id=" + chain_id, "--commission-rate=" + commission_rate})
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
