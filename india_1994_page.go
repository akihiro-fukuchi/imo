// 1994 India Regional Mathematical Olympiad
// https://mindyourdecisions.com/blog/2018/04/19/the-seemingly-impossible-missing-book-pages-puzzle-from-india/

package main

import (
    "fmt"
)

const (
    MAX = 175
    EXPECTED = 15000
)

func main() {

    for i := 1; i <= MAX; i++ {
        p := sigma(i)
        for j := 1; j <= i; j++ {
            q := p - (2 * j + 1)
            if q == EXPECTED {
                fmt.Printf("%d:%d:%d\n", i, j, q)
            }
        }
    }
    fmt.Print("not found")
}


func sigma(n int) int {
    res := 0
    for i := 0; i <= n; i++ {
        res = res + i
    }
    return res
}

