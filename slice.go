package slice

import "golang.org/x/exp/slices"

func Reverse[S ~[]E, E any](s S) (res S) {
	res = make(S, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return
}

func Windows[S ~[]E, E any](s []E, size int) (res []S) {
	if size <= 0 {
		panic("size must be > 0")
	}
	if len(s) < size {
		return
	}
	res = make([]S, 0, len(s)-(size-1))
	for len(s) >= size {
		res = append(res, s[:size])
		s = s[1:]
	}
	return
}

func RWindows[S ~[]E, E any](s S, size int) (res []S) {
	if size <= 0 {
		panic("size must be > 0")
	}
	if len(s) < size {
		return
	}
	l := len(s) - size
	res = make([]S, 0, l)
	for i := l; i >= 0; i-- {
		res = append(res, s[i:i+size])
	}
	return
}

func Chunks[S ~[]E, E any](s S, size int) (res []S) {
	if size <= 0 {
		panic("size must be > 0")
	}
	res = make([]S, 0, (len(s)+size-1)/size)
	for len(s) > 0 {
		if size > len(s) {
			size = len(s)
		}
		res = append(res, s[:size])
		s = s[size:]
	}
	return
}

func RChunks[S ~[]E, E any](s S, size int) (res []S) {
	if size <= 0 {
		panic("size must be > 0")
	}
	res = make([]S, 0, (len(s)+size-1)/size)
	for i := len(s) - size; i > -size; i -= size {
		if i < 0 {
			size += i
			i = 0
		}
		res = append(res, s[i:i+size])
	}
	return
}

func Split[S ~[]E, E comparable](s S, sep E) (res []S) {
	for len(s) > 0 {
		idx := slices.Index([]E(s), sep)
		if idx < 0 {
			res = append(res, s)
			return
		}
		res = append(res, s[:idx])
		s = s[idx+1:]
	}
	return
}

func SplitFunc[S ~[]E, E any](s S, eq func(e E) bool) (res []S) {
	for len(s) > 0 {
		idx := slices.IndexFunc([]E(s), eq)
		if idx < 0 {
			res = append(res, s)
			return
		}
		res = append(res, s[:idx])
		s = s[idx+1:]
	}
	return
}

func SplitInclusive[S ~[]E, E comparable](s S, sep E) (res []S) {
	for len(s) > 0 {
		idx := slices.Index([]E(s), sep)
		if idx < 0 {
			res = append(res, s)
			return
		}
		res = append(res, s[:idx+1])
		s = s[idx+1:]
	}
	return
}

func SplitInclusiveFunc[S ~[]E, E any](s S, eq func(e E) bool) (res []S) {
	for len(s) > 0 {
		idx := slices.IndexFunc([]E(s), eq)
		if idx < 0 {
			res = append(res, s)
			return
		}
		res = append(res, s[:idx+1])
		s = s[idx+1:]
	}
	return
}

func FlatEq[S ~[]E, E comparable](s1, s2 []S) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if !slices.Equal[E](v1, v2) {
			return false
		}
	}
	return true
}
