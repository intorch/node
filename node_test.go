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
	"testing"
	"time"

	"github.com/intorch/message"
	"github.com/stretchr/testify/assert"
)

var nodeTestEngine = Engine(func(msg message.Message) message.Message {
	return msg
})

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	n1 := Create(&nodeTestEngine, nil, "")

	assert.NotNil(n1)
	assert.NotNil(n1.Channel)
	assert.NotNil(n1.Engine)
	assert.NotEmpty(n1.ID)
	assert.True(&nodeTestEngine == n1.Engine)
}

func TestCreate_AllParameters(t *testing.T) {
	assert := assert.New(t)

	ch := NewChannel(10)
	ID := "1234567890"
	n2 := Create(&nodeTestEngine, ch, ID)

	assert.NotNil(n2)
	assert.NotNil(n2.Channel)
	assert.NotNil(n2.Engine)
	assert.NotEmpty(n2.ID)
	assert.Equal(n2.ID, "1234567890")
	assert.True(&nodeTestEngine == n2.Engine)
	assert.True(ch == n2.Channel)
}

func TestNew(t *testing.T) {
	assert := assert.New(t)

	n1 := New(&nodeTestEngine)

	assert.NotNil(n1)
	assert.NotNil(n1.Channel)
	assert.NotNil(n1.Engine)
	assert.NotEmpty(n1.ID)
	assert.True(&nodeTestEngine == n1.Engine)
}

func TestNode_Write(t *testing.T) {
	assert := assert.New(t)
	n1 := New(&nodeTestEngine)

	msg := message.New(make(message.Header), make(message.Body))
	n1.Write(*msg)

	assert.Len(n1.Channel.Input, 1)
}

func TestNode_Read(t *testing.T) {
	assert := assert.New(t)
	n1 := New(&nodeTestEngine)

	//write one message in the node
	msg := message.New(make(message.Header), make(message.Body))
	n1.Write(*msg)

	//asset that the input channel have one message
	assert.Len(n1.Channel.Input, 1)

	//start cosumer
	go func() {
		n1.Run()
	}()

	//wait for cosumer and check if output channel have one message
	time.Sleep(time.Second)
	assert.Len(n1.Channel.Output, 1)

	//read the message
	resp := n1.Read()

	//assert the message is not nil, that it's the some of input and
	//the outputchanel have no message
	assert.NotNil(resp)
	assert.True(resp.Equals(msg))
	assert.Len(n1.Channel.Output, 0)
}

func TestNode_GetReader(t *testing.T) {
	assert := assert.New(t)
	n1 := New(&nodeTestEngine)

	reader := n1.GetReader()

	assert.NotNil(reader)
	assert.Equal(reader, n1.Channel.Output)
}

func TestNode_Run(t *testing.T) {
	assert := assert.New(t)

	//create node using the fatorial engine that was created at the engine
	//test case
	node := New(&fatorial)

	//create message and send it to channel
	msg := message.New(make(message.Header), make(message.Body))
	(*msg.Body)["value"] = 10
	node.Write(*msg)

	//start node
	go func() {
		node.Run()
	}()

	//receive message
	resp := node.Read()
	value := (*resp.Body)["fat"].(int)

	assert.Equal(value, 3628800)

	//stop channel receive
	node.Stop()
}

func TestNode_Stop(t *testing.T) {
	assert := assert.New(t)
	TENMILLION := 10000000

	var doNothingEngine = Engine(func(msg message.Message) message.Message { return msg })
	node := Create(&doNothingEngine, NewChannel(uint(TENMILLION)), "test-case")

	//create message and send it to channel
	msg := message.New(make(message.Header), make(message.Body))
	(*msg.Body)["value"] = 10

	//add One Million messages to the node
	for i := 0; i < TENMILLION; i++ {
		node.Write(*msg)
	}

	assert.Len(node.Channel.Input, TENMILLION)

	//start node
	go func() {
		node.Run()
	}()

	//stop node  miliseconds after node starts
	node.Stop()

	//All One million messages should be consumed
	assert.Len(node.Channel.Output, TENMILLION)
}
