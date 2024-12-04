package main

import (
	"bufio"
	"fmt"
	"os"
) 

func main() {
	interest := "like cricket alot"
	fmt.Println(interest)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate your cricket knowlegde")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your reveiw", input)
	fmt.Printf("Type of input is %T ", input)

}
