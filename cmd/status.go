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
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/horaddrim/natsctl/nats"
	"github.com/horaddrim/natsctl/utils"
	"github.com/olekukonko/tablewriter"
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
		table := tablewriter.NewWriter(os.Stdout)

		hostIP := cmd.Flag("host").Value.String()

		color.Magenta("Connecting to server %s", hostIP)

		host := fmt.Sprintf("http://%s:8222", hostIP)

		finalURL := fmt.Sprintf("%s/connz", host)

		resp, err := client.Get(finalURL)

		if err != nil {
			fmt.Printf("Ocorreu um erro %s", err.Error())
		}

		buff, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(buff, server)

		table.SetHeader([]string{"SERVER", "TIMESTAMP", "ACTIVE_CONNECTIONS", "TOTAL_CONNECTIONS"})

		outputData := [][]string{
			[]string{server.ID, server.ServerTime.String(), strconv.FormatInt(server.ActiveConnections, 10), strconv.FormatInt(server.TotalConnections, 10)},
		}

		for _, line := range outputData {
			table.Append(line)
		}

		table.Render()
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
	statusCmd.PersistentFlags().String("host", "localhost", "This is the host aimed to connect")
	rootCmd.MarkPersistentFlagRequired("host")
}
