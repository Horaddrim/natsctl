// Copyright © 2018 Igor Franca <lee12rock@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/horaddrim/natsctl/nats"
	"github.com/horaddrim/natsctl/utils"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Retorna o status do servidor.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := utils.NewHTTPClient()
		server := new(nats.Server)

		resp, err := client.Get("http://40.121.32.164:8222/connz")

		if err != nil {
			fmt.Printf("Ocorreu um erro %s", err.Error())
		}

		buff, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(buff, server)

		fmt.Printf("TIME  ALIVE_CONNECTIONS  TOTAL_CONNECTIONS  SERVER\n")
		fmt.Printf("%s | %d | %d  | %s", server.ServerTime, server.ActiveConnections, server.TotalConnections, server.ID)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
