package stringchallengeutil

import("fmt")

type StringReverse struct{
	Actual string
}

func (stringRev StringReverse) ReverseString() string{
	runes := []rune (stringRev.Actual)

	length := len(runes)

	for i , j := 0, length-1; i<j;i,j = i+1, j-1{
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func PrintStringValue(stringValue StringProblemInterface){

	fmt.Println(stringValue.ReverseString())

}