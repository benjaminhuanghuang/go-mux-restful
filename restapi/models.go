package main

type Book struct {
	ID     string  `josn:"id"`
	Isbn   string  `josn:"isbn"`
	Title  string  `josn:"title"`
	Author *Author `josn:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
