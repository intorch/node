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

//Kernel data structure which is formed by the components necessary to
//create a micro execution layer.
type Kernel struct {
	ID      string
	Channel *Channel
	Engine  *Engine
}

//Run micro kernel. It'll start a go runtine to receive messages and
//process it until a shutdown signal is received through a closure of
//the inlet channel.
//Before run, it'll validade if all parameters are present. If anyone is
//not present the system stop execution with status 1.
func (k Kernel) Run() {
	validate(k.Engine, k.Channel, k.ID)

	for msg := range k.Channel.Input {
		resp := (*k.Engine)(msg)
		k.Channel.Output <- resp
	}
}
