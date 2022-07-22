// Sort array terlebih dahulu, kemudian rotasi ke kiri sesuai dengan nilai yang telah ditentukan.
//
// Contoh Sort array:
// Input: [4,5,2,1,3]
// Output: [1,2,3,4,5]

//Contoh RotateLeft:
//Input: 4, [1,2,3,4,5]
//Output: [5,1,2,3,4]

// Explanation RotateLeft:
// untuk melakukan rotasi kiri dengan nilai 4, array mengalami urutan perubahan berikut:
// [1,2,3,4,5] -> [2,3,4,5,1] -> [3,4,5,1,2] -> [4,5,1,2,3] -> [5,1,2,3,4]

package main

import "fmt"

func main() {
	var nums = []int{4, 5, 2, 1, 3}
	arrSorted := Sort(nums)
	fmt.Println(arrSorted)
	rotateLeft := RotateLeft(2, arrSorted)
	fmt.Println(rotateLeft)
}

func Sort(arr []int) []int {
	swapped := false                  //Untuk memeriksa apakah array sudah diurutkan; kemudian return;
	for i := 0; i < len(arr)-1; i++ { // looping parent
		for j := 0; j < len(arr)-1; j++ { // child loop
			if arr[j] > arr[j+1] {
				//elemen bertukar
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped { // kebalikan
			return arr
		}
	}
	return arr
}

func RotateLeft(d int, arr []int) []int {
	// panjang/length array -> variabel baru
	length := len(arr) // misalnya 5
	// copy arr ambil mulai indeks d sampai len(arr)-1 -> subslice
	copy := arr[d:length] // arr[3:5]
	// looping arr indeks 0 sampai sebelum indeks d -> append valuenya ke copy arr/result
	for _, val := range arr[0:d] {
		copy = append(copy, val)
	}
	// return copy
	return copy
	// return []int{} // TODO: replace this
}
