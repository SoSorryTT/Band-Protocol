package main

func main() {
	//test data
	n1 := 5
	k1 := 5
	chickenAt1 := []int{2, 5, 10, 12, 15}

	save1 := CheckCanSave(n1, k1, chickenAt1)
	println(save1)

	n2 := 6
	k2 := 10
	chickenAt2 := []int{1, 11, 30, 34, 35, 37}

	save2 := CheckCanSave(n2, k2, chickenAt2)
	println(save2)
}

func CheckCanSave(n int, k int, chickenAt []int) int {
	maxChickens := 0
	//loop for the first start number
	for i := 0; i < n; i++ {
		canSave := 0
		//loop for check how many can save
		//check if number in range itself - k
		//check equal or more than k-1 because it needs to count itself(i)
		for j := i; j < n && chickenAt[j]-chickenAt[i] <= k-1; {
			canSave += 1
			j += 1
		}
		//check if can same more than the last one
		if canSave > maxChickens {
			maxChickens = canSave
		}
	}

	return maxChickens
}
