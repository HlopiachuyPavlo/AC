package main

import (
	"testing"
)

func TestTime(t *testing.T) {
	startTrack()

	rangeStr := "cdefghijklmnopqrstuvwxyz0123456789"
	rangeN := 70

	templ = "a"
	str := "b"
	for i := 1; i <= rangeN; i++ {
		for _, r := range rangeStr {
			checkStr := str + string(r)

			actual := Distance(templ, checkStr)

			if actual != i+1 {
				t.Fatalf("Distance(%v, %v) was expected to return %v but returned %d.",
					templ, checkStr, i+1, actual)
			}

		}
		templ += "a"
		str += "b"

	}
	finishTrack()

}
