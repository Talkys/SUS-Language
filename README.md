# SUS-Language
 Small programming language for basic math

 This project is a simple language that can do the 4 basic operations plus a print operation.

 SUS is not a big project, but can be extended by adding new instructions.

 ## How it works

 SUS is like ARM Assembly, with the instruction and the operands.

 Example:

 > addi, 2, 4, 8, This is a comment. The last comma is mandatory

 the basic instruction set is:

 > addi 
 
 > subi

 > multi

 > divi

 > add
 
 > sub

 > mult

 > div

 > print

 The instruction structure is:

 > Instruction, 1st operand, 2nd operand, destination address

 The first 4 instructions use literals as operands and the next 4 use the memory address.

 The print instruction prints the memory value at the firs operand.