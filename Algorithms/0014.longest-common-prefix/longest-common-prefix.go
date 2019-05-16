func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    str := strs[0]
    for i := 1; i < len(strs); i++ {
        if len(str) < len(strs[i]) {
            str = helper(str, strs[i])
        } else {
            str = helper(strs[i], str)
        }
    }
    return str
}

func helper(s1 string, s2 string) string {
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            return s1[:i]
        }
    }
    
    return s1
}
