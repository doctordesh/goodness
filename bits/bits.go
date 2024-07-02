package bits

import (
	"fmt"
	"strconv"
)

func String(b byte) string {
	return fmt.Sprintf("%08s", strconv.FormatInt(int64(b), 2))
}

func AND(b byte, op byte) string {
	bs := String(b)
	ops := String(op)
	res := String(b & op)
	return fmt.Sprintf("   %8s\n&  %8s\n=  %8s\n", bs, ops, res)
}

func OR(b byte, op byte) string {
	bs := String(b)
	ops := String(op)
	res := String(b | op)
	return fmt.Sprintf("   %8s\n|  %8s\n=  %8s\n", bs, ops, res)
}

func XOR(b byte, op byte) string {
	bs := String(b)
	ops := String(op)
	res := String(b ^ op)
	return fmt.Sprintf("   %8s\n^  %8s\n=  %8s\n", bs, ops, res)
}

func RSHIFT(b byte, op byte) string {
	bs := String(b)
	res := String(b >> op)
	return fmt.Sprintf("   %8s\n>%d %8s\n", bs, op, res)
}

func LSHIFT(b byte, op byte) string {
	bs := String(b)
	res := String(b << op)
	return fmt.Sprintf("   %8s\n<%d %8s\n", bs, op, res)
}
