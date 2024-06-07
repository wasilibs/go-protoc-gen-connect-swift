package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-connect-swift/internal/runner"
	"github.com/wasilibs/go-protoc-gen-connect-swift/internal/wasm"
)

func main() {
	os.Exit(runner.Run("protoc-gen-connect-swift-mocks", os.Args[1:], wasm.ProtocGenConnectSwiftMocks, os.Stdin, os.Stdout, os.Stderr))
}
