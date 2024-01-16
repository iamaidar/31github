package problems

import "math/rand"

//BEST SOLUTION (HALF MY)
type RandomizedSet struct {
	m map[int]struct{}
	n []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{map[int]struct{}{}, []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.m[val] = struct{}{}
		this.n = append(this.n, val)
		return true
	}

	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	for i, v := range this.n {
		if v == val {
			delete(this.m, val)
			this.n = append(this.n[:i], this.n[i+1:]...)
			return true
		}
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.n[rand.Intn(len(this.n))]
}

//MY SOLUTION (NON OPTIMIZED)
// type RandomizedSet map[int]struct{}

// func Constructor() RandomizedSet {
// 	return RandomizedSet{}
// }

// func (this *RandomizedSet) Insert(val int) bool {
// 	if _, ok := (*this)[val]; !ok {
// 		(*this)[val] = struct{}{}
// 		return true
// 	}

// 	return false
// }

// func (this *RandomizedSet) Remove(val int) bool {
// 	if _, ok := (*this)[val]; ok {
// 		delete((*this), val)
// 		return true
// 	}

// 	return false
// }

// func (this *RandomizedSet) GetRandom() int {
// 	i := rand.Intn(len(*this))
// 	idx := 0
// 	for k, _ := range *this {
// 		if idx == i {
// 			return k
// 		}
// 		idx++
// 	}
// 	return i
// }

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
