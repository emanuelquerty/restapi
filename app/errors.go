package app

type appError struct {
	Error   error  `json:"error,omitempty"  bson:"error"`
	Message string `json:"message,omitempty"  bson:"message"`
	Code    int    `json:"code,omitempty"  bson:"code"`
}
