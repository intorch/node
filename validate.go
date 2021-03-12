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
)

func assertNotEmpty(data string, message string) {
	if len(data) == 0 {
		log.Fatal(message)
	}
}

func assertNotNil(obj interface{}, message string) {
	if obj != nil {
		log.Fatal(message)
	}
}

func validate(engine *Engine, channel *Channel, ID string) {
	assertNotEmpty(ID, "Node requires ID.")

	assertNotNil(channel, "Node requires channel.")
	assertNotNil(engine, "Node requires engine.")
}
