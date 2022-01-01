/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// peersCmd represents the peers command
var peersCmd = &cobra.Command{
	Use:   "peers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:26657/net_info")
		if err != nil {
			// handle err
			fmt.Printf("%s", err)
		}
		//We Read the response body on the line below.
		if err != nil {
			log.Fatalln(err)
		}

		var result map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&result)
		height := result["result"].(map[string]interface{})["n_peers"]
		fmt.Printf("%s", height)
	},
}

func init() {
	utilsCmd.AddCommand(peersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// peersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// peersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
