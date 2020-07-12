package p0008

import (
	"math"
)

var number = map[uint8]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
var symbol = map[uint8]int{'-': -1, '+': 1}

type State int

const (
	StateStart State = iota
	StateSigned
	StateInNumber
	StateEnd
)

type RuneType int

const (
	TypeEmpty RuneType = iota
	TypeNumber
	TypeSymbol
	TypeOther
)

func myAtoi(str string) int {
	table := map[State]map[RuneType]State{
		StateStart: {
			TypeEmpty: StateStart, TypeNumber: StateInNumber, TypeSymbol: StateSigned, TypeOther: StateEnd,
		},
		StateSigned: {
			TypeEmpty: StateEnd, TypeNumber: StateInNumber, TypeSymbol: StateEnd, TypeOther: StateEnd,
		},
		StateInNumber: {
			TypeEmpty: StateEnd, TypeNumber: StateInNumber, TypeSymbol: StateEnd, TypeOther: StateEnd,
		},
		StateEnd: {
			TypeEmpty: StateEnd, TypeNumber: StateEnd, TypeSymbol: StateEnd, TypeOther: StateEnd,
		},
	}
	result := 0
	state := StateStart
	sign := 1
	numberCount := 0
	for i := 0; i < len(str); i++ {
		c := str[i]
		state = table[state][getType(c)]
		if state == StateEnd {
			break
		}
		if numberCount > 10 {
			break
		}
		switch state {
		case StateSigned:
			if c == '-' {
				sign = -1
			}
		case StateInNumber:
			if number[c] != 0 || result > 0 {
				result = result*10 + number[c]
				numberCount++
			}

		}
	}
	result = result * sign
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}
	return result
}

func getType(c uint8) RuneType {
	if c == ' ' {
		return TypeEmpty
	}
	if _, ok := symbol[c]; ok {
		return TypeSymbol
	}
	if _, ok := number[c]; ok {
		return TypeNumber
	}
	return TypeOther
}
