package main

import (
	"fmt"

	api "../api"
)

func main() {
	res := api.RequestObject()
	fmt.Println(res[0]["licence"])
	//for _, item := range res {
	//fmt.Println(item)
	//}
}
