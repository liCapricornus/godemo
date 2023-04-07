package main

import (
	"fmt"
	"godemo/demo2/infra"
)

func getRetriever() infra.Retriever {
	return infra.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	fmt.Println("hello demo2")
	var r retriever = getRetriever()
	fmt.Println(r.Get("http://www.baidu.com"))
}