//Some additional utilities that are useful when processing CSV headers and data.
package utils

type Index map[string]int

// Return a map that maps each string in the input slice to its index in the slice.
func NewIndex(a []string) Index {
	index := make(map[string]int)
	for i, v := range a {
		index[v] = i
	}
	return index
}

// Answer true if the index contains the specified string.
func (i Index) Contains(k string) bool {
	_, ok := i[k]
	return ok
}

// Calculate the intersection between two string slices. The first returned slice
// is the intersection between the two slices. The second returned slice is
// a slice of elements in the first slice but not the second. The third returned
// slice is a slice of elements in the second slice but not the first.
func Intersect(a []string, b []string) ([]string, []string, []string) {
	index := NewIndex(a)
	result := make([]string, 0, len(b))
	aNotB := make([]string, len(a), len(a))
	copy(aNotB, a)
	bNotA := make([]string, 0, len(b))
	for _, v := range b {
		if i, ok := index[v]; ok {
			result = append(result, v)
			aNotB[i] = ""
		} else {
			bNotA = append(bNotA, v)
		}
	}
	i := 0
	for j := range a {
		present := (aNotB[j] == a[j])
		aNotB[i] = a[j]
		if present {
			i++
		}
	}
	aNotB = aNotB[0:i]
	return result, aNotB, bNotA
}
