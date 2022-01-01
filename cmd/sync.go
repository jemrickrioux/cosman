/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		rpc := viper.Get("sync").(map[string]interface{})["rpc"].(string)

		// curl -s $SNAP_RPC/block | jq -r .result.block.header.height)

		resp, err := http.Get(rpc + "/block")
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
		height := result["result"].(map[string]interface{})["block"].(map[string]interface{})["header"].(map[string]interface{})["height"]
		var new int
		if s, err := strconv.Atoi(height.(string)); err == nil {
			new = s - 1000
		}

		resp1, err := http.Get(rpc + "/block?height=" + strconv.Itoa(new))
		if err != nil {
			log.Fatalln(err)
		}

		var result2 map[string]interface{}

		json.NewDecoder(resp1.Body).Decode(&result2)
		hash := result2["result"].(map[string]interface{})["block_id"].(map[string]interface{})["hash"]
		fmt.Println(`
[statesync]
# State sync rapidly bootstraps a new node by discovering, fetching, and restoring a state machine
# snapshot from peers instead of fetching and replaying historical blocks. Requires some peers in
# the network to take and serve state machine snapshots. State sync is not attempted if the node
# has any local state (LastBlockHeight > 0). The node will have a truncated block history,
# starting from the height of the snapshot.
enable = true

# RPC servers (comma-separated) for light client verification of the synced state machine and
# retrieval of state data for node bootstrapping. Also needs a trusted height and corresponding
# header hash obtained from a trusted source, and a period during which validators can be trusted.
#
# For Cosmos SDK-based chains, trust_period should usually be about 2/3 of the unbonding time (~2
# weeks) during which they can be financially punished (slashed) for misbehavior.`)
		fmt.Println(`trust_height = "` + height.(string) + `"`)
		fmt.Println(`trust_hash = "` + hash.(string) + `"`)
		fmt.Println(`rpc_servers = "` + rpc + "," + rpc + `"`)
		fmt.Println(`trust_period = "168h0m0s"`)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
