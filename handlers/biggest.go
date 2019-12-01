package handlers

import (
	"fmt"
	"strconv"
)

func Biggest(arr []int) string {
	var n , biggest int

	for _,i := range arr {
		if i>n {
			n = i
			biggest = n
		}else{
			fmt.Println(i,"<",n)
		}
	}
	
	return strconv.Itoa(biggest+1)
}