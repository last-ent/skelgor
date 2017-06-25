package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type FileMeta struct {
	FileName string
	FileData []byte
}

var (
	HelpersTemplate = []byte(`
package common

//Empty for now.
`)

	TestHelpersTemplate = []byte(`
package common

import "testing"

const checkMark = "\u2713"
const ballotX = "\u2717"
const prefix = "\t\t - "

// OnTestSuccess is for pass.
func OnTestSuccess(t *testing.T, msg string){
	t.Log(prefix+msg, checkMark)
}

//OnTestError is for fail.
func OnTestError(t *testing.T, msg string) {
	t.Error(prefix+msg, ballotX)
}

//OnTestUnexpectedError is for unexpected fail.
func OnTestUnexpectedError(t *testing.T, err error) {
	OnTestError(t, "Unexpected Error:\n"+err.Error())
}

`)

	MainTemplate = []byte(`
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}
`)

	MakefileTemplate = []byte(`
IMAGE = docker-image-name
GOPKGS = $(shell go list ./...)
GOOS = linux
GOARCH = amd64
CGO_ENABLED = 0
BUILD_PATH = ./build
BINARY_NAME = "exec-bin"
SRC_ROOT = "."


test: lint
	@echo "RUNNING TESTS\n"
	go get ./...
	go test --cover -v $(GOPKGS)

lint:
	@echo "COMMENCING LINT CHECKS."
	go fmt ./...
	golint ./...
	go vet ./...

run.docker: test build.linux
	docker build -t $(IMAGE) .
	docker run --rm -it --name=$(IMAGE) $(IMAGE)

build.linux:
	@echo "BUILDING LINUX BINARY"
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) \
	go build -o $(BUILD_PATH)/lin-$(BINARY_NAME) $(SRC_ROOT)

build.bin:
	@echo "BUILDING BINARY"
	go build -o $(BUILD_PATH)/$(BINARY_NAME) $(SRC_ROOT)

run: build.bin
	./$(BUILD_PATH)/$(BINARY_NAME)

`)
)

func MakeProject(path string) error {
	var err error

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	if err = os.MkdirAll(path+"common", 0777); err != nil {
		fmt.Println(err)
		return err
	}

	var filesToWrite = []FileMeta{
		FileMeta{"main.go", MainTemplate},
		FileMeta{"Makefile", MakefileTemplate},
		FileMeta{"common/helpers.go", HelpersTemplate},
		FileMeta{"common/test_helpers.go", TestHelpersTemplate},
	}

	for _, fileMeta := range filesToWrite {
		ioutil.WriteFile(path+fileMeta.FileName, fileMeta.FileData, 0777)

	}

	return err
}

func main() {
	var dir string
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir = "tmp"
	}

	fmt.Println("We will now try to create project dir:", dir)

	if err := MakeProject(dir); err == nil {
		fmt.Println("Project created")
	} else {
		fmt.Println("Error while trying to create project:", err)
	}

}
