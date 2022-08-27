package main

import (
	"fmt"
	obj "objects/objects"
)

func main() {

	desk := obj.Desk{}
	desk.SetId(1)
	desk.SetName("desk name")
	desk.SetActive(true)

	fmt.Println(desk)
	desk.Open(OnFail, OnSuccess)

	desk2 := obj.Desk{}
	desk2.SetId(2)
	desk2.SetName("desk 2 name")
	desk2.SetActive(false)

	fmt.Println(desk2.GetName())
	desk2.Open(OnFail, OnSuccess)

}

func OnFail() {
	fmt.Println("open fail")
}

func OnSuccess() {
	fmt.Println("open success")
}
