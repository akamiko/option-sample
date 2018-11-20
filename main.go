package main

import (
	"github.com/akamiko/option-sample3/option"
)

func main() {
	option.Option(option.SearchTitle("hoge"), option.SearchLimit(10))
}
