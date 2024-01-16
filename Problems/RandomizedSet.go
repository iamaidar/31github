package problems

import "math/rand"

type RandomizedSet map[int]struct{}

func Constructor() RandomizedSet {
	return RandomizedSet{}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := (*this)[val]; !ok {
		(*this)[val] = struct{}{}
		return true
	}

	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := (*this)[val]; ok {
		delete((*this), val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(*this))
	idx := 0
	for k, _ := range *this {
		if idx == i {
			return k
		}
		idx++
	}
	return i
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
