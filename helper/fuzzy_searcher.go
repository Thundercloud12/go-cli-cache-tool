package helper

import (
	
)


func FuzzySearcher( a string, b string)(int){
	m:=len(a)
	n:=len(b)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}

	for i := 0; i < n+1; i++ {
		dp[0][i]=i
	}

	for j := 1; j < n+1; j++ {
		for i := 1; i < m+1; i++ {
			cost:=0
			if a[i-1] != b[j-1] {
				cost= 1
			}else {
				cost=0
			}

			dp[i][j] = min(dp[i-1][j-1]+cost, dp[i][j-1]+1, dp[i-1][j]+1)
		}
	}

	return dp[m][n]


}


