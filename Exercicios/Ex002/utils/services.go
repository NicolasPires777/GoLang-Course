package utils

import "fmt"

var Books []Book
var Publishers []Publisher
var Authors []Author

func CreateBook(name string, year uint32, author Author, publisher Publisher, gender string, value float32){
	book := Book{name, year, author, publisher, gender, value}
	Books := append(Books,book)
	fmt.Println("Book registered successfully! List of Books: ",Books)
}

func CreatePublisher(name string, nacionality string){
	publisher := Publisher{name, nacionality}
	Publishers := append(Publishers, publisher)
	fmt.Println("Publisher registered successfully! List of Publishers: ",Publishers)
}

func CreateAuthor(name string, nacionality string){
	author := Author{name, nacionality}
	Authors := append(Authors, author)
	fmt.Println("Author registered successfully! List of Authors: ",Authors)
}

func FormBook(){
	
}

