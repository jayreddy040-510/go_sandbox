package main

import "fmt"

func main() {
    fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
}

func longestCommonPrefix(strs []string) string {
    i:=0
    if len(strs) < 1 {
        return ""
    }
    if len(strs[0]) < 1 {
        return ""
    }
    for i < len(strs[0]) {
        save := strs[0][i]
        for _,word := range strs {
            if i >= len(word) || word[i] != save{
                return strs[0][:i]
            }
        }
        i++
    }
    return strs[0][:i]
}
