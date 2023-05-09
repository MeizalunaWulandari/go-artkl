package web

type CommentUpdateRequest struct {
	Id        int
	Content   string
	UserId    string
	ArticleId int
}
