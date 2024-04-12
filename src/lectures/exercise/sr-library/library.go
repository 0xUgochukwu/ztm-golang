//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import "fmt"
import "time"

type Book struct {
	title        string
	CheckedInAt  string
	CheckedOutAt string
	withMember   string
}

type Member struct {
	books []Book
	name  string
}

type Library struct {
	members []Member
	books   []Book
	logs    []string
}

func checkInBook(bookTitle string, memberName string, library *Library) {
	// get book from library
	var book *Book
	for i := range library.books {
		if library.books[i].title == bookTitle {
			book = &library.books[i]
			break
		}
	}

	if book == nil {
		panic("Book not found")
	}

	var member *Member
	for i := range library.members {
		if library.members[i].name == memberName {
			member = &library.members[i]
			break
		}
	}

	if member == nil {
		panic("Member not found")
	}

	book.CheckedOutAt = ""
	book.CheckedInAt = time.Now().String()
	book.withMember = ""
	// remove book from member
	for i, b := range member.books {
		if b.title == book.title {
			member.books = append(member.books[:i], member.books[i+1:]...)
			break
		}
	}
	library.logs = append(library.logs, "Book ["+book.title+"] checked in at "+book.CheckedInAt+" by "+member.name)
}

func checkOutBook(bookTitle string, memberName string, library *Library) {
	var book *Book
	for i := range library.books {
		if library.books[i].title == bookTitle {
			book = &library.books[i]
			break
		}
	}

	if book == nil {
		panic("Book not found")
	}

	var member *Member
	for i := range library.members {
		if library.members[i].name == memberName {
			member = &library.members[i]
			break
		}
	}

	if member == nil {
		panic("Member not found")
	}

	book.CheckedInAt = ""
	book.CheckedOutAt = time.Now().String()
	book.withMember = member.name
	member.books = append(member.books, *book)
	library.logs = append(library.logs, "Book ["+book.title+"] checked out at "+book.CheckedOutAt+" by "+member.name)
}

func printLibraryInformation(library Library) {
	fmt.Println("Library Information:")
	fmt.Println("Books:")
	// if ook is with a member label unavailable
	for _, book := range library.books {
		if book.withMember == "" {
			fmt.Println("  - " + book.title + " (available)")
		} else {
			fmt.Println("  - " + book.title + " (unavailable with " + book.withMember + ")")
		}
	}
	fmt.Println("Members:")
	for _, member := range library.members {
		fmt.Println("  - " + member.name)
	}
	fmt.Println("Logs:")
	for _, log := range library.logs {
		fmt.Println("  - " + log)
	}
	fmt.Println()
}

func main() {
	//* Perform the following:
	//  - Add at least 4 books and at least 3 members to the library
	//  - Check out a book
	//  - Check in a book
	//  - Print out initial library information, and after each change
	//* There must only ever be one copy of the library in memory at any time

	library := Library{
		books: []Book{
			{title: "Harry Potter"},
			{title: "The Hobbit"},
			{title: "The Great Gatsby"},
			{title: "The Catcher in the Rye"},
			{title: "1984"},
			{title: "To Kill a Mockingbird"},
		},
		members: []Member{
			{name: "John"},
			{name: "Jane"},
			{name: "Bob"},
			{name: "Alice"},
		},
	}

	printLibraryInformation(library)

	checkOutBook("Harry Potter", "John", &library)
	checkOutBook("The Hobbit", "Jane", &library)

	printLibraryInformation(library)

	checkInBook("Harry Potter", "John", &library)

	printLibraryInformation(library)

	checkOutBook("The Great Gatsby", "Bob", &library)

	printLibraryInformation(library)

	printLibraryInformation(library)
}
