package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

type Player struct {
	*User
	GameId int
}

func main() {
	player := Player{
		User:   &User{Id: 1, Name: "Test Name", Location: "Earth"},
		GameId: 2,
	}
	fmt.Println(player.Location, player.Id, player.GameId)
}
