package response

type UserWithNoPasswordField struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

type SingleUser struct {
	Description string                  `json:"description,omitempty"`
	User        UserWithNoPasswordField `json:"user,omitempty"`
}

type MultiUser struct {
	Description string                    `json:"description,omitempty"`
	Users       []UserWithNoPasswordField `json:"users,omitempty"`
}

type Generic struct {
	Description string
}
