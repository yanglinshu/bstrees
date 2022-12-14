package main

import (
	"bstrees/pkg/scapegoat"
	"bstrees/pkg/trait/number"
	"bstrees/pkg/util/console"
	"bufio"
	"fmt"
	"os"
)

func ReadWithPanic[T number.Integer](gin *bufio.Reader) T {
	value, err := console.Read[T](gin)
	if err != nil {
		panic(err)
	}
	return value
}

func main() {
	tree := scapegoat.New[int](0.7)
	gin := bufio.NewReader(os.Stdin)
	n := ReadWithPanic[int](gin)
	for i := 0; i < n; i++ {
		opt := ReadWithPanic[int](gin)
		value := ReadWithPanic[int](gin)
		// fmt.Println("----------------")
		// fmt.Println(opt, value)
		switch opt {
		case 1:
			tree.Insert(value)
		case 2:
			tree.Delete(value)
		case 3:
			fmt.Println(tree.Rank(value))
		case 4:
			{
				kth, err := tree.Kth(uint32(value))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(kth)
			}
		case 5:
			{
				kth, err := tree.Prev(value)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(kth)
			}
		case 6:
			{
				kth, err := tree.Next(value)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(kth)
			}
		}
		// if opt == 1 || opt == 2 {
		// 	tree.Print()
		// }
	}
}
