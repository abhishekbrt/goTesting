package arraysandslices



func SumSlice(numbers []int) int {
	ans:=0
	for _ , num:= range numbers {
		ans+=num
	}
	return ans
}