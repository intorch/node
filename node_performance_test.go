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

func TestNode_Run_OneMillionFatorial(t *testing.T) {
	assert := assert.New(t)
	ONEMILLION := 1000000
	FATORIALVALUE := 10
	FATORIALRESULT := 3628800

	//create channel
	channel := NewChannel(uint(ONEMILLION))

	//create node
	node := Create(&fatorial, channel, "")

	start := time.Now()
	//create message and post it in a channel. 10M timmes
	go func() {
		for i := 0; i < ONEMILLION; i++ {
			msg := message.New(make(message.Header), make(message.Body))
			(*msg.Body)["value"] = FATORIALVALUE

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

		assert.Equal(value, FATORIALRESULT)

		//cancel channel after receive 10M received
		if count == ONEMILLION {
			close(channel.Output)
		}
	}
	elapsed := time.Since(start).Seconds()

	//100k fatorial operations p/seconds
	assert.Greater(float64(10), elapsed)
}

func TestNode_Run_OneMillionDoNothing(t *testing.T) {
	assert := assert.New(t)
	ONEMILLION := 1000000

	//create channel
	channel := NewChannel(uint(ONEMILLION))

	//the Engine that do noting
	var doNothingEngine = Engine(func(msg message.Message) message.Message { return msg })
	node := Create(&doNothingEngine, channel, "")

	start := time.Now()
	//create message and post it in a channel. 10M timmes
	go func() {
		for i := 0; i < ONEMILLION; i++ {
			msg := message.New(make(message.Header), make(message.Body))
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
		assert.NotNil(msg)
		count = count + 1

		//cancel channel after receive 10M received
		if count == ONEMILLION {
			close(channel.Output)
		}
	}
	elapsed := time.Since(start).Seconds()

	//200k operations p/seconds
	assert.Greater(float64(5), elapsed)
}
