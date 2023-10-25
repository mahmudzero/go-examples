package main

import "fmt"

func mut(buf []byte) {
	fmt.Printf("buf: %+v\n", buf)
	buf[0] = 0xb
}

func main() {
	var buf []byte = make([]byte, 10)
	buf[0] = 0x1
	buf[1] = 0x2
	buf[2] = 0x3
	buf[3] = 0x4
	buf[4] = 0x5
	buf[5] = 0x6
	buf[6] = 0x7
	buf[7] = 0x8
	buf[8] = 0x9
	buf[9] = 0xa

	fmt.Printf("buf: %+v\n", buf)

	mut(buf[:5])

	fmt.Printf("buf: %+v\n", buf)
}
