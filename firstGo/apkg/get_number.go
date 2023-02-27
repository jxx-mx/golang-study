package apkg

var number = 0

func GetNumber() int {
	number++
	return number
}