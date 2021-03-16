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

//the Engine that do fatorial
var fatorial = Engine(func(msg message.Message) message.Message {
	body := (*msg.Body)

	vl := body["value"]

	acc := 1
	for i := vl.(int); i > 0; i-- {
		acc = acc * i
	}

	body["fat"] = acc

	return msg
})

func TestNew(t *testing.T) {
	assert := assert.New(t)

	node := New(&fatorial, nil, "")

	assert.NotNil(node)
	assert.NotNil(node.Channel)
	assert.NotNil(node.Engine)
	assert.NotEmpty(node.ID)
}

func TestNewEquals(t *testing.T) {
	assert := assert.New(t)

	ch := &Channel{}

	node := New(&fatorial, ch, "")

	assert.NotNil(node)
	assert.NotNil(node.Channel)
	assert.NotNil(node.Engine)
	assert.NotEmpty(node.ID)

	assert.Equal(node.Channel, ch)
}
