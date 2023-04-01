# ArvanCloud CDN Go

[![Release](https://github.com/arvancloud/cdn-go/actions/workflows/release.yaml/badge.svg)](https://github.com/arvancloud/cdn-go/actions/workflows/release.yaml) [![CodeQL](https://github.com/arvancloud/cdn-go/actions/workflows/codeql.yaml/badge.svg)](https://github.com/arvancloud/cdn-go/actions/workflows/codeql.yaml) ![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/r1cloud/cdn?sort=semver) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/arvancloud/cdn-go) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/arvancloud/cdn-go?display_name=tag&sort=semver)

![logo](.github/logo.svg)

It's a Go library for interacting with the ArvanCloud CDN API.

> **Note**: This project is under active development and may have problems and shortcomings.

## Installation

```bash
go get github.com/arvancloud/cdn-go
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    arvancloud "github.com/arvancloud/cdn-go/pkg"
)

func main() {
    api, err := arvancloud.New(
    os.Getenv("ARVANCLOUD_API_KEY"),
        arvancloud.Debug(false),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    resource := arvancloud.Resource{
        Domain: "test.ir",
    }

    record := arvancloud.CreateDNSRecordParams{
        Type: "A",
        Name: "C",
        Value: []arvancloud.DNSRecord_Value_A{
            {
                IP: "1.1.1.1",
            },
        },
        TTL:           120,
        UpstreamHTTPS: "default",
        IPFilterMode: arvancloud.DNSRecord_IPFilterMode{
            Count:     "single",
            Order:     "none",
            GeoFilter: "none",
        },
    }

    u, err := api.CreateDNSRecord(ctx, resource, record)
    if err != nil {
        log.Fatal(err.Error())
    }

    fmt.Printf("%+v", u)
}
```

## Features

- [ ] Products
  - [x] DNS
  - [ ] Firewall
  - [ ] WAF
  - [ ] DDoS
  - [ ] Cache
  - [ ] Load Balancer
  - [ ] Page Rule
  - [ ] Acceleration
  - [ ] Custom Pages
  - [ ] Redirect
  - [ ] Log Forwarder
  - [ ] L4 Proxy
  - [ ] Rate Limit
- [ ] Package
  - [x] CLI
  - [ ] Documentation
  - [x] Official Release
  - [x] CI/CD

## Contributing

We welcome contributions from the community. Please consider that this project is under active development as we expand it to cover the CDN API.

Please report any issues you find in the [Issues page](https://github.com/arvancloud/cdn-go/issues) or send us an email at [cdn@arvancloud.ir](mailto:cdn@arvancloud.ir).
