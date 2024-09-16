package main

import (
	"fmt"

	"github.com/bitfield/script"
)

func main() {
	fmt.Println("=========READ FILE=============")
	contents, err := script.File("test.txt").String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contents)
	fmt.Println("==========COUNT LINE============")
	numLines, err := script.File("test.txt").CountLines()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(numLines)
	fmt.Println("==========COUNT WORDS============")
	numhello, err := script.File("test.txt").Match("hello").CountLines()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(numhello)
	//fmt.Println("=========Take INPUNE, PRINT IF MATCH hello=============")
	//script.Stdin().Match("hello").Stdout()
	fmt.Println("=========SHA256Sum=============")
	sum, err := script.Echo("hello").SHA256Sum()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
	fmt.Println("=========SHA256Sums=============")
	script.Echo("test.txt").SHA256Sums().Stdout()
	//result, err := script.Echo("hello\nworld").SHA256Sums().String()
	/*
		result, err := script.Echo("test.txt").SHA256Sums().String()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	*/
}
