package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [2246]")
	examplesIn := [][]string{
		{
			`[-1,0,0,1,1,2]`,
			`"abacbe"`,
		},
		{
			`[-1,0,0,0]`,
			`"aabc"`,
		},
	}
	examplesOut := [][]string{
		{
			`3`,
		},
		{
			`3`,
		},
	}
	if err := testutil.RunLeetCodeFunc(t, longestPath, examplesIn, examplesOut); err != nil {
		t.Fatal(err)
	}
}
