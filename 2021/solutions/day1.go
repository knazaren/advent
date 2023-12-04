package solutions

func countIncreasing(numbers []int64) int64 {
	var count int64 = 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			count++
		}
	}
	return count
}
