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

package node

import (
	"time"

	"github.com/google/uuid"
	"github.com/intorch/message"
)

//Node data structure which is formed by the components necessary to
//create a micro execution layer.
type Node struct {
	ID      string
	Channel *Channel
	Engine  *Engine
}

func New(engine *Engine) *Node {
	return Create(engine, nil, "")
}

func Create(engine *Engine, channel *Channel, ID string) *Node {
	assertNotNil(engine, "The node cannot be created withot an engine")

	if len(ID) == 0 {
		ID = uuid.NewString()
	}

	if channel == nil {
		CAPACITY := 10
		channel = NewChannel(uint(CAPACITY))
	}

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

//Write add new message to be processed by the engine
func (n Node) Write(msg message.Message) {
	n.Channel.Input <- msg
}

//Read message that already been processed
func (n Node) Read() message.Message {
	return <-n.Channel.Output
}

//GetReader get the output channel
func (n Node) GetReader() chan message.Message {
	return n.Channel.Output
}

//Stop node shutdown
func (n Node) Stop() {
	close(n.Channel.Input)

	for len(n.Channel.Input) > 0 {
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(50 * time.Millisecond)
}
