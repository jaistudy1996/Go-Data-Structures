package main

import ("fmt")

type author struct {
	firstName string
	lastName string
	age int
}

func (a author) fullName() string {
	return a.firstName + " " + a.lastName
}

type post struct {
	title string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)	
	fmt.Println("Content: ", p.content)
	fmt.Println("Name of author: ", p.author.fullName())
}

func main() {
	author1 := author{
		firstName: "First",
		lastName: "Last",
		age: 21,
	}

	post := post{
		title: "Title",
		content: "Content for title",
		author: author1,
	}
	post.details()
}
