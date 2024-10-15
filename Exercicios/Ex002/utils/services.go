package utils

import (
	"fmt"
	"bufio"
	"os"
)

var Books []Book
var Publishers []Publisher
var Authors []Author

var year uint32
var gender string
var value float32
var author int8
var publisher int8
var nacionality string

func CreateBook(name string, year uint32, author Author, publisher Publisher, gender string, value float32){
	book := Book{name, year, author, publisher, gender, value}
	Books = append(Books,book)
	fmt.Println("Book registered successfully! List of Books: ",Books)
}

func CreatePublisher(name string, nacionality string){
	publisher := Publisher{name: name, nacionality: nacionality}
	Publishers = append(Publishers, publisher)
	fmt.Println("Publisher registered successfully! List of Publishers: ")
	ListPublishers()
}

func CreateAuthor(name string, nacionality string){
	author := Author{name, nacionality}
	Authors = append(Authors, author)
	fmt.Println("Author registered successfully! List of Authors: ")
	ListAuthors()
}

func FormBook(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the book name: ")
	name, err := reader.ReadString('\n')
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	fmt.Print("Enter the book year: ")
	_, err= fmt.Scanln(&year)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	fmt.Print("Enter the book gender: ")
	gender, err= reader.ReadString('\n')
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

	fmt.Println("Registered Authors: ")
	ListAuthors()
	fmt.Print("Enter the book author: ")
	_, err= fmt.Scanln(&author)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	fmt.Println("Registered Publishers")
	ListPublishers()
	fmt.Print("Enter the book publisher: ")
	_, err= fmt.Scanln(&publisher)
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	CreateBook(name,year,Authors[author-1],Publishers[publisher-1],gender,value)
}

func FormAuthor(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the author name: ")
	name, err := reader.ReadString('\n')
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	fmt.Print("Enter the author nacionality: ")
	nacionality, err= reader.ReadString('\n')
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	CreateAuthor(name,nacionality)
}

func FormPublisher(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the publisher name: ")
	name, err := reader.ReadString('\n')
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	fmt.Print("Enter the publisher nacionality: ")
	nacionality, err= reader.ReadString('\n')
	if err!= nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	CreatePublisher(name,nacionality)
}

func ListBooks(){
	if len(Books) == 0{
		fmt.Println("No registered author")
	} else {
		for i, book := range Books{
			fmt.Println(i+1," - ",book.name)
		}
	}
}

func ListAuthors(){
	if len(Authors) == 0{
		fmt.Println("No registered author")
	} else {
		for i, author := range Authors{
			fmt.Println(i+1," - ",author.name)
		}
	}
}

func ListPublishers(){
	if len(Publishers) == 0{
		fmt.Println("No registered publisher")
	} else {
		for i, publisher := range Publishers{
			fmt.Println(i+1," - ",publisher.name)
		}
	}
}

