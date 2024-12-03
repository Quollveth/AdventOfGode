package day3

import (
	"fmt"
	"strconv"

	"github.com/quollveth/AdventOfGode/util"
)

type stack []rune

func (s *stack) push(v rune) {
	*s = append(*s, v)
}

func (s *stack) pop() rune {
	if len(*s) == 0 {
		return 0
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *stack) peek() rune {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

func (s *stack) clear() {
	*s = []rune{}
}

func isNum(c rune) bool {
	return c >= '0' && c <= '9'
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func performMul(s *stack) int {
	// assume the stack is already correct

	num1 := -1
	num2 := -1

	n := ""
	for range len(*s) {
		cur := s.pop()
		if cur == ')' { // ignore
			continue
		}
		// comma marks the end of num2
		if cur == ',' {
			num2, _ = strconv.Atoi(reverseString(n))
			n = ""
			continue
		}
		// open parenthesis marks the end of num1
		if cur == '(' {
			num1, _ = strconv.Atoi(reverseString(n))
			break
		}

		n += string(cur)
	}

	return num1 * num2
}

const (
	doReady   int = 1
	dontReady int = 2
)

func Run() {
	input := util.ReadFileFull("day3/input")

	sum := 0
	stack := stack{}
	enabled := true
	ready := 0
	for _, c := range input {
		top := stack.peek()

		switch rune(c) {
		case 'd':
			stack.clear()
			stack.push('d')
		case 'o':
			if top == 'd' {
				stack.push('o')
				continue
			}
			stack.clear()
		case 'n':
			if top == 'o' {
				stack.push('n')
				continue
			}
			stack.clear()
		case '\'':
			if top == 'n' {
				stack.push('\'')
				continue
			}
			stack.clear()
		case 't':
			if top == '\'' {
				stack.push('t')
				continue
			}
			stack.clear()
		case 'm':
			if !enabled {
				stack.clear()
				continue
			}

			stack.clear()
			stack.push('m')
		case 'u':
			if !enabled {
				stack.clear()
				continue
			}
			if top == 'm' {
				stack.push('u')
				break
			}
			stack.clear() // empty stack on invalid token
		case 'l':
			if !enabled {
				stack.clear()
				continue
			}
			if top == 'u' {
				stack.push('l')
				break
			}
			stack.clear()
		case ',':
			if !enabled {
				stack.clear()
				continue
			}
			if isNum(top) {
				stack.push(',')
				break
			}
			stack.clear()
		case '(':
			if top == 'o' {
				ready = doReady
				stack.push('(')
				continue
			}
			if top == 't' {
				ready = dontReady
				stack.push(')')
				continue
			}

			if !enabled {
				stack.clear()
				continue
			}
			if top == 'l' {
				stack.push('(')
				break
			}
			stack.clear()
		case ')':
			if ready == doReady {
				enabled = true
				ready = 0
				stack.clear()
				continue
			}
			if ready == dontReady {
				enabled = false
				ready = 0
				stack.clear()
				continue
			}

			if !enabled {
				stack.clear()
				continue
			}
			if isNum(top) {
				stack.push(')')
				// we reached the end of a mul() operation, calculate
				sum += performMul(&stack) // this function will clear the stack
			}
			stack.clear()
		default:
			if !enabled {
				stack.clear()
				continue
			}
			if isNum(rune(c)) {
				// a number has to follow another number, a comma, or an open parenthesis
				val := top
				if val == '(' || val == ',' || isNum(val) {
					stack.push(rune(c))
					continue
				}
				stack.clear()
			}
			stack.clear() // clear on invalid token
		}
	}

	fmt.Println(sum)
}
