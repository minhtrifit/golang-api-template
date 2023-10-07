package helper

import (
	"fmt"
	"go-api/datatype"
)

func DeleteSlice(originalArray []datatype.Album, i int) (newArray []datatype.Album) {
	// check if the index is within array bounds
	if i < 0 || i >= len(originalArray) {
		fmt.Println("The given index is out of bounds.");
	} else {
		// delete an element from the array
		var newLength = 0;

		for index := range originalArray {
			if index != i {
				originalArray[newLength] = originalArray[index];
				newLength++;
			}
		}

		// reslice the array to remove extra index
		newArray = originalArray[:len(originalArray) - 1];
	}
	return
}