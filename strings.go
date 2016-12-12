package xstringutils

func FullName(f, l string) (string, int) {
	full := f + " " + l + " from github package xstringutils."
	length := len(full)

	return full, length
}
