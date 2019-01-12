package fizzbuzz_test

import "testing"
import "github.com/vichar/fizzbuzz"

func TestFizzBuzz(t *testing.T) {
	t.Run("1 should say 1", func(t *testing.T) {
		get := fizzbuzz.Say(1)
		want := "1"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})

	t.Run("2 should say 2", func(t *testing.T) {
		get := fizzbuzz.Say(2)
		want := "2"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})

	t.Run("3 should say Fizz", func(t *testing.T) {
		get := fizzbuzz.Say(3)
		want := "Fizz"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})

	t.Run("4 should say 4", func(t *testing.T) {
		get := fizzbuzz.Say(4)
		want := "4"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})

	t.Run("5 should say Buzz", func(t *testing.T) {
		get := fizzbuzz.Say(5)
		want := "Buzz"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})

}
