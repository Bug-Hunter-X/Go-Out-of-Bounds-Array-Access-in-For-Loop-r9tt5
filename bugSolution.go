func myFunc(a []int) int {
    sum := 0
    for i := 0; i < len(a); i++ {
        sum += a[i]
    }
    return sum
} 

//Alternative using range
func myFuncRange(a []int) int {
    sum := 0
    for _, v := range a {
        sum += v
    }
    return sum
}