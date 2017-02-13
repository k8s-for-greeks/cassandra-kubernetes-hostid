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

	"github.com/k8s-for-greeks/cassandra-kubernetes-hostid/pkg/hostid"
	"github.com/spf13/cobra"
)

func SetupCassandraClient(cmd *cobra.Command, nodeTool string)(c *hostid.CasssandraHostId, err error) {

	p, err := cmd.Parent().PersistentFlags().GetString("pod")
	if err != nil {
		return nil, fmt.Errorf("error getting pod flag: %s\n", err)
	} else if p == "" {
		return nil, fmt.Errorf("pod flag not set")
	}

	ns, err := cmd.Parent().PersistentFlags().GetString("namespace")
	if err != nil {
		return nil, fmt.Errorf("error getting namespace flag: %s\n", err)
	}

	ann, err := cmd.Parent().PersistentFlags().GetString("annotation")
	if err != nil {
		return nil, fmt.Errorf("error getting annotation flag: %s\n", err)
	} else if ann == "" {
		ann = "cassandra"
	}

	return hostid.CreateCasssandraHostId(nodeTool, p, ns, ann)
}

