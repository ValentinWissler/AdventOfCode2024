package computer

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Computer struct {
	a, b, c int
	cmds    []string
	output  []string
}

func NewComputer(a, b, c int, cmds ...string) *Computer {
	return &Computer{a: a, b: b, c: c, cmds: cmds}
}

func (c *Computer) ProcessCmds() string {
	for i := 0; i < len(c.cmds); i += 2 {
		instruction := c.cmds[i]
		if i+1 >= len(c.cmds) {
			i = -2
			continue
		}
		next := c.operand(c.cmds[i+1])
		switch instruction {
		case "0":
			c.adv(next)
		case "1":
			n, _ := strconv.Atoi(c.cmds[i+1])
			c.bxl(c.b, n)
		case "2":
			c.bst(next)
		case "3":
			// Pointer goes forward 1
			if c.jnz(c.a) {
				i--
			}
		case "4":
			c.bxc()
		case "5":
			c.out(next)
		case "6":
			c.bdv(next)
		case "7":
			c.cdv(next)
		default:
			panic(fmt.Sprintf("we tried to process an unknown instruction type: %s", instruction))
		}
	}
	return strings.Join(c.output, ",")
}

func (c *Computer) operand(op string) int {
	next, _ := strconv.Atoi(op)
	if next == 4 {
		return c.a
	} else if next == 5 {
		return c.b
	} else if next == 6 {
		return c.c
	} else {
		return next
	}
}

/*
The adv instruction (opcode 0) performs division. The numerator is the value in the A register.
The denominator is found by raising 2 to the power of the instruction's combo operand.
(So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.)
The result of the division operation is truncated to an integer and then written to the A register
*/
func (c *Computer) adv(denominator int) {
	c.a = c.a / int(math.Pow(2, float64(denominator)))
}

/*
The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand, then stores the result in register B.
*/
func (c *Computer) bxl(registerB, next int) {
	c.b = registerB ^ next
}

/*
The bst instruction (opcode 2) calculates the value of its combo operand modulo 8 (thereby keeping only its lowest 3 bits),
then writes that value to the B register.
*/
func (c *Computer) bst(operand int) {
	c.b = operand % 8
}

/*
The jnz instruction (opcode 3) does nothing if the A register is 0. However, if the A register is not zero,
it jumps by setting the instruction pointer to the value of its literal operand; if this instruction jumps,
the instruction pointer is not increased by 2 after this instruction.
*/
func (c *Computer) jnz(registerA int) bool {
	return registerA != 0
}

/*
The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C,
then stores the result in register B. (For legacy reasons, this instruction reads an operand but ignores it.)
*/
func (c *Computer) bxc() {
	c.b = c.b ^ c.c
}

/*
The out instruction (opcode 5) calculates the value of its combo operand modulo 8, then outputs that value.
(If a program outputs multiple values, they are separated by commas.)
*/
func (c *Computer) out(combo int) {
	c.output = append(c.output, fmt.Sprintf("%d", combo%8))
}

/*
The bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored in the B register.
The numerator is still read from the A register.
*/
func (c *Computer) bdv(denominator int) {
	c.b = c.a / int(math.Pow(2, float64(denominator)))
}

/*
The cdv instruction (opcode 7) works exactly like the adv instruction except that the result is stored in the C register.
The numerator is still read from the A register.
*/
func (c *Computer) cdv(denominator int) {
	c.c = c.a / int(math.Pow(2, float64(denominator)))
}
