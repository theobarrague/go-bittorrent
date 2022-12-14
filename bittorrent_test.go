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
	"fmt"
	"testing"
)

func TestHashEncoding(t *testing.T) {
	encoded, err := encodeHash("123456789abcdef123456789abcdef123456789a")

	fmt.Println(encoded)

	if err != nil {
		t.FailNow()
	}

	if encoded != "%124Vx%9a%bc%de%f1%23Eg%89%ab%cd%ef%124Vx%9a" {
		t.FailNow()
	}
}
