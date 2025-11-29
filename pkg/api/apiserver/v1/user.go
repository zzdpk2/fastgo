package v1

import "time"

type User struct {
	UserID    string    `json:"userID"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	PostCount int64     `json:"postCount"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateUserRequest struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Nickname *string `json:"nickname"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
}

type CreateUserResponse struct {
	UserID string `json:"userID"`
}

type UpdateUserRequest struct {
	Username *string `json:"username"`
	Nickname *string `json:"nickname"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}

type UpdateUserResponse struct{}

type DeleteUserRequest struct{}

type GetUserRequest struct{}

type GetUserResponse struct {
	User *User `json:"user"`
}

type ListUserRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type ListUserResponse struct {
	TotalCount int64   `json:"totalCount"`
	Users      []*User `json:"users"`
}
