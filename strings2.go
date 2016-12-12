package xstringutils

func FullName(f, l string) (string, int) {
	full := f + " " + l + " from github package xstringutils UPDATED 2."
	length := len(full)

	return full, length
}
