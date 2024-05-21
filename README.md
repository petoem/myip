# myip

Tool to discover your public IP address with STUN (Session Traversal Utilities for NAT)

## Install

```console
go install github.com/petoem/myip/cmd/myip@latest
```

## Usage

```console
> myip --help
Usage of myip:
  -n    print newline after IP address
  -stun string
        server to use for discovery (e.g. stun:your.server.host:3478)
  -verbose
        verbose output
  -version
        display version information
>
```
