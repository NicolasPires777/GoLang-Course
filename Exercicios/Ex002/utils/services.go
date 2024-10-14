package utils

import "fmt"

var Books []Book
var Publishers []Publisher
var Authors []Author

var name string
var year uint32
var gender string
var value float32
var author int8
var publisher int8
var nacionality string

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
	fmt.Print("Enter the book name: ")
	_, err:= fmt.Scanln(&name)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	fmt.Print("Enter the book year:")
	_, err= fmt.Scanln(&year)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	fmt.Print("Enter the book gender: ")
	_, err= fmt.Scanln(&gender)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	fmt.Print("Enter the book price: ")
	_, err= fmt.Scanln(&value)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	ListAuthors()
	fmt.Print("Enter the book author: ")
	_, err= fmt.Scanln(&author)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	ListPublishers()
	fmt.Print("Enter the book publisher: ")
	_, err= fmt.Scanln(&publisher)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	CreateBook(name,year,Authors[author],Publishers[publisher],gender,value)
}

func FormAuthor(){
	fmt.Println("Enter the author name:")
	_, err := fmt.Scanln(&name)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	fmt.Println("Enter the author nacionality: ")
	_, err= fmt.Scanln(&nacionality)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	CreateAuthor(name,nacionality)
}

func FormPublisher(){
	fmt.Println("Enter the publisher name:")
	_, err := fmt.Scanln(&name)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	fmt.Println("Enter the publisher nacionality: ")
	_, err= fmt.Scanln(&nacionality)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	CreatePublisher(name,nacionality)
}

func ListAuthors(){
	if len(Authors) == 0{
		fmt.Println("No registered author")
	} else {
		for i, author := range Authors{
			fmt.Println(i," - ",author.name)
		}
	}
}

func ListPublishers(){
	if len(Publishers) == 0{
		fmt.Println("No registered publisher")
	} else {
		for i, publisher := range Publishers{
			fmt.Println(i," - ",publisher.name)
		}
	}
}

