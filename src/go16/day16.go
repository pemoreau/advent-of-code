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
var input_day string

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
		b, _ := bit(bytes, i)
		if b {
			res = 2*res + 1
		} else {
			res = res * 2
		}
	}
	return res, bit_start + length
}

type Header struct {
	version uint8
	id      uint8
}

type LitteralValue struct {
	Header
	value int64
}

type Operator struct {
	Header
	packets []Packet
}

type Packet interface {
	sumVersion() int
	eval() int
}

func (l LitteralValue) sumVersion() int {
	return int(l.Header.version)
}

func (o Operator) sumVersion() int {
	res := int(o.Header.version)
	for _, p := range o.packets {
		res += p.sumVersion()
	}
	return res
}

func (l LitteralValue) eval() int {
	return int(l.value)
}

func (o Operator) eval() int {
	switch o.id {
	case 0:
		res := 0
		for _, p := range o.packets {
			res += p.eval()
		}
		return res
	case 1:
		res := 1
		for _, p := range o.packets {
			res *= p.eval()
		}
		return res
	case 2:
		res := math.MaxInt
		for _, p := range o.packets {
			a := p.eval()
			if a < res {
				res = a
			}
		}
		return res
	case 3:
		res := 0
		for _, p := range o.packets {
			a := p.eval()
			if a > res {
				res = a
			}
		}
		return res
	case 5:
		return toInt(o.packets[0].eval() > o.packets[1].eval())
	case 6:
		return toInt(o.packets[0].eval() < o.packets[1].eval())
	case 7:
		return toInt(o.packets[0].eval() == o.packets[1].eval())
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
func (l LitteralValue) String() string {
	return fmt.Sprintf("LV: %s value=%d", l.Header, l.value)
}
func (o Operator) String() string {
	return fmt.Sprintf("Op: %s packets=%v", o.Header, o.packets)
}

func decode(bytes []byte, bit_start uint32) (Packet, uint32) {
	var index uint32 = bit_start
	version, index := extract(bytes, uint32(index), 3)
	id, index := extract(bytes, index, 3)
	header := Header{version: uint8(version), id: uint8(id)}

	if id == 4 {
		var next bool
		var res uint64

		next, index = bit(bytes, index)
		res, index = extract(bytes, index, 4)
		for next {
			next, index = bit(bytes, index)
			var v uint64
			v, index = extract(bytes, index, 4)
			res = res<<4 + v
		}
		return LitteralValue{
			Header: header,
			value:  int64(res),
		}, index
	} else {
		var lengthTypeId bool
		var length uint64

		lengthTypeId, index = bit(bytes, index)
		if !lengthTypeId {
			length, index = extract(bytes, index, 15)
			end := index + uint32(length)
			packets := make([]Packet, 0)
			for index < end {
				var res Packet
				res, index = decode(bytes, index)
				packets = append(packets, res)
			}
			return Operator{
				Header:  header,
				packets: packets,
			}, index
		} else {
			length, index = extract(bytes, index, 11)
			packets := make([]Packet, length)
			for i := 0; i < int(length); i++ {
				var res Packet
				res, index = decode(bytes, index)
				packets[i] = res
			}
			return Operator{
				Header:  header,
				packets: packets,
			}, index

		}
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
	return res.sumVersion()
}

func Part2(input string) int {
	s := strings.TrimSuffix(input, "\n")
	res := decodeString(s)
	return res.eval()
}

func main() {
	start := time.Now()
	fmt.Println("part1: ", Part1(string(input_day)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input_day)))
	fmt.Println(time.Since(start))
}
