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

	"io"
	"os"

	"github.com/spf13/cobra"
)

type CreateHostIdOptions struct {
}

func NewCreateHostIdCmd(out io.Writer) *cobra.Command {

	// createCmd represents the create command
	c := &cobra.Command{
		Use:   "create",
		Short: "create a hostid for a cassandra instance",
		Long:  `TODO`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO check errors better
			RunCreateHostIdCmd(cmd, out)
		},
	}

	return c
}

func RunCreateHostIdCmd(cmd *cobra.Command, out io.Writer) error {

	// TODO do we always need this path??
	n, err := cmd.Parent().PersistentFlags().GetString("nodetool")
	if err != nil {
		fmt.Println("nodetool flag not set")
		os.Exit(2)
	}

	c, err := SetupCassandraClient(cmd,n)

	if err != nil {
		fmt.Printf("error setting up host id: %s\n", err)
		os.Exit(2)
	}


	err = c.SaveHostId()

	if err != nil {
		fmt.Printf("error saving host id: %s\n", err)
		os.Exit(2)
	}

	return nil
}
