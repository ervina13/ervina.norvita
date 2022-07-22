package main

import (
	//"container/list"

	"strings"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/parentheses-validation/stack"
)

// Salah satu problem populer yang dapat diselesaikan dengan menggunakan Stack adalah mengecek validitas tanda kurung.
// Diberikan sebuah string yang hanya terdapat karakter '(', ')', '{', '}', '[', dan ']'.
// Tentukan apakah sebuah string merupakan sekumpulan tanda kurung yang valid.
// String tanda kurung yang valid adalah
// 1. Tanda kurung buka harus ditutup oleh pasangannya.
// 2. Tanda kurung buka harus ditutup di urutan yang benar.

// Contoh 1
// Input: s = "()"
// Output: true
// Penjelasan: tanda kurung buka '(' ditutup dengan pasangannya yaitu ')'.

// Contoh 2
// Input: s = "{)"
// Output: false
// Penjelasan: tanda kurung buka '{' tidak ditutup oleh pasangannya.

// Contoh 3
// Input: s = "([])"
// Output: true
// Penjelasan: tanda kurung buka ditutup dengan pasangannya dan sesuai dengan urutan.

func IsValidParentheses(s string) bool {
	// TODO: answer here
	if len(s)%2 != 0 {
		return false
	}
	Stack := new(stack.Stack)
	Stack.Top = -1

	data := map[string]int{
		"(": 1, ")": -1, "{": 2, "}": -2, "[": 3, "]": -3,
	}
	for i := 0; i < strings.Count(s, "")-1; i++ {
		if data[string(s[i])] > 0 {
			Stack.Push(rune(s[i]))
		} else {
			peekVal, err := Stack.Peek()
			if err == nil {
				if data[string(s[i])]+data[string(peekVal)] == 0 {
					Stack.Pop()
				}
			}
		}
	}
	if len(Stack.Data) == 0 && data[string(s[0])] > 0 {
		return true
	}
	return false
}
