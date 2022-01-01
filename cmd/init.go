/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func shell(program string, args []string) error {
	cmd := exec.Command(program, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		daemon := viper.Get("settings").(map[string]interface{})["daemon"]
		moniker := viper.Get("settings").(map[string]interface{})["moniker"]
		chain_id := viper.Get("settings").(map[string]interface{})["chain_id"]
		home := viper.Get("settings").(map[string]interface{})["home"]
		genesis := viper.Get("settings").(map[string]interface{})["genesis"]
		shell(daemon.(string), []string{"init", moniker.(string), "--chain-id", chain_id.(string)})
		fmt.Println("HOME DIRECTORY", home.(string)+"/config/genesis.json")
		os.Remove(home.(string) + "/config/genesis.json")
		shell("wget", []string{genesis.(string)})
		err := os.Rename("genesis.json", home.(string)+"/config/genesis.json")
		if err != nil {
			fmt.Println(err)
		}
		os.Remove("genesis.json")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
