package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
   #######    Very sus
  #       #  /
  #  ######
### #      #
# # #      #
# #  ######
# #       #
# #       #
# #       #
###       #
  #  ###  #
  #  # #  #
  #### ####
*/

func main() {
	fmt.Print("SUSVM V")
	fmt.Println(version)
	var mem memory

	inputfile, err := os.Open(os.Args[1])

	if err != nil {
		panic("Error reading the input file")
	}

	//fmt.Println("Passou")
	filescanner := bufio.NewScanner(inputfile)
	filescanner.Split(bufio.ScanLines)

	for filescanner.Scan() {
		//fmt.Println("Lendo instrucao " + filescanner.Text())
		execute(&mem, filescanner.Text())
	}
}
