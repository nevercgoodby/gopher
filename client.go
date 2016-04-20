package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func Client_connect_server(host string, port int16) net.Conn {

	addr := net.ParseIP(host)
	if addr == nil {
		fmt.Println("Invalid address")
		return nil
	} else {
		fmt.Println("Connect to: ", addr.String())
	}

	return nil
}

func Client_connect() {

	conn, err := net.DialTimeout("tcp4", "172.16.10.208:22", time.Second*10)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	var deadtime time.Time
	conn.SetDeadline(deadtime)
	var buf []byte
	buf = make([]byte, 128)
	wn, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	rn, err := conn.Read(buf)
	fmt.Println(string(buf), err, wn, rn)
}

func Client_download() {

}

func Client_DailRedirect() {

}

func Client_test() {

	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(b)
	fmt.Println("BUF:", buf)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Print(pi, buf)
}

func Client_test1() {
	var pi float64 = 3.141592653589793
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, pi)
	// 这里可以继续往buf里写, 都存在buf里
	// binary.Write(buf, binary.LittleEndian, uint16(12345))
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Print(buf.Bytes())
	// [24 45 68 84 251 33 9 64]

}
func Client_test2() {
	buf := new(bytes.Buffer)
	var data = []interface{}{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	fmt.Printf("XXX%x\n", buf.Bytes())
	// beefcafe 这个是16进制串
	// 这里转换为了16进制整数的串?
}
