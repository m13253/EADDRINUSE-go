/*
  EADDRINUSE-go
  Written in 2018, by Star Brilliant <coder@poorlab.com>

  This is free and unencumbered software released into the public domain.

  Anyone is free to copy, modify, publish, use, compile, sell, or
  distribute this software, either in source code form or as a compiled
  binary, for any purpose, commercial or non-commercial, and by any
  means.

  In jurisdictions that recognize copyright laws, the author or authors
  of this software dedicate any and all copyright interest in the
  software to the public domain. We make this dedication for the benefit
  of the public at large and to the detriment of our heirs and
  successors. We intend this dedication to be an overt act of
  relinquishment in perpetuity of all present and future rights to this
  software under copyright law.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
  IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
  OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
  ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
  OTHER DEALINGS IN THE SOFTWARE.

  For more information, please refer to <http://unlicense.org>
*/

package EADDRINUSE_test

import (
	"fmt"
	"net"

	EADDRINUSE "github.com/m13253/EADDRINUSE-go"
)

func Example() {
	l1, addr1, err := listenOnNextAddr("127.0.0.1:10000")
	if err != nil {
		panic(err)
	}
	defer l1.Close()
	l2, addr2, err := listenOnNextAddr("[::1]:10000")
	if err != nil {
		panic(err)
	}
	defer l2.Close()
	l3, addr3, err := listenOnNextAddr("127.0.0.1:10000")
	if err != nil {
		panic(err)
	}
	defer l3.Close()
	l4, addr4, err := listenOnNextAddr("[::1]:10000")
	if err != nil {
		panic(err)
	}
	defer l4.Close()
	fmt.Println(addr1)
	fmt.Println(addr2)
	fmt.Println(addr3)
	fmt.Println(addr4)
	// Output:
	// 127.0.0.1:10000
	// [::1]:10000
	// 127.0.0.1:10001
	// [::1]:10001
}

func listenOnNextAddr(addr string) (net.Listener, *net.TCPAddr, error) {
	originalAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, nil, err
	}

	var l net.Listener
	availableAddr := new(net.TCPAddr)
	*availableAddr = *originalAddr

	for availableAddr.Port = originalAddr.Port; availableAddr.Port < 65535 && availableAddr.Port-originalAddr.Port < 10; availableAddr.Port++ {
		l, err = net.ListenTCP("tcp", availableAddr)
		if err != nil {
			if EADDRINUSE.Tell(err) {
				continue
			} else {
				return nil, nil, err
			}
		}
		return l, availableAddr, nil
	}
	return nil, nil, err
}
