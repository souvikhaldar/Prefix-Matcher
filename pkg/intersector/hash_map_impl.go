package intersector

type HashMapIntersector struct{}

func NewHashMapIntersector() *HashMapIntersector {
	return &HashMapIntersector{}
}

// return the list of common strings between the two
// list of strings in the input
func (h *HashMapIntersector) Intersect(first []string, second []string) []string {
	inFirst := make(map[string]struct{})
	for _, el := range first {
		inFirst[el] = struct{}{}
	}
	var common []string
	for _, el := range second {
		if _, ok := inFirst[el]; ok {
			common = append(common, el)
		}
	}
	return common
}
