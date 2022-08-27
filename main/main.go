package main

import (
	"fmt"
	obj "objects/objects"
)

func main() {

	desk := obj.Desk{}
	desk.SetCallback(func(flag bool) {
		fmt.Println(flag)
	})
	desk.SetId(1)
	desk.SetName("desk name")
	desk.SetActive(true)

	desk.Open(OnFail, OnSuccess)

	desk2 := obj.Desk{}
	desk2.SetCallback(func(flag bool) {
		fmt.Println(flag)
	})

	desk2.SetId(2)
	desk2.SetName("desk 2 name")
	desk2.SetActive(false)

	desk2.Open(OnFail, OnSuccess)

}

func OnFail() {
	fmt.Println("open fail")
}

func OnSuccess() {
	fmt.Println("open success")
}
