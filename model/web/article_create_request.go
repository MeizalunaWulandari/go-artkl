package web

type ArticleCreateRequest struct {
	Slug       string
	Title      string
	Content    string
	CategoryId int
	UserId     int
}
