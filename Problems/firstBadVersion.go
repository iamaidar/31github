package problems

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l := 1
	m := l + (n-l)/2
	for l <= n {
		m = l + (n-l)/2
		if isBadVersion(m) {
			break
		} else {
			l = m + 1
		}
	}

	for isBadVersion(m) {
		m--
	}

	return m + 1
}
