package problems

func nextGreatestLetter(letters []byte, target byte) byte {
	l, h := 0, len(letters)
	for l < h {
		mid := l + (h-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			h = mid
		}
	}

	if l == len(letters) {
		return letters[0]
	}

	return letters[l]
}
