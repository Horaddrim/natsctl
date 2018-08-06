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
	"strconv"

	"github.com/fatih/color"
	"github.com/horaddrim/natsctl/nats"
	"github.com/horaddrim/natsctl/utils"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Returns the server status",
	Long:  `This command will return a table pretty printing you server's info to you.`,
	Run: func(cmd *cobra.Command, args []string) {
		var protocol string
		client := utils.NewHTTPClient()
		server := new(nats.Server)

		hostIP := cmd.Flag("host").Value.String()
		serverPort := cmd.Flag("port").Value.String()
		forceHTTPS, err := strconv.ParseBool(cmd.Flag("force-https").Value.String())

		if err != nil {
			fmt.Printf("An error have ocured %s", err.Error())
		}

		if forceHTTPS {
			protocol = "https"
		} else {
			protocol = "http"
		}

		finalURL := nats.PrepareHost(protocol, hostIP, serverPort, nats.VariableRoute)

		color.Magenta("Connecting to server %s", hostIP)

		resp, err := client.Get(finalURL)

		if err != nil {
			fmt.Printf("An error have ocured %s", err.Error())
		}

		buff, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(buff, server)

		memoryUsage := strconv.FormatInt((server.MemoryUsage / 1024 / 1024), 10)

		outputData := [][]string{
			[]string{server.ID,
				server.ServerTime.String(),
				server.Age,
				fmt.Sprintf("%s MB", memoryUsage),
			},
		}

		table := utils.NewTable(nats.DefaultHeaders)

		table.AppendBulk(outputData)

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
	statusCmd.PersistentFlags().String("port", "8222", "This is the port to connect on the given host")
	statusCmd.PersistentFlags().Bool("force-https", true, "Flag to force the use of the HTTPS protocol")
}
