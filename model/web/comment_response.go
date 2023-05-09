package web

type CommentResponse struct {
	Id        int
	Content   string
	UserId    int
	ArticleId int
	CreatedAt string
}
