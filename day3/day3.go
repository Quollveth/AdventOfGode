package day3

import (
	"fmt"

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

func reverseInt(n int) int {
	new := 0
	for n > 0 {
		r := n % 10
		new += 10
		new += r
		n /= 10
	}

	return new
}

func performMul(s *stack) int {
	// assume the stack is already correct
	fmt.Println("Received stack:", string(*s))

	num1 := -1
	num2 := -1

	n := 0
	for range len(*s) {
		cur := s.pop()
		if cur == ')' { // ignore
			continue
		}
		// comma marks the end of num2
		if cur == ',' {
			num2 = reverseInt(n)
			n = 0
			continue
		}
		// open parenthesis marks the end of num1
		if cur == '(' {
			num1 = reverseInt(n)
			break
		}

		n *= 10
		n += int(cur - '0')
	}

	fmt.Printf("Multiplying %v and %v\n\n", num1, num2)
	return num1 * num2
}

func Run() {
	input := util.ReadFileFull("day3/testin")

	sum := 0
	stack := stack{}
	for _, c := range input {
		switch rune(c) {
		case 'm':
			stack.clear()
			stack.push('m')
		case 'u':
			if stack.peek() == 'm' {
				stack.push('u')
				break
			}
			stack.clear() // empty stack on invalid token
		case 'l':
			if stack.peek() == 'u' {
				stack.push('l')
				break
			}
			stack.clear()
		case ',':
			if isNum(stack.peek()) {
				stack.push(',')
				break
			}
			stack.clear()
		case '(':
			if stack.peek() == 'l' {
				stack.push('(')
				break
			}
			stack.clear()
		case ')':
			if isNum(stack.peek()) {
				stack.push(')')
				// we reached the end of a mul() operation, calculate
				sum += performMul(&stack) // this function will clear the stack
			}
			stack.clear()
		default:
			if isNum(rune(c)) {
				// a number has to follow another number, a comma, or an open parenthesis
				val := stack.peek()
				if val == '(' || val == ',' || isNum(val) {
					stack.push(rune(c))
					continue
				}
				stack.clear() // clear on invalid token
			}
		}
	}

	fmt.Println(sum)
}
