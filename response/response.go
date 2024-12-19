package response

import (
	"restapi/domain"
)

type Response struct {
	Description  string               `json:"description,omitempty"  bson:"description"`
	ErrorMessage string               `json:"error_message,omitempty"  bson:"error_message"`
	User         *domain.PublicUser   `json:"user,omitempty"  bson:"user"`
	Users        *[]domain.PublicUser `json:"users,omitempty"  bson:"users"`
}

func New(descripion string) *Response {
	return &Response{
		Description: descripion,
	}
}

func (r *Response) WithUser(user domain.User) *Response {
	publicUser := mapToPublicUser(user)
	r.User = &publicUser
	return r
}

func (r *Response) WithUsers(users []domain.User) *Response {
	var publicUsers []domain.PublicUser

	for _, currUser := range users {
		user := mapToPublicUser(currUser)
		publicUsers = append(publicUsers, user)
	}
	r.Users = &publicUsers
	return r
}
