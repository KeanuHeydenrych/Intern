package main

func Add(x int, y int) int {

	return x + y
}

func Multi(x int, y int) int {

	return x * y
}

func Swap(x, y string) (string, string) {

	return y, x
}

func Split(sum int) (x, y int) {

	x = sum * 4 / 9
	y = sum - x
	return
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func NeedInt(x int) int {
	return x*10 + 1
}
func NeedFloat(x float64) float64 {
	return x * 0.1
}

func Bark() string {
	return "WOOOOF!"
}
