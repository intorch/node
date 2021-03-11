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
	"github.com/intorch/message"
)

//Channel data structure to support the message flow in the application micro
//Kernel. This data structure use channels (chan) to control the inbound and
//outbound messges.
//Note that this channel has no direction. It's important, becouse the output
//message in this kernel can be an input in other one.
type Channel struct {

	//Input message chan. In this Channel, this one will behave itself as the
	//inbound channel
	Input chan message.Message

	//Output message chan. In this Channel, this one will behave itself as the
	//outbound channel
	Output chan message.Message
}

//CreateInput function to create a new channel message. Case the input is already
//set, it'll be ovewrited by the new one.
//The parameter size is used to set the channel size. The minimum acceptable is 1,
//so, case 0 is set it'll replaced to 1.
func (ch *Channel) CreateInput(size uint) {
	if size == 0 {
		size = 1
	}

	ch.Input = make(chan message.Message, size)
}

//CreateOutput function to create a new channel message. Case the output is already
//set, it'll be ovewrited by the new one.
//The parameter size is used to set the channel size. The minimum acceptable is 1,
//so, case 0 is set it'll replaced to 1.
func (ch *Channel) CreateOutput(size uint) {
	if size == 0 {
		size = 1
	}

	ch.Output = make(chan message.Message, size)
}
