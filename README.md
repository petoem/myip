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
  -list-stun-servers
        list baked in public stun server
  -n    print newline after the output
  -stun string
        server to use for discovery (e.g. stun:your.server.host:3478)
  -v6
        discover IPv6 address
  -verbose
        verbose output
  -version
        display version information
>
```
