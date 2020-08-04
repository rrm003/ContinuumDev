package test

//Divide ...
func Divide(x, y int) (float64, string) {
	if y == 0 {
		return float64(0), "Divide by zero"
	}
	return float64(x / y), ""
}
