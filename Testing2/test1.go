package temp

//go:generate mockgen -source test1.go -destination test1_mock.go -package temp
type shape interface {
	area(int) int
}

func area(x int) int {
	return x * x
}
