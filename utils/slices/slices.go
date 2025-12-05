package slices

import (
	"advent-of-code/utils"
	"strings"
)

func StringToIntSlice(str string) []int {
	var int_slice []int

	for _, v := range strings.Fields(str) {
		int_slice = append(int_slice, utils.Atoi(v))
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

func RemoveDuplicate[T string | int](sliceList []T) []T {
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

// Deprecated: do not use as the signature is confusing
func ToIntSlice(str_slice []string) []int {
	var int_slice []int

	for _, str := range str_slice {
		for _, v := range strings.Fields(str) {
			int_slice = append(int_slice, utils.Atoi(v))
		}
	}

	return int_slice
}

func RemoveAt[T comparable](slice []T, i int) []T {
	return append(slice[:i], slice[i+1:]...)
}
