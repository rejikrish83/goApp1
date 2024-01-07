package main

import (
	"github.com/rejikrish83/goApp1/stringchallengeutil"
)

func main() {
	var reversString = stringchallengeutil.StringReverse{Actual: "Malayalam"}
	stringchallengeutil.PrintStringValue(reversString)
}
