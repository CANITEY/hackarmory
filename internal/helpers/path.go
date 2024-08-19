func CheckPath(p string) bool {
	pathVar := os.Getenv("PATH")
	pathSlice := strings.Split(pathVar, ":")
	return slices.Contains(pathSlice, p)
}
