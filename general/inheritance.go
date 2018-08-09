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

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("Contents of website")
	for _, v := range w.posts {
		v.details()
	}
}

func main() {
	author1 := author{
		firstName: "First",
		lastName: "Last",
		age: 21,
	}

	post1 := post{
		title: "Title",
		content: "Content for title",
		author: author1,
	}

	post2 := post{
		title: "Title for second post",
		content: "Content for second post",
		author: author1,
	}

	post3 := post{
		title: "Title for third post",
		content: "Content for third post",
		author: author1,
	}

	website := website{
		posts: []post{post1, post2, post3},
	}
	
	website.contents()
}
