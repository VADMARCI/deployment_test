package utils

func IntSliceContains(items []int, key int) bool {
	for _, item := range items {
		if item == key {
			return true
		}
	}

	return false
}

func StringSliceContains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}

	return false
}

func Deduplicate(items []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0, len(items))

	for _, item := range items {
		if _, ok := keys[item]; !ok {
			keys[item] = true
			list = append(list, item)
		}
	}

	return list
}

func DeduplicateInts(items []int) []int {
	keys := make(map[int]bool)
	list := make([]int, 0, len(items))

	for _, item := range items {
		if _, ok := keys[item]; !ok {
			keys[item] = true
			list = append(list, item)
		}
	}

	return list
}

func FilterOutDefaultInts(items []int) []int {
	list := make([]int, 0, len(items))
	for _, v := range items {
		if v != 0 {
			list = append(list, v)
		}
	}
	return list
}
