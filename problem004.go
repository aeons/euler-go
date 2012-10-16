package main

func init() {
	solvers[4] = Solve4
}

func Solve4() (solution int) {
	var isPalindromeDigit func(n, d int) bool

	isPalindrome := func(n int) bool {
		// fmt.Println(n)
		switch {
		// 1 digits
		case n < 10:
			return true
			// 2 digits
		case n < 100:
			if n/10 == n%10 {
				return true
			}
			// 3 digits
		case n < 1000:
			if n/100 == n%10 {
				return true
			}
			// 4 digits
		case n < 10000:
			if f, l := n/1000, n%10; f == l {
				return isPalindromeDigit((n-1000*f)/10, 2)
			}
			// 5 digits
		case n < 100000:
			if f, l := n/10000, n%10; f == l {
				return isPalindromeDigit((n-10000*f)/10, 3)
			}
			// 6 digits
		default:
			if f, l := n/100000, n%10; f == l {
				return isPalindromeDigit((n-100000*f)/10, 4)
			}
		}
		return false
	}

	isPalindromeDigit = func(n, d int) bool {
		switch d {
		case 2:
			if n/10 == n%10 {
				return true
			}
		case 3:
			if n/100 == n%10 {
				return true
			}
		case 4:
			if f, l := n/1000, n%10; f == l {
				return isPalindromeDigit((n-1000*f)/10, 2)
			}
		}
		return false
	}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if k := i * j; isPalindrome(k) && k > solution {
				solution = k
			}
		}
	}
	return
}
