/*
X = BDCABA
Y = ABCBDAB => Longest Comman Subsequence is B C B

Dynamic Programming method : O ( n )
*/

package base_1

import "fmt"

func Max(more ...int) int {
	max_num := more[0]
	for _, elem := range more {
		if max_num < elem {
			max_num = elem
		}
	}
	return max_num
}

func Longest(str1, str2 string) int {
	len1 := len(str1)
	len2 := len(str2)

	//in C++,
	//int tab[m + 1][n + 1];
	//tab := make([][100]int, len1+1)

	tab := make([][]int, len1+1)
	for i := range tab {
		tab[i] = make([]int, len2+1)
	}

	i, j := 0, 0
	for i = 0; i <= len1; i++ {
		for j = 0; j <= len2; j++ {
			if i == 0 || j == 0 {
				tab[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				tab[i][j] = tab[i-1][j-1] + 1
				if i < len1 {
					fmt.Printf("%c", str1[i])
				}
			} else {
				tab[i][j] = Max(tab[i-1][j], tab[i][j-1])
			}
		}
	}
	fmt.Println()
	return tab[len1][len2]
}

func main() {
	str1 := "AGGTABTABTABTAB"
	str2 := "GXTXAYBTABTABTAB"
	fmt.Println(Longest(str1, str2))
	//Actual Longest Common Subsequence: GTABTABTABTAB
	//GGGGGTAAAABBBBTTTTAAAABBBBTTTTAAAABBBBTTTTAAAABBBB
	//13

	str3 := "AGGTABGHSRCBYJSVDWFVDVSBCBVDWFDWVV"
	str4 := "GXTXAYBRGDVCBDVCCXVXCWQRVCBDJXCVQSQQ"
	fmt.Println(Longest(str3, str4))
	//Actual Longest Common Subsequence: ?
	//GGGTTABGGGHHRCCBBBBBBYYYJSVDDDDDWWWFDDDDDVVVSSSSSBCCCBBBBBBVVVDDDDDWWWFWWWVVVVVV
	//14
}
