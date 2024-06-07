# go-protoc-gen-connect-swift

go-protoc-gen-connect-swift is a distribution of the [protoc-gen-connect-swift][1].
It does not actually reimplement any functionality of swift-protobuf in Go, instead compiling it
to WebAssembly, and executing with the pure Go Wasm runtime [wazero][2].
This means that `go install` or `go run` can be used to execute it, with no need to rely on external
package managers such as Homebrew, on any platform that Go supports.

## Installation

Precompiled binaries are available in the [releases](https://github.com/wasilibs/go-protoc-gen-connect-swift/releases).
Alternatively, install the plugin using `go install`.

```bash
$ go install github.com/wasilibs/go-protoc-gen-connect-swift/cmd/protoc-gen-connect-swift@latest
```

As long as `$GOPATH/bin`, e.g. `~/go/bin` is on the `PATH`, you can use it with protoc as normal.

```bash
$ protoc --connect-swift_out=out/swift -Iprotos protos/helloworld.proto
```

For [buf][3] users, to avoid installation entirely, it can be convenient to use `go run` in `buf.gen.yaml`.

```yaml
version: v1
plugins:
  - plugin: swift
    out: out/swift
    path:
      [
        "go",
        "run",
        "github.com/wasilibs/go-protoc-gen-connect-swift/cmd/protoc-gen-connect-swift@latest",
      ]
```

This makes it possible to have full protobuf/gRPC generation with no installation of tools,
besides Go itself, on any platform that Go supports. The above examples use `@latest`, but it is
recommended to specify a version, in which case all of the developers on your codebase will use the
same version of the tool with no special steps.

See a full [example][5] in `go-protoc-gen-builtins`. To generate protos, enter the directory and run
`go run github.com/bufbuild/buf/cmd/buf@v1.32.0 generate`. As long as your machine has Go installed,
you will be able to generate protos. The first time using `go run` for a command, Go automatically builds
it making it slower, but subsequent invocations should be quite fast.

[1]: https://github.com/connectrpc/connect-swift
[2]: https://wazero.io/
[3]: https://buf.build/
[4]: https://github.com/wasilibs/go-protoc-gen-builtins
[5]: https://github.com/wasilibs/go-protoc-gen-builtins/tree/main/example
