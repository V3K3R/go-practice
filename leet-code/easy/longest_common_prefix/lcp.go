package longestcommonprefix
func lcp(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := ""
	for i:=0; i<len(strs[0]);i++ {
		for j:=1; j<len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return prefix
			}
		}
		prefix += string(strs[0][i])
	}
	return prefix

}