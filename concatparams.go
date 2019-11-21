package piscine

//ConcatParams is a function
func ConcatParams(args []string) string {
	build := ""
	for index := range args {
		build += args[index] + "\n"
	}
	return build
}
