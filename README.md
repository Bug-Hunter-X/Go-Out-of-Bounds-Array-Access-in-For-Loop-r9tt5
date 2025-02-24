# Go Out-of-Bounds Array Access Bug

This repository demonstrates a common off-by-one error in Go when iterating over arrays using a `for` loop. The bug is in the loop condition `i <= len(a)`, which causes an out-of-bounds access when `i` reaches `len(a)`. The correct condition should be `i < len(a)`.  This example highlights the importance of carefully considering loop boundaries when working with arrays and slices in Go. 

## Bug Report
The provided `bug.go` file contains a function `myFunc` that calculates the sum of elements in an integer array. However, due to an incorrect loop condition, it attempts to access an element beyond the array's bounds, leading to a runtime panic.