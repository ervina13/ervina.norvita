// Mengecek jika dua string adalah anagram
// Anagram adalah kata yang dibentuk melalui penataan ulang huruf-huruf dari beberapa kata lain.
//
// Contoh 1
// Input: a = "keen", b = "knee"
// Output: "Anagram"
// Explanation: Jika ditata, "knee" dan "keen" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 2
// Input: a = "fried", b = "fired"
// Output: "Anagram"
// Explanation: Jika ditata, "fried" dan "fired" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 3
// Input: a = "apple", b = "paddle"
// Output: "Bukan Anagram"
// Explanation: Jika ditata, "apple" dan "paddle" memiliki huruf-huruf yang berbeda

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 = "fried"
	var str2 = "fired"
	fmt.Println(AnagramsChecker(str1, str2))
}

func AnagramsChecker(str1 string, str2 string) string {
	var isCheck1 []bool
	var isCheck2 []bool
	var countAns int

	if strings.Count(str1, "") != strings.Count(str2, "") {
		return "Bukan Anagram"
	}

	for k := 0; k < strings.Count(str1, ""); k++ {
		isCheck1 = append(isCheck1, false)
		isCheck2 = append(isCheck2, false)
	}

	for i := 0; i < strings.Count(str1, "")-1; i++ {
		for j := 0; j < strings.Count(str2, "")-1; j++ {
			if isCheck1[i] == false && isCheck2[j] == false {
				countAns++
				isCheck1[i] = true
				isCheck2[j] = true
				break
			}
		}
	}

	if countAns == strings.Count(str1, "")-1 {
		return "Anagram"
	}

	return "Bukan Anagram"
}
