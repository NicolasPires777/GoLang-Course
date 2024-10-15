package main

import (
	"fmt"
	"modulo/utils"
	"time"
)

func main(){
	var decision uint8

	for {
		fmt.Print("1 - Register a book\n2 - Register a publisher\n3 - Register a author\n4 - List Authors\n5 - List Publishers\n6 - List Books\n7 - Stop Program\nSelect an option: ")
		fmt.Scanln(&decision)

		if(decision==1){
			utils.FormBook()
		} else if (decision == 2){
			utils.FormPublisher()
		} else if (decision == 3){
			utils.FormAuthor()
		} else if (decision == 4){
			utils.ListAuthors()
		} else if (decision == 5){
			utils.ListPublishers()
		} else if (decision == 6){
			utils.ListBooks()
		} else if (decision == 7){
			break
		} else {
			fmt.Println("This option doesn't exists")
		}
		time.Sleep(2*time.Second)
		fmt.Println("=============================================================================") 
		}
}