package piscine

//BasicJoin is a function
func BasicJoin(strs []string) string {
	build := ""
	for x := range strs {
		build += strs[x]
	}
	return build
}
