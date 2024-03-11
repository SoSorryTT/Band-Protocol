package main

func Problem1() {
	// Testcase data
	input1 := "SRSSRRR"
	input2 := "RSSRR"
	input3 := "SSSRRRRS"
	input4 := "SSRR"
	input5 := "SRRSSR"
	input6 := "SSRSRRRRR"

	test1 := CheckGoodBoy(input1)
	println("test1 : ", test1)
	test2 := CheckGoodBoy(input2)
	println("test2 : ", test2)
	test3 := CheckGoodBoy(input3)
	println("test3 : ", test3)
	test4 := CheckGoodBoy(input4)
	println("test4 : ", test4)
	test5 := CheckGoodBoy(input5)
	println("test5 : ", test5)
	test6 := CheckGoodBoy(input6)
	println("test6 : ", test6)
}

// CheckGoodBoy run time is O(n) because it do only 1 loop and the other are O(1)
func CheckGoodBoy(inputString string) string {
	// If big boss shoot first and cannot revenge return Bad boy
	if inputString[0] == 'R' || inputString[len(inputString)-1] == 'S'  {
		return "Bad boy"
	}

	countS := 0
	for _, c := range inputString {
		if c == 'S' {
			countS += 1
		} else if countS > 0 {
			countS -= 1
		}
	}
	if countS > 0 {
		return "Bad boy"
	}
	return "Good boy"
}
