package main

import (
	//不用关心gopath src的问题 相对引用
	"hello-module/pkg"
)

func main() {
	pkg.Start()
}
