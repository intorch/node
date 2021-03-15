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
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//funcName function to get the test function Name. It's used to run the test case
//in a subprocess
func funcName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	slice := strings.Split(f.Name(), ".")

	return slice[len(slice)-1]
}

//executeTestInSubprocess function to execute the test case in a subprocess. It is util to
//get any king of errors like os.Exit(n)
func executeTestInSubprocess(testName string, argName string) error {
	cmd := exec.Command(
		os.Args[0],
		fmt.Sprintf("-test.run=%s", testName),
	)
	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("%s=1", argName),
	)

	ex := cmd.Run()

	return ex
}

//assertNotEmptyTestCase test case for assertNotEmpty function
func assertNotEmptyTestCase(funcName string, value string) error {
	argName := strings.ToUpper(strings.Replace(funcName, "Test_", "", 1))

	// Run the crashing code when FLAG is set
	if os.Getenv(argName) == "1" {
		assertNotEmpty(value, "some error message")
		return nil
	}

	return executeTestInSubprocess(funcName, argName)
}

//assertNotNilTestCase test case for assertNotNil
func assertNotNilTestCase(funcName string, value interface{}) error {
	argName := strings.ToUpper(strings.Replace(funcName, "Test_", "", 1))

	print(value)

	// Run the crashing code when FLAG is set
	if os.Getenv(argName) == "1" {
		assertNotNil(value, "some error message")
		return nil
	}

	return executeTestInSubprocess(funcName, argName)
}

//validateTestCase test case for assertNotNil
func validateTestCase(funcName string, engine *Engine, channel *Channel, ID string) error {
	argName := strings.ToUpper(strings.Replace(funcName, "Test_", "", 1))

	// Run the crashing code when FLAG is set
	if os.Getenv(argName) == "1" {
		validate(engine, channel, ID)
		return nil
	}

	return executeTestInSubprocess(funcName, argName)
}

func Test_assertNotEmpty(t *testing.T) {
	assert := assert.New(t)

	someValue := "Some value"
	err := assertNotEmptyTestCase(funcName(), someValue)

	assert.Nil(err)
}

func Test_assertNotEmptyFail(t *testing.T) {
	assert := assert.New(t)

	emptyValue := ""
	err := assertNotEmptyTestCase(funcName(), emptyValue)

	assert.NotNil(err)
	assert.Equal("exit status 1", err.Error())
}

func Test_assertNotNil(t *testing.T) {
	assert := assert.New(t)

	someValue := &Channel{}
	err := assertNotNilTestCase(funcName(), someValue)

	assert.Nil(err)
}

// func Test_assertNotNilError(t *testing.T) {
// 	assert := assert.New(t)

// 	err := assertNotNilTestCase(funcName(), nil)

// 	assert.NotNil(err)
// 	assert.Equal("exit status 1", err.Error())
// }

// func Test_validate(t *testing.T) {
// 	assert := assert.New(t)

// 	enValue := Engine(func(msg message.Message) message.Message {
// 		return msg
// 	})

// 	chValue := &Channel{}

// 	idValue := "some-id"

// 	err := validateTestCase(funcName(), &enValue, chValue, idValue)

// 	assert.Nil(err)
// }
