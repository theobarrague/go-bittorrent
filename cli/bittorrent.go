/*
 * Copyright © 2022, Théo BARRAGUÉ
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the “Software”), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * The Software is provided “as is”, without warranty of any kind, express or
 * implied, including but not limited to the warranties of merchantability, fitness
 * for a particular purpose and noninfringement. In no event shall the authors or
 * copyright holders X be liable for any claim, damages or other liability, whether
 * in an action of contract, tort or otherwise, arising from, out of or in
 * connection with the software or the use or other dealings in the Software.
 *
 * Except as contained in this notice, the name of the Théo BARRAGUÉ shall not be
 * used in advertising or otherwise to promote the sale, use or other dealings in
 * this Software without prior written authorization from the Théo BARRAGUÉ.
 */

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/theobarrague/go-bittorrent"
)

func main() {
	var hash string
	var tracker string
	var upload int
	var download int

	flag.StringVar(&hash, "hash", "", "torrent hash")
	flag.StringVar(&tracker, "tracker", "", "tracker url")
	flag.IntVar(&upload, "upload", 0, "bytes")
	flag.IntVar(&download, "download", 0, "bytes")

	flag.Parse()

	if hash == "" {
		fmt.Fprintln(os.Stderr, "set hash")
		os.Exit(1)
	}

	if tracker == "" {
		fmt.Fprintln(os.Stderr, "set tracker")
		os.Exit(1)
	}

	if upload < 0 || download < 0 {
		fmt.Fprintln(os.Stderr, "set upload and / or download with positive values")
		os.Exit(1)
	}

	if upload == 0 && download == 0 {
		fmt.Fprintln(os.Stderr, "set upload or download")
		os.Exit(1)
	}

	peerId := randomPeerId()
	port := randomPort()

	client := bittorrent.New(peerId, port)

	if upload > 0 {
		err := client.Upload(tracker, hash, upload)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Println("Uploaded", upload, "bytes")
		}
	}

	if download > 0 {
		err := client.Download(tracker, hash, download)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Println("Downloaded", upload, "bytes")
		}
	}
}

func randomPort() int {
	return rand.Intn(4096) + 49152
}

func randomPeerId() string {
	peerId := "-DE13F0-"

	for i := 0; i < 6; i++ {
		peerId += string(rune(rand.Intn(26) + 65))
	}

	return peerId
}
