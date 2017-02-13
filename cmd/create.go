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

	c, err := hostid.CreateCasssandraHostId(n, p, ns, ann)

	if err != nil {
		os.Exit(2)
	}

	err = c.SaveHostId()

	if err != nil {
		os.Exit(2)
	}
	fmt.Println("Create hostId")

	return nil
}
