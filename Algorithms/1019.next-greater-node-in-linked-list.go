func nextLargerNodes(head *ListNode) []int {
    res := []int{}
    list := []int{}
    
    ptr := head
    for ptr != nil {
        list = append(list, ptr.Val)
        res = append(res, 0)
        ptr = ptr.Next
    }
    
    for i := 1; i < len(list); i++ {
        for j := 0; j < i; j++ {
            if res[j] == 0 && list[j] < list[i] {
                res[j] = list[i]
            }
        }
    }
    
    return res
}
