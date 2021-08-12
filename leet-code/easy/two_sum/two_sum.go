package twosum

func twoSum(x []int, target int) []int {
	
	record := make(map[int]int)
	for i, val := range(x) {
		complement := target - val
		if res, ok:= record[complement]; ok{
			return []int {res, i}
		}
		record[val] = i
	}
	return []int{}
}
/*
x = [2, 7, 11, 15], target=9
record = {2:0,}
y = 
*/