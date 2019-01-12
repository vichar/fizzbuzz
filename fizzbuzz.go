package fizzbuzz

import "strconv"

func main() {

}

func Say(i int) string {
	if i%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(i)
}
