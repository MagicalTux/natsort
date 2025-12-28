// Package natsort implements natural strings sorting
package natsort

import (
	"sort"
	"strings"
)

type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Less(a, b int) bool {
	return Compare(s[a], s[b])
}

func (s stringSlice) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

// Sort sorts a list of strings in a natural order
func Sort(l []string) {
	sort.Sort(stringSlice(l))
}

// Compare returns true if the first string precedes the second one according to natural order
func Compare(a, b string) bool {
	lenA := len(a)
	lenB := len(b)
	posA := 0
	posB := 0

	for {
		if lenA <= posA {
			if lenB <= posB {
				// eof on both at the same time (equal)
				return false
			}
			return true
		} else if lenB <= posB {
			// eof on b
			return false
		}

		av, bv := a[posA], b[posB]

		if av >= '0' && av <= '9' && bv >= '0' && bv <= '9' {
			// go into numeric mode
			intLenA := 1
			intLenB := 1
			for {
				if posA+intLenA >= lenA {
					break
				}
				x := a[posA+intLenA]
				if av == '0' && x >= '0' && x <= '9' {
					posA += 1
					av = x
					continue
				}
				if x >= '0' && x <= '9' {
					intLenA += 1
				} else {
					break
				}
			}
			for {
				if posB+intLenB >= lenB {
					break
				}
				x := b[posB+intLenB]
				if bv == '0' && x >= '0' && x <= '9' {
					posB += 1
					bv = x
					continue
				}
				if x >= '0' && x <= '9' {
					intLenB += 1
				} else {
					break
				}
			}
			if intLenB > intLenA {
				// length of b is longer, means b is a bigger number
				return true
			} else if intLenA > intLenB {
				return false
			}
			// both have same length, let's compare as string
			v := strings.Compare(a[posA:posA+intLenA], b[posB:posB+intLenB])
			if v < 0 {
				return true
			} else if v > 0 {
				return false
			}
			// equal
			posA += intLenA
			posB += intLenB
			continue
		}

		if av == bv {
			posA += 1
			posB += 1
			continue
		}

		return av < bv
	}
}
