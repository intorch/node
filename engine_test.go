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

func TestEngine_ThisIsJustAnExample(t *testing.T) {
	assert := assert.New(t)

	msg := message.New(
		make(message.Header),
		make(message.Body),
	)

	resp := ThisIsJustAnExample(*msg)

	assert.NotNil(resp)
	assert.Equal(*msg, resp)
}

func TestEngine_ResponseNotEquals(t *testing.T) {
	assert := assert.New(t)

	eng := Engine(func(msg message.Message) message.Message {
		(*msg.Body)["hello"] = "world"

		return msg
	})

	msg := message.New(
		make(message.Header),
		make(message.Body),
	)

	resp := eng(*msg)

	assert.NotNil(resp)
	assert.True((*msg).Equals(&resp))
	assert.Equal((*resp.Body)["hello"], "world")
}
