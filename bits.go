package goodness

import (
	"fmt"
	"strconv"
)

func bits(b byte) string {
	return fmt.Sprintf("%08s", strconv.FormatInt(int64(b), 2))
}

func bitsAND(b byte, op byte) string {
	bs := bits(b)
	ops := bits(op)
	res := bits(b & op)
	return fmt.Sprintf("   %8s\n+  %8s\n=  %8s\n", bs, ops, res)
}

func bitsOR(b byte, op byte) string {
	bs := bits(b)
	ops := bits(op)
	res := bits(b | op)
	return fmt.Sprintf("   %8s\n|  %8s\n=  %8s\n", bs, ops, res)
}

func bitsXOR(b byte, op byte) string {
	bs := bits(b)
	ops := bits(op)
	res := bits(b ^ op)
	return fmt.Sprintf("   %8s\n^  %8s\n=  %8s\n", bs, ops, res)
}

func bitsRSHIFT(b byte, op byte) string {
	bs := bits(b)
	res := bits(b >> op)
	return fmt.Sprintf("   %8s\n>%d %8s\n", bs, op, res)
}

func bitsLSHIFT(b byte, op byte) string {
	bs := bits(b)
	res := bits(b << op)
	return fmt.Sprintf("   %8s\n<%d %8s\n", bs, op, res)
}
