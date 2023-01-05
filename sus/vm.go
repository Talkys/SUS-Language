package main

import (
	"strconv"
	"strings"
)

type instruction_struct struct {
	operation      uint8
	first_operand  int16
	second_operand int16
	destination    uint8
}

func decode(instruction string) uint8 {
	//Must do a can't fetch here, but I'll do that later
	_, ok := instruction_map[instruction]
	if ok {
		return instruction_map[instruction]
	}
	panic("Can't find instruction: " + instruction + "\n")
}

func parse(instruction string) instruction_struct {
	//Removes spaces and splits the instruction by commas
	instruction = strings.ReplaceAll(instruction, " ", "")
	instruction_split := strings.Split(instruction, ",")

	var parsed_instruction instruction_struct
	//fmt.Println(instruction_split)
	switch len(instruction_split) {
	case 5:
		//Fetch the instruction on the map ignoring capitalization
		parsed_instruction.operation = decode(strings.ToLower(instruction_split[0]))

		//Parse the operands to their correct places, just repeat for all of them
		first_operand, err := strconv.Atoi(instruction_split[1])
		if err == nil {
			parsed_instruction.first_operand = int16(first_operand)
		} else {
			panic("Error parsing operand \"" + instruction_split[1] + "\" at instruction " + instruction)
		}
		second_operand, err := strconv.Atoi(instruction_split[2])
		if err == nil {
			parsed_instruction.second_operand = int16(second_operand)
		} else {
			panic("Error parsing operand \"" + instruction_split[2] + "\" at instruction " + instruction)
		}
		destination, err := strconv.Atoi(instruction_split[3])
		if err == nil {
			parsed_instruction.destination = uint8(destination)
		} else {
			panic("Error parsing operand \"" + instruction_split[3] + "\" at instruction " + instruction)
		}

	default:
		panic("Error parsing instruction " + instruction + "\nToo few arguments!")
		//return instruction_struct{operation: 0} // error return value
	}

	//This must be set to be returned, exceptions must return that value above
	return parsed_instruction
}

func execute(mem *memory, instruction string) {
	parsed_instruction := parse(instruction)

	//fmt.Printf("Jooj%d\n", parsed_instruction.operation)

	//Call the function by the hash index
	instruction_hash[parsed_instruction.operation](
		parsed_instruction,
		mem,
	)

	//fmt.Println(mem.tape[3])

}
