package main

import "fmt"

func inst_print(
	instruction instruction_struct,
	mem *memory) {
	fmt.Println(mem.tape[instruction.first_operand])
}
