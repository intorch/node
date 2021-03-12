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
	"bytes"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_assertNotEmpty(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	log.SetOutput(&buf)

	assertNotEmpty("Hello World", "It should not be printed")

	msg := buf.String()
	assert.Empty(msg)
}

func Test_assertNotEmptyFail(t *testing.T) {
	//assert := assert.New(t)

	// Run the crashing code when FLAG is set
	if os.Getenv("FLAG") == "1" {
		assertNotEmpty("", "Should be printed")
		return
	}

	// Run the test in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=TestGetConfig")
	cmd.Env = append(os.Environ(), "FLAG=1")
	err := cmd.Run()

	// Cast the error as *exec.ExitError and compare the result
	//e, ok := err.(*exec.ExitError)

	println(err.Error())
	//	print(ok)

	// expectedErrorString := "exit status 1"
	// assert.Equal(true, ok)
	// assert.Equal(expectedErrorString, e.Error())
}

// assert := assert.New(t)

// eng := Engine(func(msg message.Message) message.Message {
// 	return msg
// })

// cha := &Channel{}
