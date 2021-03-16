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

import (
	"testing"

	"github.com/intorch/message"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	node := New(&fatorial, nil, "")

	assert.NotNil(node)
	assert.NotNil(node.Channel)
	assert.NotNil(node.Engine)
	assert.NotEmpty(node.ID)
}

func TestNewChannelEquals(t *testing.T) {
	assert := assert.New(t)

	ch := &Channel{}

	node := New(&fatorial, ch, "1234567890")

	assert.NotNil(node)
	assert.NotNil(node.Channel)
	assert.NotNil(node.Engine)
	assert.NotEmpty(node.ID)

	assert.Equal(node.Channel, ch)
	assert.Equal("1234567890", node.ID)
}

func TestNode_Run(t *testing.T) {
	assert := assert.New(t)

	//create channel
	channel := &Channel{}
	channel.CreateInput(1)
	channel.CreateOutput(1)

	//create node
	node := New(&fatorial, channel, "")

	//create message and send it to channel
	msg := message.New(make(message.Header), make(message.Body))
	(*msg.Body)["value"] = 10
	channel.Input <- *msg

	//start node
	go func() {
		node.Run()
	}()

	//receive message
	resp := <-channel.Output
	value := (*resp.Body)["fat"].(int)

	assert.Equal(value, 3628800)

	//stop channel receive
	close(channel.Output)
}

func TestNode_Run10MilionFatorial(t *testing.T) {
	assert := assert.New(t)

	//create channel
	channel := &Channel{}
	channel.CreateInput(1000000)
	channel.CreateOutput(1000000)

	//create node
	node := New(&fatorial, channel, "")

	//create message and post it in a channel. 10M timmes
	go func() {
		ONEMILLION := 10000000
		for i := 0; i < ONEMILLION; i++ {
			msg := message.New(
				make(message.Header),
				make(message.Body),
			)

			(*msg.Body)["value"] = 10

			channel.Input <- *msg
		}
	}()

	//start node
	go func() {
		node.Run()
	}()

	//start message receive
	count := 0
	for msg := range channel.Output {
		value := (*msg.Body)["fat"].(int)
		count = count + 1

		assert.Equal(value, 3628800)

		//cancel channel after receive 10M received
		if count == 10000000 {
			close(channel.Output)
		}
	}
}

func TestNode_Run10MilionDoNothing(t *testing.T) {
	assert := assert.New(t)

	//create channel
	channel := &Channel{}
	channel.CreateInput(1000000)
	channel.CreateOutput(1000000)

	//create node
	node := New(&fatorial, channel, "")

	//create message and post it in a channel. 10M timmes
	go func() {
		ONEMILLION := 10000000
		for i := 0; i < ONEMILLION; i++ {
			msg := message.New(
				make(message.Header),
				make(message.Body),
			)

			(*msg.Body)["value"] = 10

			channel.Input <- *msg
		}
	}()

	//start node
	go func() {
		node.Run()
	}()

	//start message receive
	count := 0
	for msg := range channel.Output {
		// value := (*msg.Body)["fat"].(int)
		value := msg
		count = count + 1

		assert.NotNil(value)

		//cancel channel after receive 10M received
		if count == 10000000 {
			close(channel.Output)
		}
	}
}
