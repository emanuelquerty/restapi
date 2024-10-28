package testdata

import "restapi/domain"

var Users = []domain.User{
	domain.User{
		ID:           1,
		FirstName:    "Peter",
		LastName:     "Petrelli",
		Email:        "ppetrelli@email.com",
		PasswordHash: "bGVhbGRhZGU=khjs8e90020283jsl0",
	},
	domain.User{
		ID:           2,
		FirstName:    "Leroy",
		LastName:     "Jenkis",
		Email:        "ljenkins@email.com",
		PasswordHash: "bGVhbGRhZGU=t5348hnns028873kasxx",
	},
}
