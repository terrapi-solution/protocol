# TerrAPI Protocol

## Overview

Terrapi Protocol is a gRPC-based protocol in the TerrAPI system.

## Prerequisites

- Go 1.16 or later
- Protocol Buffers compiler (`protoc`)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

## Installation

Install the Protocol Buffers compiler (`protoc`):

```sh
# On Ubuntu
sudo apt-get install -y protobuf-compiler

# On macOS
brew install protobuf
```

Install the Go plugins for Protocol Buffers:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Ensure the `protoc-gen-go` and `protoc-gen-go-grpc` binaries are in your PATH:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

Generating Code

To generate the Go code from the Protocol Buffers definitions, run the following script:

On Unix-based systems:

```sh
./gen.sh
```

On Windows:

```sh
gen.cmd
```

## Security

If you discover a security vulnerability within this package, please send an e-mail to <contact@thomas-illiet.fr>.

## Contributing

Fork -> Patch -> Push -> Pull Request

## License

This project is licensed under the Apache License, Version 2.0.
See the [LICENSE](LICENSE) file for details.
