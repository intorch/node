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

func TestEngine_OneMilionDoNothing(t *testing.T) {
	assert := assert.New(t)

	//the Engine that do noting
	var doNothingEngine = Engine(func(msg message.Message) message.Message { return msg })

	msg := message.New(make(message.Header), make(message.Body))

	ONEMILLION := 1000000

	start := time.Now()
	for i := 0; i < ONEMILLION; i++ {
		resp := doNothingEngine(*msg)

		assert.NotNil(resp)
		assert.Equal(*msg, resp)
	}
	elapsed := time.Since(start).Seconds()

	assert.Greater(float64(15), elapsed)
}

func TestEngine_OneMilionFatorial(t *testing.T) {
	assert := assert.New(t)

	msg := message.New(make(message.Header), make(message.Body))

	(*msg.Body)["value"] = 10

	ONEMILLION := 1000000

	start := time.Now()
	for i := 0; i < ONEMILLION; i++ {
		resp := fatorial(*msg)

		assert.NotNil(resp)
		assert.True((*msg).Equals(&resp))
		assert.Equal((*resp.Body)["fat"], 3628800)
	}
	elapsed := time.Since(start).Seconds()

	assert.Greater(float64(20), elapsed)
}
