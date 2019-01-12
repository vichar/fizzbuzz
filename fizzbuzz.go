package fizzbuzz

import "strconv"

func main() {

}

func Say(i int) string {
	if i%3 == 0 {
		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(i)
}
