package iteration

import "strings"

// Strings are immutable, so everytime you add another string to the "repeated" variable, 
// it copies memory to accomodate new string which can make it ineffecient
func Repeat(str string, count int) string {
	var repeated string
	for i:=0; i < count; i++ {
		repeated = repeated + str
	}
	return repeated
}

// This is more effecieint because it minimizes memory copying, WriteString concatenates strings together 
func RepeatOptimized(str string, count int) string {
	var repeated strings.Builder
	for i:=0; i < count; i++ {
		repeated.WriteString(str)
	}
	return repeated.String()
}

