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
	"io"

	"github.com/spf13/cobra"
)

type HostIdOptions struct {
	Pod        string
	Namespace  string
	Annotation string
	Nodetool   string
}

func NewHostIdCmd(out io.Writer) *cobra.Command {

	options := &HostIdOptions{}

	c := &cobra.Command{
		Use:   "hostid",
		Short: "hostid",
		Long:  `todo`,
	}

	c.PersistentFlags().StringVarP(&options.Pod, "pod", "p", "", "Pod Name")
	c.PersistentFlags().StringVarP(&options.Namespace, "namespace", "n", "", "Pod Namespace")
	c.PersistentFlags().StringVarP(&options.Annotation, "annotation", "a", "", "Annotation Prefix")
	c.PersistentFlags().StringVarP(&options.Nodetool, "nodetool", "t", "", "Nodetool Path")

	c.AddCommand(NewCreateHostIdCmd(out))
	c.AddCommand(NewGetHostIdCmd(out))
	return c
}
