package wasm

import _ "embed"

//go:embed protoc-gen-connect-swift.wasm
var ProtocGenConnectSwift []byte

//go:embed protoc-gen-connect-swift-mocks.wasm
var ProtocGenConnectSwiftMocks []byte
