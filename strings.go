package xstringutils

func FullName(f, l string) (string, int) {
	full := f + " " + l + " from github package xstringutils UPDATED."
	length := len(full)

	return full, length
}
