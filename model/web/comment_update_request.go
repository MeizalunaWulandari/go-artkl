package web

type CommentUpdateRequest struct {
	Id        int
	Content   string
	UserId    int
	ArticleId int
}
