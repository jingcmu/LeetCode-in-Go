func generateParenthesis(n int) []string {
    ret := []string{}
      
    backtrack(&ret, "", 0, 0, n)
    return ret
    
}

func backtrack(ret *[]string, ans string, left int, right int, n int){
    if len(ans) == 2*n {
        *ret = append(*ret, ans)
        return 
    }
    
    if left < n {
        backtrack(ret, ans+"(", left+1, right, n)
    }
    
    if right < left {
        backtrack(ret, ans+")", left, right+1, n)
    }
}
