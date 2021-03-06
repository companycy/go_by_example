/**********************************************************************************
*
* Suppose a sorted array is rotated at some pivot unknown to you beforehand.
*
* (i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
*
* Find the minimum element.
*
* You may assume no duplicate exists in the array.
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getMin(l []int, low, high int) int {
    if len(l) == 0 || low > high {
        return -1
    } else if low == high {
        return l[low]
    } else if (high - low) == 1 {
        if l[low] > l[high] {
            return l[high]
        } else {
            return l[low]
        }
    }

    mid := low + (high-low)/2
    fmt.Println(low, mid, high)
    fmt.Println(l[low], l[mid], l[high])
    fmt.Println()
    if l[low] < l[mid] {
        return getMin(l, mid, high)
    } else {
        return getMin(l, low, mid)
    }
}

func main() {
    l := []int{
        // 4, 5, 6, 7, 1, 2,
        4, 5, 6, 7, 0, 1, 2,
    }

    min := getMin(l, 0, len(l)-1)
    fmt.Println(min)
}
