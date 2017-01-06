package arrayperm

func ArrayPermutations(A [][]int) [][]int {
	perm_num := 0
	perm_len := 0
	for _, arr := range A {
		if len(arr) > 0 {
			if perm_num == 0 {
				perm_num = 1
			}
			perm_num *= len(arr)
			perm_len += 1
		}
	}
	perms := make([][]int, perm_num)
	for p := range perms {
		perms[p] = make([]int, perm_len)
	}
	rot := perm_num
	digit := 0
	for _, arr := range A {
		if len(arr) > 0 {
			rot = rot / len(arr)
			for p := 0; p < perm_num; p++ {
				val_ind := (p / rot) % len(arr)
				perms[p][digit] = arr[val_ind]
			}
			digit++
		}
	}
	return perms
}
