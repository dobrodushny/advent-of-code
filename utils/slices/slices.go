package slices

import (
	"strconv"
	"strings"
)

func ToIntSlice(str_slice []string) []int {
	var int_slice []int
	for _, str := range str_slice {
		for _, v := range strings.Fields(str) {
			int_v, _ := strconv.Atoi(v)
			int_slice = append(int_slice, int_v)
		}
	}

	return int_slice
}

func Intersect[T comparable](s1, s2 []T) []T {
	var result []T
	set := make(map[T]bool)

	for _, el := range s1 {
		if !set[el] {
			set[el] = true
		}
	}

	for _, el := range s2 {
		if set[el] {
			result = append(result, el)
		}
	}

	return result
}

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
