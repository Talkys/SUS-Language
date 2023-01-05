package main

// This should not be used outside the instructions
type memory struct {
	tape [256]int16
}

//Those two are constants

/*Instruction map*/
/*It can only holds 255 instructions (0 is flag)*/
/*
The original set is just the 4 basic
operations with their immediate variations.

Extended instructions should be separated with
a blank line for better reading.
*/
var instruction_map = map[string]uint8{
	"addi":  1,
	"subi":  2,
	"multi": 3,
	"divi":  4,
	"add":   5,
	"sub":   6,
	"mult":  7,
	"div":   8,

	//Implement those on a different file
	"print": 9,
}

/*Instruction hash*/
var instruction_hash = []func(
	instruction instruction_struct,
	mem *memory,
){
	inst_nop,
	inst_addi,
	inst_subi,
	inst_multi,
	inst_divi,
	inst_add,
	inst_sub,
	inst_mult,
	inst_div,

	inst_print,
}

func inst_nop(
	instruction instruction_struct,
	mem *memory,
) {
	return //This is just a placeholder
}

func inst_addi(
	instruction instruction_struct,
	mem *memory) {
	mem.tape[instruction.destination] = instruction.first_operand + instruction.second_operand
}

func inst_subi(
	instruction instruction_struct,
	mem *memory) {
	mem.tape[instruction.destination] = instruction.first_operand - instruction.second_operand
}

func inst_multi(
	instruction instruction_struct,
	mem *memory) {
	mem.tape[instruction.destination] = instruction.first_operand * instruction.second_operand
}

func inst_divi(
	instruction instruction_struct,
	mem *memory) {
	mem.tape[instruction.destination] = instruction.first_operand / instruction.second_operand
}

func inst_add(
	instruction instruction_struct,
	mem *memory) {
	mem.tape[instruction.destination] = mem.tape[instruction.first_operand] + mem.tape[instruction.second_operand]
}

func inst_sub(
	instruction instruction_struct,
	mem *memory) {
	mem.tape[instruction.destination] = mem.tape[instruction.first_operand] - mem.tape[instruction.second_operand]
}

func inst_mult(
	instruction instruction_struct,
	mem *memory) {
	mem.tape[instruction.destination] = mem.tape[instruction.first_operand] * mem.tape[instruction.second_operand]
}

func inst_div(
	instruction instruction_struct,
	mem *memory) {
	mem.tape[instruction.destination] = mem.tape[instruction.first_operand] / mem.tape[instruction.second_operand]
}
