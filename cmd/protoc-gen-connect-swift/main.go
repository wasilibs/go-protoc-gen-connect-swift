package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-connect-swift/internal/runner"
	"github.com/wasilibs/go-protoc-gen-connect-swift/internal/wasm"
)

func main() {
	os.Exit(runner.Run("protoc-gen-connect-swift", os.Args[1:], wasm.ProtocGenConnectSwift, os.Stdin, os.Stdout, os.Stderr))
}
