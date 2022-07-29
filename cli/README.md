# go-bittorrent-cli

go-bittorrent is a simple Go module for sending informations to BitTorrent trackers.

## Overview

All BitTorrent trackers are supported.

## Build

```shell
go build -ldflags="-s -w" -o bittorrent
```

## Reduce binary size

```shell
upx -9 bittorrent
````

## Usage

### upload

```shell
bittorrent --tracker http://localhost:8080/announce --hash 123456789abcdef123456789abcdef123456789a --upload 4096
```

### download

```shell
bittorrent --tracker http://localhost:8080/announce --hash 123456789abcdef123456789abcdef123456789a --download 4096
```

## License

Copyright © 2022, [Théo BARRAGUÉ](https://www.github.com/theobarrague)

Licensed under the MIT License 2.0
