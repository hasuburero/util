package byteutils

import (
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	shift_bits = 8
	bytesize   = 256
)

func ByteSize(arg uint) int {
	var i int = 0
	for ; i < 4; i++ {
		arg = arg / bytesize
		if arg == 0 {
			break
		}
	}
	return i + 1
}

func Byte2Int32(arg1 []byte) (int32, error) {
	var result int32
	arglen := len(arg1)
	if arglen > 4 {
		return 0, errors.New("out of range int32")
	}
	byte_buf := []byte{}
	for i := 0; i < 4-arglen; i++ {
		byte_buf = append(byte_buf, 0x00)
	}
	byte_buf = append(byte_buf, arg1...)
	result = int32(binary.BigEndian.Uint32(byte_buf))
	return result, nil
}

func Byte2Int64(arg1 []byte) (int64, error) {
	var result int64 = 0
	arglen := len(arg1)
	if arglen > 8 {
		return 0, errors.New("out of range int64")
	}
	byte_buf := []byte{}
	for i := 0; i < 8-arglen; i++ {
		byte_buf = append(byte_buf, 0x00)
	}
	byte_buf = append(byte_buf, arg1...)
	result = int64(binary.BigEndian.Uint64(byte_buf))
	return result, nil
}

func Int322Byte(arg2 int32) ([]byte, error) {
	result := make([]byte, 4)
	binary.BigEndian.PutUint32(result, uint32(arg2))
	return result, nil
}

func Int642Byte(arg2 int64) ([]byte, error) {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, uint64(arg2))
	return result, nil
}

func Bytecmp(arg1, arg2 []byte) bool {
	len1 := len(arg1)
	len2 := len(arg2)
	if len1 != len2 {
		return false
	}
	for i := 0; i < len1; i++ {
		if arg1[i] != arg2[i] {
			return false
		}
	}

	return true
}

func PrintByte(arg1 []byte) {
	for _, ctx := range arg1 {
		fmt.Printf("%x ", ctx)
	}
	fmt.Println("")
}
