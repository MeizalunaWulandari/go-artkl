package domain

type Article struct {
	Id         int
	Slug       string
	Title      string
	Content    string
	CategoryId int
	UserId     int
	Status     int
}
