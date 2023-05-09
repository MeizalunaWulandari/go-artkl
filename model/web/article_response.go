package web

type ArticleResponse struct {
	Id         int
	Slug       string
	Title      string
	Content    string
	CategoryId int
	UserId     int
	Status     int
	Views      int
	CreatedAt  string
}
