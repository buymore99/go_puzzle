package main


func longestConsecutive(nums []int) int {
	numsMap := make(map[int]int)
	lenNums := len(nums)

	for i := 0; i< lenNums; i++ {
		numsMap[nums[i]] = 0
	}

	maxLCS := 0

	for i := 0; i < lenNums ; i++ {
		currentLen := 1
		counter := 1

		for {
			val, ok := numsMap[nums[i] + counter]
			if ok {
				if val != 0 {
					currentLen += val
					break
				} else {
					currentLen += 1
					counter += 1
				}
			} else {
				break
			}
		}

		if currentLen > maxLCS {
			maxLCS = currentLen
		}
		numsMap[nums[i]] = currentLen
	}

	return maxLCS
}

func main() {
	output := longestConsecutive([]int{4, 7, 3, 8, 2, 1})
	fmt.Println(output)

	output = longestConsecutive([]int{4, 7, 3, 8, 2, 1, 9, 24, 10, 11})
	fmt.Println(output)
}