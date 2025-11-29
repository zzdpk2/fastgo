package v1

import (
	"time"
)

type Post struct {
	PostID    string    `json:"postID"`
	UserID    string    `json:"userID"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreatePostResponse struct {
	PostID string `json:"postID"`
}

type UpdatePostRequest struct {
	PostID  string  `json:"postID" uri:"postID"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

type UpdatePostResponse struct{}

type DeletePostRequest struct {
	PostIDs []string `json:"postIDs"`
}

type DeletePostResponse struct{}

type GetPostRequest struct {
	PostID string `json:"postID" uri:"postID"`
}

type GetPostResponse struct {
	Post *Post `json:"post"`
}

type ListPostRequest struct {
	Offset int64   `json:"offset"`
	Limit  int64   `json:"limit"`
	Title  *string `json:"title"`
}

type ListPostResponse struct {
	TotalCount int64   `json:"total_count"`
	Posts      []*Post `json:"posts"`
}
