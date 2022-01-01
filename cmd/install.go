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

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long:  `This command makes the necessary installation for the chain`,
	Run: func(cmd *cobra.Command, args []string) {
		github := viper.Get("settings").(map[string]interface{})["github"]
		folder_name := viper.Get("settings").(map[string]interface{})["folder_name"]
		home := viper.Get("settings").(map[string]interface{})["home"]
		genesis := viper.Get("settings").(map[string]interface{})["genesis"]
		moniker := viper.Get("settings").(map[string]interface{})["moniker"]
		fmt.Println(github, folder_name, home, genesis, moniker)
		cmdr := exec.Command("/bin/sh", "-c", "git clone "+github.(string))
		err := cmdr.Run()
		if err != nil {
			fmt.Println(err)
		}
		os.Chdir("./" + folder_name.(string))
		newDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		cmdr1 := exec.Command("/bin/sh", "-c", "make install")
		err = cmdr1.Run()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Current Working Direcotry: %s\n", newDir)

	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
