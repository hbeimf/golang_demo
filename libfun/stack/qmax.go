package stack

import (
	// llstack "github.com/emirpasic/gods/stacks/linkedliststack"
	"fmt"
)

func getMaxWindow(list []int32, w int) {

	fmt.Println(list)
	fmt.Println(w)
}


func TestGetMaxWindow() {

	list := []int32{4, 3, 5, 4, 3, 3, 6, 7}

	getMaxWindow(list, 3)

	// fmt.Println(list)


}
