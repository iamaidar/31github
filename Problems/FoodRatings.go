package problems

type FoodRatings struct {
	foods    []string
	cuisines []string
	ratings  []int
}

func lex(a, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			return true
		} else if a[i] > b[j] {
			return false
		} else {
			i++
			j++
		}
	}

	return false
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	return FoodRatings{foods, cuisines, ratings}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	idx := -1
	for i := 0; i < len(this.foods); i++ {
		if this.foods[i] == food {
			idx = i
			break
		}
	}

	this.ratings[idx] = newRating
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	max := -999999999
	idx := -1
	for i := 0; i < len(this.ratings); i++ {
		if this.ratings[i] > max && this.cuisines[i] == cuisine {
			max = this.ratings[i]
			idx = i
		} else if this.ratings[i] == max && this.cuisines[i] == cuisine {
			if lex(this.foods[i], this.foods[idx]) {
				max = this.ratings[i]
				idx = i
			}
		}
	}

	return this.foods[idx]
}
