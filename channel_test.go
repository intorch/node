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

func TestChannel_CreateInput(t *testing.T) {
	assert := assert.New(t)
	ch := &Channel{}

	//check the existence of channel
	assert.NotNil(ch)

	//grant the both, input and output born as nil
	assert.Nil(ch.Input)
	assert.Nil(ch.Output)

	//grant that once create a new Input with 0 as size
	//it'll have the capacity equal 1.
	//The output should continue as nil
	ch.CreateInput(0)
	assert.NotNil(ch.Input)
	assert.Nil(ch.Output)
	assert.Equal(1, cap(ch.Input))

	//grant that once create a new Input with 1 as size it'll have the capacity
	// equal 1.
	//The output should continue as nil
	//It also intrinsically test the capacity of chanel ovewriting
	ch.CreateInput(1)
	assert.NotNil(ch.Input)
	assert.Nil(ch.Output)
	assert.Equal(1, cap(ch.Input))

	//grant that once create a new Input with 10 as size it'll have the capacity
	//equal 10.
	//The output should continue as nil
	//It also intrinsically test the capacity of chanel ovewriting
	ch.CreateInput(10)
	assert.NotNil(ch.Input)
	assert.Nil(ch.Output)
	assert.Equal(10, cap(ch.Input))
}

func TestChannel_CreateInputManualy(t *testing.T) {
	assert := assert.New(t)
	ch := &Channel{}

	//check the existence of channel
	assert.NotNil(ch)

	//grant the both, input and output born as nil
	assert.Nil(ch.Input)
	assert.Nil(ch.Output)

	//grant that Input with have the capacity equal 1.
	//The output should continue as nil
	ch.Input = make(chan message.Message)
	assert.NotNil(ch.Input)
	assert.Nil(ch.Output)
	assert.Equal(0, cap(ch.Input))

	//grant that Input with have the capacity equal 10.
	//The output should continue as nil
	//It also intrinsically test the capacity of chanel ovewriting
	ch.Input = make(chan message.Message, 10)
	assert.NotNil(ch.Input)
	assert.Nil(ch.Output)
	assert.Equal(10, cap(ch.Input))
}

func TestChannel_CreateOutputManualy(t *testing.T) {
	assert := assert.New(t)
	ch := &Channel{}

	//check the existence of channel
	assert.NotNil(ch)

	//grant the both, input and output born as nil
	assert.Nil(ch.Input)
	assert.Nil(ch.Output)

	//grant that Output with have the capacity equal 1.
	//The output should continue as nil
	ch.Output = make(chan message.Message)
	assert.NotNil(ch.Output)
	assert.Nil(ch.Input)
	assert.Equal(0, cap(ch.Output))

	//grant that Output with have the capacity equal 10.
	//The output should continue as nil
	//It also intrinsically test the capacity of chanel ovewriting
	ch.Output = make(chan message.Message, 10)
	assert.NotNil(ch.Output)
	assert.Nil(ch.Input)
	assert.Equal(10, cap(ch.Output))
}

func TestChannel_CreateOutput(t *testing.T) {
	assert := assert.New(t)

	ch := &Channel{}

	//check the existence of channel
	assert.NotNil(ch)

	//grant the both, input and output born as nil
	assert.Nil(ch.Input)
	assert.Nil(ch.Output)

	//grant that once create a new Output with 0 as size
	//it'll have the capacity equal 1.
	//The output should continue as nil
	ch.CreateOutput(0)
	assert.NotNil(ch.Output)
	assert.Nil(ch.Input)
	assert.Equal(1, cap(ch.Output))

	//grant that once create a new Input with 1 as size it'll have the capacity
	// equal 1.
	//The output should continue as nil
	//It also intrinsically test the capacity of chanel ovewriting
	ch.CreateOutput(1)
	assert.NotNil(ch.Output)
	assert.Nil(ch.Input)
	assert.Equal(1, cap(ch.Output))

	//grant that once create a new Input with 10 as size it'll have the capacity
	//equal 10.
	//The output should continue as nil
	//It also intrinsically test the capacity of chanel ovewriting
	ch.CreateOutput(10)
	assert.NotNil(ch.Output)
	assert.Nil(ch.Input)
	assert.Equal(10, cap(ch.Output))
}
