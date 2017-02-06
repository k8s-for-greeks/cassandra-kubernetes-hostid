// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/k8s-for-greeks/cassandra-kubernetes-hostid/pkg/hostid"
	"os"
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

		fmt.Println("create called")
		n, err := cmd.Flags().GetString("nodetool")
		if err != nil {
			fmt.Println("nodetool flag not set")
			os.Exit(2)
		}
		p, err := cmd.Flags().GetString("pod")
		if err != nil {
			fmt.Println("pod flag not set")
			os.Exit(2)
		}
		ns, err := cmd.Flags().GetString("namespace")
		if err != nil {
			fmt.Println("pod flag not set")
			os.Exit(2)
		}
		ann, err := cmd.Flags().GetString("annotation")
		if err != nil {
			os.Exit(2)
		}
		c := hostid.CreateCasssandraHostId(n,p,ns,ann)

		err = c.SaveHostId()

		if err != nil {
			os.Exit(2)
		}
		fmt.Println("Create hostId")
	},
}

func init() {
	hostIdCmd.cobraCommand.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
