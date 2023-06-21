func Gimme(array [3]int) int { 
	arrayCopy := array

	arrSlice := arrayCopy[:]
	sort.Ints(arrSlice)
	val := arrSlice[1]

	for i, v := range array {
		if v == val {
			return i
		}
	}

	return 0
}
