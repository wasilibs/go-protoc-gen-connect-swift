FROM  ghcr.io/swiftwasm/swift:5.10-focal

RUN apt-get update && apt-get install -y binaryen curl

ADD buildtools/wasm/version.txt version.txt
RUN mkdir -p /workspace && curl -L https://github.com/connectrpc/connect-swift/archive/$(cat version.txt).tar.gz | tar -xz --strip-components 1 -C /workspace
WORKDIR /workspace
ADD buildtools/wasm/patch.txt patch.txt
RUN patch -p1 < patch.txt

RUN swift build -c release --triple wasm32-unknown-wasi --product protoc-gen-connect-swift && swift build -c release --triple wasm32-unknown-wasi --product protoc-gen-connect-swift-mocks

CMD ["cp", ".build/wasm32-unknown-wasi/release/protoc-gen-connect-swift.wasm", ".build/wasm32-unknown-wasi/release/protoc-gen-connect-swift-mocks.wasm", "/out/"]
