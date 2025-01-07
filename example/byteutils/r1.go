package main

import (
	"fmt"
	"github.com/hasuburero/util/byteutils"
)

func main() {
	var A []byte = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	var B []byte = []byte{0xff, 0xff, 0xff, 0xff}
	var C []byte = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	int_buf1, err := byteutils.Byte2Int32(A)
	if err != nil {
		fmt.Println(err)
		fmt.Println("byteutils.Byte2Int32 error\n")
	} else {
		fmt.Println(int_buf1)
		fmt.Println("")
	}

	int_buf2, err := byteutils.Byte2Int64(A)
	if err != nil {
		fmt.Println(err)
		fmt.Println("byteutils.Byte2Int64 error\n")
	} else {
		fmt.Println(int_buf2)
		fmt.Println("")
	}

	int_buf3, err := byteutils.Byte2Int32(B)
	if err != nil {
		fmt.Println(err)
		fmt.Println("byteutils.Byte2Int32 error\n")
	} else {
		fmt.Println(int_buf3)
		fmt.Println("")
	}

	int_buf4, err := byteutils.Byte2Int64(B)
	if err != nil {
		fmt.Println(err)
		fmt.Println("byteutils.Byte2Int64 error\n")
	} else {
		fmt.Println(int_buf4)
		fmt.Println("")
	}

	int_buf5, err := byteutils.Byte2Int64(C)
	if err != nil {
		fmt.Println(err)
		fmt.Println("byteutils.Byte2Int64 error\n")
	} else {
		fmt.Println(int_buf5)
		fmt.Println("")
	}
	int_buf6, err := byteutils.Byte2Int32(C)
	if err != nil {
		fmt.Println(err)
		fmt.Println("byteutils.Byte2Int64 error\n")
	} else {
		fmt.Println(int_buf6)
		fmt.Println("")
	}

	var a int32 = 1
	var b int64 = 2
	byte_buf1, err := byteutils.Int322Byte(a)
	if err != nil {
		fmt.Println(err)
		fmt.Println("byteutils.Int322Byte error")
	} else {
		byteutils.PrintByte(byte_buf1)
	}
	byte_buf2, err := byteutils.Int642Byte(b)
	if err != nil {
		fmt.Println(err)
		fmt.Println("byteutils.Int642Byte error")
	} else {
		byteutils.PrintByte(byte_buf2)
	}

	return
}
