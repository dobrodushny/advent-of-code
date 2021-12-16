package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

const VALUE_PACKET_TYPE = 4

type Packet interface {
	Version() int
	Type() int
	Value() int
}

type ValuePacket struct {
	v   int
	t   int
	val int
}

func (vp *ValuePacket) Version() int {
	return vp.v
}

func (vp *ValuePacket) Type() int {
	return vp.t
}

func (vp *ValuePacket) Value() int {
	return vp.val
}

type OperatorPacket struct {
	v          int
	t          int
	SubPackets []*Packet
}

func (op *OperatorPacket) Version() int {
	return op.v
}

func (op *OperatorPacket) Type() int {
	return op.t
}

func (op *OperatorPacket) Value() int {
	var res int

	switch op.Type() {
	case 0:
		for _, sub := range op.SubPackets {
			res += (*sub).Value()
		}
	case 1:
		if res == 0 {
			res = 1
		}
		for _, sub := range op.SubPackets {
			res *= (*sub).Value()
		}
	case 2:
		min := 999999999999
		for _, sub := range op.SubPackets {
			val := (*sub).Value()
			if val < min {
				min = val
			}
		}
		res = min
	case 3:
		max := 0
		for _, sub := range op.SubPackets {
			val := (*sub).Value()
			if val > max {
				max = val
			}
		}
		res = max
	case 5:
		first := op.SubPackets[0]
		second := op.SubPackets[1]
		if (*first).Value() > (*second).Value() {
			res = 1
		} else {
			res = 0
		}
	case 6:
		first := op.SubPackets[0]
		second := op.SubPackets[1]
		if (*first).Value() < (*second).Value() {
			res = 1
		} else {
			res = 0
		}
	case 7:
		first := op.SubPackets[0]
		second := op.SubPackets[1]
		if (*first).Value() == (*second).Value() {
			res = 1
		} else {
			res = 0
		}
	}

	return res
}

func main() {
	now := time.Now()

	solve()

	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve() {
	binary := parseInput()
	p, _ := parsePacket(binary)
	fmt.Println((*p).Value())
}

func parsePacket(bin string) (*Packet, string) {
	v := binToInt(bin[0:3])
	t := binToInt(bin[3:6])

	var p Packet
	var remainder string
	if t == VALUE_PACKET_TYPE {
		var packet *ValuePacket
		packet, remainder = parseValuePacket(bin[6:], v)

		p = packet
	} else {
		var packet *OperatorPacket
		packet, remainder = parseOperatorPacket(bin[6:], v, t)

		p = packet
	}

	return &p, remainder
}

func parseOperatorPacket(bin string, version int, t int) (*OperatorPacket, string) {
	var remainder string

	p := OperatorPacket{v: version, t: t}
	l := bin[0]

	if l == '0' {
		bitsLength := binToInt(bin[1:16])

		packagesPart := bin[16 : 16+bitsLength]
		for len(packagesPart) > 0 {
			var subP *Packet
			var rem string
			subP, rem = parsePacket(packagesPart)

			p.SubPackets = append(p.SubPackets, subP)
			packagesPart = rem
		}
		remainder = bin[16+bitsLength:]
	} else {
		packetsCount := binToInt(bin[1:12])

		packagesPart := bin[12:]
		for i := 0; i < packetsCount; i++ {
			var subP *Packet
			var rem string
			subP, rem = parsePacket(packagesPart)

			p.SubPackets = append(p.SubPackets, subP)
			packagesPart = rem
		}
		remainder = packagesPart
	}

	return &p, remainder
}

func parseValuePacket(bin string, version int) (*ValuePacket, string) {
	var remainder string
	p := ValuePacket{v: version, t: VALUE_PACKET_TYPE}

	var binVal string
	for len(bin) > 4 {
		flag := bin[0]
		binVal += bin[1:5]

		if flag == '0' {
			remainder = bin[5:]
			break
		} else {
			bin = bin[5:]
		}
	}
	p.val = binToInt(binVal)

	return &p, remainder
}

func binToInt(bin string) int {
	res, _ := strconv.ParseInt(bin, 2, 64)

	return int(res)
}

func parseInput() string {
	input, _ := ioutil.ReadFile("input.txt")
	mapping := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	var res string
	for _, v := range strings.Split(string(input), "") {
		res += mapping[v]
	}

	return res
}
