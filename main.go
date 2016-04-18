package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func test() {

	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Print(pi)
}
func main() {
	var (
		i    uint32
		size uint32
	)
	size = 512
	bmap := new(Bitmap)
	bmap.Init(size)
	fmt.Println(bmap)

	for i = 0; i < size*6; i++ {
		bmap.SetBit(uint32(i), 1)
	}
	fmt.Println(bmap)
	fmt.Println(bmap.GetBit(size * 5))
	fmt.Println(bmap.GetBit(size * 7))
	test()
}
