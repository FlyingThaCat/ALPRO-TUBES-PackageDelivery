package utils

func SortArray(arr []int, mode string) {
	if mode == "asc" {
		for a := 1; a < len(arr); a++ {
			ptr := arr[a]
			b := a - 1

			for b >= 0 && arr[b] > ptr {
				arr[b+1] = arr[b]
				b--
			}
			arr[b+1] = ptr
		}
	} else if mode == "desc" {
		for a := 1; a < len(arr); a++ {
			ptr := arr[a]
			b := a - 1

			for b >= 0 && arr[b] < ptr {
				arr[b+1] = arr[b]
				b--
			}
			arr[b+1] = ptr
		}
	}
}
