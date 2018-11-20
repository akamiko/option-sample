package option

import "fmt"

type Search struct {
	title string
	limit int
}

type SearchOption func(*Search)

func SearchTitle(title string) SearchOption {
	return func(s *Search) {
		s.title = title
	}
}

func SearchLimit(limit int) SearchOption {
	return func(s *Search) {
		s.limit = limit
	}
}

func Option(options ...SearchOption) {
	opt := Search{}
	for _, o := range options {
		o(&opt)
	}
	fmt.Println(opt)
	// use opt
}
