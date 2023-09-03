package main

func compact[S ~[]E, E comparable](s S) func(yield func(val E) bool) bool {
	return func(yield func(val E) bool) bool {
		for i, v := range s {
			if i+1 == len(s) {
				if !yield(v) {
					return false
				}
				return true
			}

			if v == s[i+1] {
				continue
			}

			if !yield(v) {
				return false
			}
		}
		return true
	}
}
