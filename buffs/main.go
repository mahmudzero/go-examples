package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	num := 10
	b   := new(bytes.Buffer)
	err := binary.Write(b, binary.LittleEndian, uint16(num))
	if err != nil {
		fmt.Printf("broken w/ e: %+v\n", err)
	}

	fmt.Printf("bytes: %+v\n", b.Bytes())

	buf := make([]byte, 4)
	binary.LittleEndian.PutUint16(buf, uint16(num))
	fmt.Printf("bytes: %+v\n", buf[:3])
}
