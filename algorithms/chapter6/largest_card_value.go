package main

import (
	"fmt"
)

// 6.13
// s[1...n]  has value v[1...n]
// two players take turns to pick up first or last card to get max total value
// M[i, j] max total value when choose between s[i] and s[j]

// assume competitor also picks his max total value, so
// when pick s[i],
// case 1:
// my         [          ]
// i    i+1    i+2 .... j
//     other
// => M[i+2, j]

// case 2:
// my  [           ]
// i    i+1 ... j-1   j
//                  other
// => M[i+1, j-1]

// so M[i] = vi + min(M[i+2, j], M[i+1, j-1]) => other only leaves smaller one

// when pick s[j],
// case 1:
//       [                ]  my
//  i     i+1   ....   j-1   j
// other
// => M[i+1, j-1]

// case 2:
//[             ]       my
// i         j-2  j-1   j
//               other 
// => M[i, j-2]

// then M[i, j] = max(M[i], M[j])

// if j = i+1, M[i, j] = max(vi, vj)


func main() {
	
}
