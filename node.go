// Copyright 2021 intorch.org. All Rights Reserved.
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

package core

import "github.com/google/uuid"

//Node data structure which is formed by the components necessary to
//create a micro execution layer.
type Node struct {
	ID      string
	Channel *Channel
	Engine  *Engine
}

func New(engine *Engine, channel *Channel, ID string) *Node {
	if len(ID) == 0 {
		ID = uuid.NewString()
	}

	validate(engine, channel, ID)

	return &Node{
		Engine:  engine,
		Channel: channel,
		ID:      ID,
	}
}

//Run micro kernel. It'll start a go runtine to receive messages and
//process it until a shutdown signal is received through a closure of
//the inlet channel.
//Before run, it'll validade if all parameters are present. If anyone is
//not present the system stop execution with status 1.
func (n Node) Run() {
	validate(n.Engine, n.Channel, n.ID)

	for msg := range n.Channel.Input {
		resp := (*n.Engine)(msg)
		n.Channel.Output <- resp
	}
}
