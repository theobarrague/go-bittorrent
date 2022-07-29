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

package bittorrent

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type Client struct {
	peerId string
	port   int
}

func New(peerId string, port int) *Client {
	client := Client{peerId: peerId, port: port}
	return &client
}

func (client Client) Upload(tracker string, hash string, bytes int) error {
	encoded, err := encodeHash(hash)

	if err != nil {
		return err
	}

	request := []string{tracker, "?info_hash=", encoded, "&peer_id=", client.peerId, "&port=", strconv.Itoa(client.port), "&uploaded=", strconv.Itoa(bytes), "&downloaded=", strconv.Itoa(0), "&compact=1"}

	result, err := http.Get(strings.Join(request, ""))

	if err != nil {
		return err
	} else if result.StatusCode != http.StatusFound {
		return errors.New("http status code " + strconv.Itoa(result.StatusCode))
	}

	return nil
}

func (client Client) Download(tracker string, hash string, bytes int) error {
	encoded, err := encodeHash(hash)

	if err != nil {
		return err
	}

	request := []string{tracker, "?info_hash=", encoded, "&peer_id=", client.peerId, "&port=", strconv.Itoa(client.port), "&uploaded=", strconv.Itoa(0), "&downloaded=", strconv.Itoa(bytes), "&compact=1"}

	result, err := http.Get(strings.Join(request, ""))

	if err != nil {
		return err
	} else if result.StatusCode != http.StatusFound {
		return errors.New("http status code " + strconv.Itoa(result.StatusCode))
	}

	return nil
}

func encodeHash(hash string) (string, error) {
	var encoded strings.Builder

	for i := 0; i < len(hash); i += 2 {
		byte := hash[i : i+2]
		v, err := strconv.ParseInt(byte, 16, 16)

		if err != nil {
			return "", err
		}

		if v >= 48 && v <= 57 { // 0-9
			_, err = encoded.WriteString(string(rune(v)))
		} else if v >= 65 && v <= 90 { // A-Z
			_, err = encoded.WriteString(string(rune(v)))
		} else if v >= 97 && v <= 122 { // a-z
			_, err = encoded.WriteString(string(rune(v)))
		} else if v == 45 || v == 46 || v == 95 || v == 126 { // dot, dash, underscore, tilde
			_, err = encoded.WriteString(string(rune(v)))
		} else {
			_, err = encoded.WriteString("%" + byte)
		}

		if err != nil {
			return "", err
		}
	}

	return encoded.String(), nil
}
