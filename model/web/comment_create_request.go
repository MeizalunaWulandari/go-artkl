package web

type CommentCreateRequest struct {
	Content   string
	UserId    int
	ArticleId int
}
