package utils

func Swap(str1, str2 *string) {
	*str1, *str2 = *str2, *str1
}
