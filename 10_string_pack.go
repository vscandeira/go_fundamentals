//example code with string package

package main

import (
	"fmt"
	"strings"
)

func main() {
	//Contains function returns true if the string contains the substring
	fmt.Println(strings.Contains("I am a string", "am")) //true

	//Count function returns the number of non-overlapping instances of a substring
	fmt.Println(strings.Count("I am a string", "a")) //2

	//HasPrefix function returns true if the string starts with the specified prefix
	fmt.Println(strings.HasPrefix("I am a string", "I")) //true

	//HasSuffix function returns true if the string ends with the specified suffix
	fmt.Println(strings.HasSuffix("I am a string", "string")) //true

	//Index function returns the index of the first instance of the substring in the string
	fmt.Println(strings.Index("I am a string", "a")) //2

	//Join function concatenates the elements of a string array to create a single string
	fmt.Println(strings.Join([]string{"I", "am", "a", "string"}, " ")) //I am a string

	//Repeat function returns a new string consisting of count copies of the string
	fmt.Println(strings.Repeat("a", 5)) //aaaaa

	//Replace function returns a copy of the string with the first n instances of old replaced by new
	fmt.Println(strings.Replace("I am a string", "a", "an", 2)) //I anm an string

	//Split function splits the string into a slice of substrings separated by the separator
	fmt.Println(strings.Split("I am a string", " ")) //[I am a string]

	//ToLower function returns a copy of the string with all Unicode letters mapped to their lower case
	fmt.Println(strings.ToLower("I AM A STRING")) //i am a string

	//ToUpper function returns a copy of the string with all Unicode letters mapped to their upper case
	fmt.Println(strings.ToUpper("i am a string")) //I AM A STRING

	//Trim function returns a slice of the string with all leading and trailing characters contained in cutset removed
	fmt.Println(strings.Trim("I am a string", "I")) // am a string
}
