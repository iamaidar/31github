package problems

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return       -1 if num is higher than the picked number
 *         1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l := 1
	r := 9223372036854775807

	for l <= r {
		m := l + (r-l)/2
		t := guess(m)
		if t == 0 {
			return m
		} else if t == -1 {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return 0

}
