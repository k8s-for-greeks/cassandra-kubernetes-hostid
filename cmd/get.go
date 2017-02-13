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

	"github.com/k8s-for-greeks/cassandra-kubernetes-hostid/pkg/hostid"
	"github.com/spf13/cobra"
)

type GetHostIdOptions struct {
}

func NewGetHostIdCmd(out io.Writer) *cobra.Command {
	// getCmd represents the get command
	c := &cobra.Command{
		Use:   "get",
		Short: "get hostid",
		Long:  `TODO`,
		Run: func(cmd *cobra.Command, args []string) {
			RunGetHostIdCmd(cmd, out)
		},
	}

	return c
}

func RunGetHostIdCmd(cmd *cobra.Command, out io.Writer) error {
	//n, err := cmd.Flags().GetString("nodetool")
	n, err := cmd.Parent().PersistentFlags().GetString("nodetool")
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

	c, err := hostid.CreateCasssandraHostId(n, p, ns, ann)

	if err != nil {
		fmt.Println("error setting up host id: %s", err)
		os.Exit(2)
	}

	h, err := c.GetHostId()

	if err != nil {
		fmt.Println("error getting host id: %s", err)
		os.Exit(2)
	}

	// TODO fix
	fmt.Printf("%s", h)

	return nil
}
