package main

import (
	_ "embed"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func bit(b []byte, i uint32) (bool, uint32) {
	idx, offset := (i / 8), (i % 8)
	return (b[idx] & (1 << uint(7-offset))) != 0, i + 1
}

func extract(bytes []byte, bit_start uint32, length uint32) (uint64, uint32) {
	if length > 64 {
		log.Fatal("length too long: ", length)
	}
	var res uint64 = 0
	for i := bit_start; i < bit_start+length; i++ {
		res = res << 1
		b, _ := bit(bytes, i)
		if b {
			res += 1
		}
	}
	return res, bit_start + length
}

// pattern explained here: https://eli.thegreenplace.net/2018/go-and-algebraic-data-types/
type Packet interface {
	isTree()
}

type Header struct {
	version uint8
	id      uint8
}

type Value struct {
	Header
	value uint64
}

type Operator struct {
	Header
	packets []Packet
}

func (_ Value) isTree()    {}
func (_ Operator) isTree() {}

func sumVersion(t Packet) int {
	switch nt := t.(type) {
	case Value:
		return int(nt.version)
	case Operator:
		res := int(nt.version)
		for _, p := range nt.packets {
			res += sumVersion(p)
		}
		return res
	default:
		log.Fatalf("unexpected type %T", nt)
	}
	return 0
}

func eval(t Packet) int {
	switch nt := t.(type) {
	case Value:
		return int(nt.value)
	case Operator:
		switch nt.id {
		case 0:
			res := 0
			for _, p := range nt.packets {
				res += eval(p)
			}
			return res
		case 1:
			res := 1
			for _, p := range nt.packets {
				res *= eval(p)
			}
			return res
		case 2:
			res := math.MaxInt
			for _, p := range nt.packets {
				a := eval(p)
				if a < res {
					res = a
				}
			}
			return res
		case 3:
			res := 0
			for _, p := range nt.packets {
				a := eval(p)
				if a > res {
					res = a
				}
			}
			return res
		case 5:
			return toInt(eval(nt.packets[0]) > eval(nt.packets[1]))
		case 6:
			return toInt(eval(nt.packets[0]) < eval(nt.packets[1]))
		case 7:
			return toInt(eval(nt.packets[0]) == eval(nt.packets[1]))
		}
		return 0

	default:
		log.Fatalf("unexpected type %T", nt)
	}
	return 0
}

func toInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func (h Header) String() string {
	return fmt.Sprintf("version:%d id:%d", h.version, h.id)
}
func (l Value) String() string {
	return fmt.Sprintf("LV: %s value=%d", l.Header, l.value)
}
func (o Operator) String() string {
	return fmt.Sprintf("op: %s packets=%v", o.Header, o.packets)
}

func decode(bytes []byte, bit_start uint32) (packet Packet, index uint32) {
	version, index := extract(bytes, uint32(bit_start), 3)
	id, index := extract(bytes, index, 3)
	header := Header{version: uint8(version), id: uint8(id)}

	if id == 4 {
		var next bool
		var value uint64

		next, index = bit(bytes, index)
		value, index = extract(bytes, index, 4)
		for next {
			next, index = bit(bytes, index)
			var v uint64
			v, index = extract(bytes, index, 4)
			value = value<<4 + v
		}
		packet = Value{
			Header: header,
			value:  value,
		}
		return
	} else {
		var lengthTypeId bool

		var packets []Packet
		lengthTypeId, index = bit(bytes, index)
		if !lengthTypeId {
			var length uint64
			length, index = extract(bytes, index, 15)
			end := index + uint32(length)
			for index < end {
				var res Packet
				res, index = decode(bytes, index)
				packets = append(packets, res)
			}
		} else {
			var length uint64
			length, index = extract(bytes, index, 11)
			for i := 0; i < int(length); i++ {
				var res Packet
				res, index = decode(bytes, index)
				packets = append(packets, res)
			}
		}
		packet = Operator{
			Header:  header,
			packets: packets,
		}
		return
	}
}

func decodeString(s string) Packet {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	res, _ := decode(bytes, 0)
	return res
}

func Part1(input string) int {
	s := strings.TrimSuffix(input, "\n")
	res := decodeString(s)
	return sumVersion(res)
}

func Part2(input string) int {
	s := strings.TrimSuffix(input, "\n")
	res := decodeString(s)
	return eval(res)
}

func main() {
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
