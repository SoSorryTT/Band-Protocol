package main

func Problem1() {
	//testcase data
	input1 := "SRSSRRR"
	input2 := "RSSRR"
	input3 := "SSSRRRRS"
	input4 := "SSRR"
	input5 := "SRRSSR"

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
}

// CheckGoodBoy run time is O(n) because it do only 1 loop and the other are O(1)
func CheckGoodBoy(inputString string) string {
	//if big boss shoot first return Bad boy
	if inputString[0] == 'R' {
		return "Bad boy"
	}
	countShoot := 0
	countRevenge := 0
	checkRevenge := false
	for _, c := range inputString {
		if c == 'S' {
			//if big boss revenge check condition
			if checkRevenge && countRevenge-countShoot < 0 {
				return "Bad boy" //return Bad boy if revenge fail
			} else if checkRevenge {
				//if revenge not fail reset value
				checkRevenge = false
				countShoot = 0
				countRevenge = 0
			}
			countShoot += 1
		} else if c == 'R' {
			countRevenge += 1
			checkRevenge = true //check if big boss revenge?
		}
	}
	//if revenge success return Good boy
	if countRevenge >= countShoot && inputString[len(inputString)-1] != 'S' {
		return "Good boy"
	}

	return "Bad boy"
}
