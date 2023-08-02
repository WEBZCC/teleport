// Copyright 2017 Gravitational, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/gravitational/teleport/tool/tctl/common"
)

func main() {

	// aggregate common and oss-specific command variants
	commands := common.Commands()
	commands = append(commands, common.OSSCommands()...)

	// TODO(zmb3): add devices command to common commands
	// after the command has been removed from enterprise
	commands = append(commands, &common.DevicesCommand{})

	common.Run(commands)
}
