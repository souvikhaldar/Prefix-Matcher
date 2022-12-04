package intersector

import "testing"

func TestIntersector(t *testing.T) {
	i := NewHashMapIntersector()
	c := i.Intersect([]string{"souvik", "gouri", "moti"}, []string{"souvik", "gouri"})
	if len(c) != 2 {
		t.Fatal("Wrong number of common")
	}
}
