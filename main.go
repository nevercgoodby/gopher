package main

import (
	"fmt"
)

func main() {
	var bm Bitmap
	fmt.Println(bm)
	bitmap_test()
	Client_test()
	Client_test1()
	Client_test2()
	Client_connect_server("172.16.10.208", 16379)
	Client_connect()
}
