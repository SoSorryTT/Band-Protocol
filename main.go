package main

import (
	"assignment/band-protocol/problem3"
	"os"
)

func main() {
	command := os.Args[1:][len(os.Args[1:])-1]
	switch command {
	case "problem1":
		Problem1()
	case "problem2":
		Problem2()
	case "problem3":
		problem3.Problem3()
	}
}
