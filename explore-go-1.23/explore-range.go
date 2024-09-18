package newfeatures

import "fmt"

func Reverse(K []int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := range K {
			if !yield(i) {
				return
			}
		}
	}
}

func NewRange() {
	// the new Go 1.22 supports the range operator on functions the return signaute
	// func(func(int) bool)
	nums := []int{3, 2, 1}
	// for indx := range Reverse(nums){
	//	  fmt.Println(indx)
	//	}

	Reverse(nums)(func(indx int) bool {
		//ops inside the loop
		fmt.Println(indx)
		return true
	})
}
