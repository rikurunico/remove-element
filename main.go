package main

import "fmt"

func removeLeftRight(arr []interface{}, position string) []interface{} {
	if position == "left" {
		return arr[1:]
	} else if position == "right" {
		return arr[:len(arr)-1]
	} else {
		return arr
	}
}

func main() {
	arr := []interface{}{1, 2, 3, 4, 5}
	arr = removeLeftRight(arr, "left")
	arr = removeLeftRight(arr, "right")
	fmt.Println(arr)

	strArr := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	strArr = removeLeftRight(strArr, "left")
	strArr = removeLeftRight(strArr, "right")
	fmt.Println(strArr)
}
