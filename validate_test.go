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

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	slice := strings.Split(f.Name(), ".")

	return slice[len(slice)-1]
}

func executeTestInSubprocess(testName string) error {
	cmd := exec.Command(
		os.Args[0],
		fmt.Sprintf("-test.run=%s", testName),
	)
	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("%s=1", testName),
	)

	return cmd.Run()
}

func Test_assertNotEmpty(t *testing.T) {
	assert := assert.New(t)

	name := runFuncName()
	print(name)

	// Run the crashing code when FLAG is set
	if os.Getenv("ASSERT_NOT_EMPTY") == "1" {
		assertNotEmpty("tst", "some error message")
		return
	}

	// Run the test in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=Test_assertNotEmpty")
	cmd.Env = append(os.Environ(), "ASSERT_NOT_EMPTY=1")
	err := cmd.Run()

	assert.Nil(err)
}

func Test_assertNotEmptyFail(t *testing.T) {
	assert := assert.New(t)

	// Run the crashing code when FLAG is set
	if os.Getenv("ASSERT_NOT_EMPTY_FAIL") == "1" {
		assertNotEmpty("", "some error message")
		return
	}

	// Run the test in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=Test_assertNotEmptyFail")
	cmd.Env = append(os.Environ(), "ASSERT_NOT_EMPTY_FAIL=1")
	err := cmd.Run()

	assert.NotNil(err)
	assert.Equal("exit status 1", err.Error())
}

// func Test_assertNotNil(t *testing.T) {
// 	assert := assert.New(t)

// 	// Run the crashing code when FLAG is set
// 	if os.Getenv("ASSERT_NOTnIL") == "1" {
// 		print("run")
// 		assertNotNil(Channel{}, "some error message")
// 		return
// 	}

// func Test_assertNotNilFail(t *testing.T) {
// 	assert := assert.New(t)

// 	// Run the crashing code when FLAG is set
// 	if os.Getenv("FLAG") == "1" {
// 		print("run")
// 		assertNotNil(nil, "some error message")
// 		return
// 	}

// 	// Run the test in a subprocess
// 	cmd := exec.Command(os.Args[0], "-test.run=Test_assertNotNilFail")
// 	cmd.Env = append(os.Environ(), "FLAG=1")
// 	err := cmd.Run()

// 	assert.NotNil(err)
// 	assert.Equal("exit status 1", err.Error())
// }

// 	// Run the test in a subprocess
// 	cmd := exec.Command(os.Args[0], "-test.run=Test_assertNotNil")
// 	cmd.Env = append(os.Environ(), "FLAG=1")
// 	err := cmd.Run()

// 	assert.Nil(err)
// }

// func Test_validate(t *testing.T) {
// 	assert := assert.New(t)

// 	// Run the crashing code when FLAG is set
// 	if os.Getenv("FLAG") == "1" {
// 		print("run")
// 		validate(&ThisIsJustAnExample, &Channel{}, "test")
// 		return
// 	}

// 	// Run the test in a subprocess
// 	cmd := exec.Command(os.Args[0], "-test.run=Test_validate")
// 	cmd.Env = append(os.Environ(), "FLAG=1")
// 	err := cmd.Run()

// 	assert.Nil(err)
// }

// func Test_validateErrEngine(t *testing.T) {
// 	assert := assert.New(t)

// 	// Run the crashing code when FLAG is set
// 	if os.Getenv("FLAG") == "1" {
// 		print("run")
// 		validate(nil, &Channel{}, "test")
// 		return
// 	}

// 	// Run the test in a subprocess
// 	cmd := exec.Command(os.Args[0], "-test.run=Test_validateErrEngine")
// 	cmd.Env = append(os.Environ(), "FLAG=1")
// 	err := cmd.Run()

// 	assert.NotNil(err)
// 	assert.Equal("exit status 1", err.Error())
// }

// func Test_validateErrChannel(t *testing.T) {
// 	assert := assert.New(t)

// 	// Run the crashing code when FLAG is set
// 	if os.Getenv("FLAG") == "1" {
// 		print("run")
// 		validate(&ThisIsJustAnExample, nil, "test")
// 		return
// 	}

// 	// Run the test in a subprocess
// 	cmd := exec.Command(os.Args[0], "-test.run=Test_validateErrChannel")
// 	cmd.Env = append(os.Environ(), "FLAG=1")
// 	err := cmd.Run()

// 	assert.NotNil(err)
// 	assert.Equal("exit status 1", err.Error())
// }

// func Test_validateErrID(t *testing.T) {
// 	assert := assert.New(t)

// 	// Run the crashing code when FLAG is set
// 	if os.Getenv("FLAG") == "1" {
// 		print("run")
// 		validate(&ThisIsJustAnExample, &Channel{}, "")
// 		return
// 	}

// 	// Run the test in a subprocess
// 	cmd := exec.Command(os.Args[0], "-test.run=Test_validateErrID")
// 	cmd.Env = append(os.Environ(), "FLAG=1")
// 	err := cmd.Run()

// 	assert.NotNil(err)
// 	assert.Equal("exit status 1", err.Error())
// }
