package web

type ArticleUpdateRequest struct {
	Id         int
	Slug       string
	Title      string
	Content    string
	CategoryId int
	UserId     int
}
