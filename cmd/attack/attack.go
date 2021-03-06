// Copyright 2020 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package attack

import "github.com/spf13/cobra"

func NewAttackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attack <subcommand>",
		Short: "Attack related commands",
	}

	cmd.AddCommand(
		NewProcessAttackCommand(),
		NewNetworkAttackCommand(),
		NewStressAttackCommand(),
		NewDiskAttackCommand(),
		NewHostAttackCommand(),
	)

	return cmd
}
