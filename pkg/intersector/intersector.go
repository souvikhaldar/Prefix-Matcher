package intersector

type Intersector interface {
	// return the list of common strings between the two
	// list of strings in the input
	Intersect([]string, []string) []string
}
