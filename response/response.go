package response

import (
	"restapi/domain"
)

type Response struct {
	Description  string  `json:"description,omitempty"  bson:"description"`
	ErrorMessage string  `json:"error_message,omitempty"  bson:"error_message"`
	User         *User   `json:"user,omitempty"  bson:"user"`
	Users        *[]User `json:"users,omitempty"  bson:"users"`
}

func New(descripion string) *Response {
	return &Response{
		Description: descripion,
	}
}

func (r *Response) WithError(err error) *Response {
	r.ErrorMessage = err.Error()
	return r
}

func (r *Response) WithUser(user domain.User) *Response {
	respUser := domainToResponseUser(user)
	r.User = &respUser
	return r
}

func (r *Response) WithUsers(users []domain.User) *Response {
	var respUsers []User

	for _, currUser := range users {
		user := domainToResponseUser(currUser)
		respUsers = append(respUsers, user)
	}
	r.Users = &respUsers
	return r
}
