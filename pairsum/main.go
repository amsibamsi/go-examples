package pairsum

func HasPairWithSum(a []int, sum int) bool {
	seen := make(map[int]bool)
	found := false
	for _, v := range a {
		comp := sum - v
		if seen[comp] {
			found = true
			break
		} else {
			seen[v] = true
		}
	}
	return found
}
