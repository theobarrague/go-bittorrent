# go-bittorrent

go-bittorrent is a simple Go module for sending informations to BitTorrent trackers.

## Overview

All BitTorrent trackers are supported.

You can build the [bittorrent client tool](cli).

Or you can do development by using this golang module as below.

## Installation

```shell
go get -u github.com/theobarrague/go-bittorrent
```

## Importing

```go
import (
    "github.com/theobarrague/go-bittorrent"
)
```

## Example

### upload

```go
client := bittorrent.New("-DE13F0-ABCDEF", 49152)

err := client.Upload(tracker, hash, bytes)

if err == nil {
    fmt.Println("Uploaded", bytes, "bytes")
}
```

### download

```go
client := bittorrent.New("-DE13F0-ABCDEF", 49152)

err := client.Download(tracker, hash, bytes)

if err == nil {
    fmt.Println("Downloaded", bytes, "bytes")
}
```

## License

Copyright © 2022, [Théo BARRAGUÉ](https://www.github.com/theobarrague)

Licensed under the MIT License
