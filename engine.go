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
	"log"

	"github.com/intorch/message"
)

//Engine type to define the way to operate and execute every component inside
//the kernel.
type Engine func(msg message.Message) message.Message

//ThisIsJustAnExample Do not use it in production! It's here just for a test purpose.
var ThisIsJustAnExample = Engine(func(msg message.Message) message.Message {
	log.Println("Do not use it in production!!!")

	return msg
})
